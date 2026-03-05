//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestPostWithTask(t *testing.T) { //nolint: funlen
	tests := map[string]struct {
		response         http.Response
		expectedTask     *TaskMonitorInfo
		expectedResponse bool
		expectedErr      string
	}{
		"post with sync response": {
			response: http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString("{}")),
			},
			expectedResponse: true,
		},
		"post with task monitor location": {
			response: http.Response{
				StatusCode: http.StatusAccepted,
				Header: http.Header{
					"Location": []string{"/TaskMonitor/1"},
				},
				Body: io.NopCloser(bytes.NewBufferString("garbage")),
			},
			expectedTask: &TaskMonitorInfo{
				TaskMonitor: "/TaskMonitor/1",
			},
		},
		"post with task monitor location and retry after seconds": {
			response: http.Response{
				StatusCode: http.StatusAccepted,
				Header: http.Header{
					"Location": []string{"/TaskMonitor/1"},
					// Retry-After is tested separately, use an absolute time to avoid delta diffs
					"Retry-After": []string{"Wed, 21 Oct 2015 07:28:00 GMT"},
				},
				Body: io.NopCloser(bytes.NewBufferString("garbage")),
			},
			expectedTask: &TaskMonitorInfo{
				TaskMonitor: "/TaskMonitor/1",
				RetryAfter:  time.Date(2015, 10, 21, 7, 28, 0, 0, time.UTC),
			},
		},
		"post with task monitor location, retry after and task body": {
			response: http.Response{
				StatusCode: http.StatusAccepted,
				Header: http.Header{
					"Location": []string{"/TaskMonitor/1"},
					// Retry-After is tested separately, use an absolute time to avoid delta diffs
					"Retry-After": []string{"Wed, 21 Oct 2015 07:28:00 GMT"},
				},
				Body: io.NopCloser(
					bytes.NewBufferString(
						`{"@odata.id": "/TaskService/Tasks/1", "@odata.type": "Task"}`,
					),
				),
			},
			expectedTask: &TaskMonitorInfo{
				TaskMonitor: "/TaskMonitor/1",
				RetryAfter:  time.Date(2015, 10, 21, 7, 28, 0, 0, time.UTC),
				Task: &Task{
					Entity: Entity{
						ODataID: "/TaskService/Tasks/1",
					},
					ODataType: "Task", // not actually the spec string
				},
			},
		},
		"post with error response": {
			response: http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(bytes.NewBufferString("{}")),
			},
			expectedErr: "500: {}",
		},
	}

	testURI := "/Action/test"
	testPayload := map[string]string{}
	var testHeaders map[string]string

	for name, test := range tests {
		test := test // to support older go
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			c := &TestClient{
				CustomReturnForActions: map[string][]any{
					http.MethodPost: {&test.response},
				},
			}

			res, taskMonitorInfo, err := PostWithTask(c, testURI, testPayload, testHeaders, false)
			if test.expectedErr != "" {
				RequireErrorContains(t, err, test.expectedErr)
			} else {
				RequireNoError(t, err)
			}

			if test.expectedTask != nil && test.expectedTask.Task != nil {
				test.expectedTask.Task.SetClient(c)
			}

			if test.expectedResponse {
				AssertEqual(t, &test.response, res)
			} else if test.expectedTask != nil {
				AssertEqualMsg(t, test.expectedTask.TaskMonitor, taskMonitorInfo.TaskMonitor,
					"expected task monitor %q, actual %q",
					test.expectedTask.TaskMonitor, taskMonitorInfo.TaskMonitor)
				if test.expectedTask.Task != nil {
					AssertEqualMsg(t, test.expectedTask.Task.ID, taskMonitorInfo.Task.ID,
						"expected task ID [%q], actual [%q]",
						test.expectedTask.Task.ID, taskMonitorInfo.Task.ID)
				}
			}
		})
	}
}

func TestWaitForTaskMonitor(t *testing.T) { //nolint: funlen
	startTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	testTaskMonitorURI := "/TaskService/TaskMonitors/1"

	complete0 := uint(0)
	complete100 := uint(100)
	tests := map[string]struct {
		PollRate        time.Duration
		CtxFunc         func(ctx context.Context) (context.Context, func())
		TaskMonitor     *TaskMonitorInfo
		ExpectedUpdates []*Task
		StatusChan      bool
		EndTime         time.Time
		ExpectErr       string
		Responses       []*http.Response
	}{
		"nil task monitor": {
			ExpectErr:   "task monitor is nil",
			TaskMonitor: nil,
			EndTime:     startTime,
		},
		"missing task uri": {
			ExpectErr:   "task monitor URI is missing",
			TaskMonitor: &TaskMonitorInfo{},
			EndTime:     startTime,
		},
		"done when first polled, error": {
			Responses: []*http.Response{
				{StatusCode: 500, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectErr:       "500:",
			ExpectedUpdates: []*Task{},
			StatusChan:      true,
			EndTime:         startTime, // no wait until first poll
			TaskMonitor: &TaskMonitorInfo{
				TaskMonitor: testTaskMonitorURI,
			},
		},
		"done when first polled, success": {
			Responses: []*http.Response{
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*Task{},
			StatusChan:      true,
			EndTime:         startTime,
			TaskMonitor: &TaskMonitorInfo{
				TaskMonitor: testTaskMonitorURI,
			},
		},
		"two polls, then success": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*Task{nil, nil},
			StatusChan:      true,
			EndTime:         startTime.Add(20 * time.Second),
			TaskMonitor: &TaskMonitorInfo{
				TaskMonitor: testTaskMonitorURI,
			},
		},
		"two polls, then success, nil status channel": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			EndTime: startTime.Add(20 * time.Second),
			TaskMonitor: &TaskMonitorInfo{
				TaskMonitor: testTaskMonitorURI,
			},
		},
		"two polls, then fail": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 500, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*Task{nil, nil},
			StatusChan:      true,
			EndTime:         startTime.Add(20 * time.Second),
			TaskMonitor: &TaskMonitorInfo{
				TaskMonitor: testTaskMonitorURI,
			},
			ExpectErr: "500:",
		},
		"two polls, then success, initial wait time": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*Task{nil, nil},
			StatusChan:      true,
			EndTime:         startTime.Add(40 * time.Second),
			TaskMonitor: &TaskMonitorInfo{
				RetryAfter:  startTime.Add(20 * time.Second),
				TaskMonitor: testTaskMonitorURI,
			},
		},
		"two polls, with retry-after in response": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202,
					Header: http.Header{"Retry-After": []string{"1"}},
					Body:   io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*Task{nil, nil},
			StatusChan:      true,
			EndTime:         startTime.Add((1 + 20 + 10) * time.Second),
			TaskMonitor: &TaskMonitorInfo{
				RetryAfter:  startTime.Add(20 * time.Second),
				TaskMonitor: testTaskMonitorURI,
			},
		},
		"context timeout during initial wait": {
			CtxFunc: func(ctx context.Context) (context.Context, func()) {
				return context.WithTimeout(ctx, 1*time.Nanosecond)
			},
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202,
					Header: http.Header{"Retry-After": []string{"1"}},
					Body:   io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*Task{},
			StatusChan:      true,
			EndTime:         startTime.Add(10 * time.Second),
			TaskMonitor: &TaskMonitorInfo{
				RetryAfter:  startTime.Add(20 * time.Second),
				TaskMonitor: testTaskMonitorURI,
			},
			ExpectErr: context.DeadlineExceeded.Error(),
		},
		"context timeout during task poll": {
			CtxFunc: func(ctx context.Context) (context.Context, func()) {
				return context.WithTimeout(ctx, 10*time.Millisecond)
			},
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202,
					Header: http.Header{"Retry-After": []string{"1"}},
					Body:   io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*Task{nil},
			StatusChan:      true,
			EndTime:         startTime.Add(10 * time.Second),
			TaskMonitor: &TaskMonitorInfo{
				RetryAfter:  startTime.Add(2 * time.Second),
				TaskMonitor: testTaskMonitorURI,
			},
			ExpectErr: context.DeadlineExceeded.Error(),
		},
		"two polls, then success with task status": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(
					`{"TaskState": "Running", "PercentComplete": 0}`,
				))},
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(
					`{"TaskState": "Running", "PercentComplete": 100}`,
				))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*Task{
				{TaskState: "Running", PercentComplete: &complete0},
				{TaskState: "Running", PercentComplete: &complete100},
			},
			StatusChan: true,
			EndTime:    startTime.Add(20 * time.Second),
			TaskMonitor: &TaskMonitorInfo{
				TaskMonitor: testTaskMonitorURI,
			},
		},
	}

	for name, test := range tests {
		test := test // to support older go versions
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// set a default poll rate so tests run quickly
			if test.PollRate == 0 {
				test.PollRate = 20 * time.Millisecond
			}

			// rescale retry-after to milliseconds until we have synctest
			if test.TaskMonitor != nil && !test.TaskMonitor.RetryAfter.IsZero() {
				test.TaskMonitor.RetryAfter = time.Now().Add(test.TaskMonitor.RetryAfter.Sub(startTime) * (time.Millisecond / time.Second))
			}

			// synctest.Test(t, func(t *testing.T) {
			c := &TestClient{
				CustomReturnForActions: map[string][]any{},
			}

			for _, resp := range test.Responses {
				c.CustomReturnForActions[http.MethodGet] = append(c.CustomReturnForActions[http.MethodGet], resp)
			}

			var statusChan chan *Task
			var statusChanWg sync.WaitGroup
			if test.StatusChan {
				statusChan = make(chan *Task)
				statusChanWg.Add(1)
				go func() {
					defer statusChanWg.Done()
					i := 0
					for status := range statusChan {
						RequireLessMsg(t, i, len(test.ExpectedUpdates), "test got %d response but only expects %d", i+1, len(test.ExpectedUpdates))
						expectedStatus := test.ExpectedUpdates[i]
						if expectedStatus != nil {
							expectedStatus.SetClient(c)
						}

						if expectedStatus != nil && status != nil {
							AssertEqualMsg(t, expectedStatus.TaskState, status.TaskState,
								"status update [%d], expected task state %q, actual %q",
								i, expectedStatus.TaskState, status.TaskState)
							if expectedStatus.PercentComplete == nil {
								AssertEqualMsg(t, nil, status.PercentComplete,
									"status update [%d], expected nil percent complete, actual [%d]",
									i, *status.PercentComplete)
							} else {
								AssertEqualMsg(t, *expectedStatus.PercentComplete, *status.PercentComplete,
									"status update [%d], expected percent complete [%d], actual [%d]",
									i, expectedStatus.PercentComplete, status.PercentComplete)
							}
						}
						// AssertEqualMsg(t, expectedStatus, status
						i++
					}
					AssertEqual(t, len(test.ExpectedUpdates), i)
				}()
			}

			ctx := context.Background()
			if test.CtxFunc != nil {
				var cancel func()
				ctx, cancel = test.CtxFunc(ctx)
				if cancel != nil {
					t.Cleanup(cancel)
				}
			}

			var resp *http.Response
			var err error
			done := make(chan bool)
			go func() {
				resp, err = WaitForTaskMonitor(ctx, c, test.PollRate, test.TaskMonitor, statusChan)
				if statusChan != nil {
					close(statusChan)
				}
				done <- true
			}()

			// for {
			// 	synctest.Wait()
			// 	select {
			// 	case <-done:
			// 	default:
			// 		time.Sleep(1 * time.Second) // advance time in the bubble
			// 		continue
			// 	}

			<-done
			statusChanWg.Wait()

			// EndTime comparisons will be a bit flaky without synctest, disabling
			// currTime := time.Now().UTC()
			// AssertEqual(t, test.EndTime.UTC(), currTime)

			if test.ExpectErr != "" {
				RequireErrorContains(t, err, test.ExpectErr)
			} else {
				RequireNoError(t, err)
				AssertEqual(t, test.Responses[len(test.Responses)-1], resp)
			}
			// 		return
			// 	}
			// })
		})
	}
}

func TestWaitForTaskMonitorObject(t *testing.T) {
	raw := `{"PowerState": "On"}`
	c := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {
				&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(raw)),
				},
			},
		},
	}

	expected := &ComputerSystem{
		PowerState: "On",
		RawData:    []byte(raw),
	}
	expected.SetClient(c)

	ctx := context.Background() // go1.24 adds t.Context()

	object, _, err := WaitForTaskMonitorObject[ComputerSystem](
		ctx, c, 0, &TaskMonitorInfo{TaskMonitor: "/Monitor"}, nil,
	)
	RequireNoError(t, err)
	AssertEqual(t, expected, object)
}
