package redfish_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostWithTask(t *testing.T) {
	tests := map[string]struct {
		response         http.Response
		expectedTask     *redfish.TaskMonitorInfo
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
			expectedTask: &redfish.TaskMonitorInfo{
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
			expectedTask: &redfish.TaskMonitorInfo{
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
			expectedTask: &redfish.TaskMonitorInfo{
				TaskMonitor: "/TaskMonitor/1",
				RetryAfter:  time.Date(2015, 10, 21, 7, 28, 0, 0, time.UTC),
				Task: &redfish.Task{
					Entity: common.Entity{
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

	testUri := "/Action/test"
	testPayload := map[string]string{}
	var testHeaders map[string]string

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			c := &common.TestClient{
				CustomReturnForActions: map[string][]interface{}{
					http.MethodPost: {&test.response},
				},
			}

			res, taskMonitorInfo, err := redfish.PostWithTask(c, testUri, testPayload, testHeaders, false)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
			}

			if test.expectedTask != nil && test.expectedTask.Task != nil {
				test.expectedTask.Task.SetClient(c)
			}

			if test.expectedResponse {
				assert.Equal(t, &test.response, res)
			} else {
				assert.Equal(t, test.expectedTask, taskMonitorInfo)
			}
		})
	}
}

func TestWaitForTaskMonitor(t *testing.T) {
	startTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	testTaskMonitorUri := "/TaskService/TaskMonitors/1"

	tests := map[string]struct {
		PollRate        time.Duration
		CtxFunc         func(ctx context.Context) (context.Context, func())
		TaskMonitor     *redfish.TaskMonitorInfo
		ExpectedUpdates []*redfish.Task
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
			TaskMonitor: &redfish.TaskMonitorInfo{},
			EndTime:     startTime,
		},
		"done when first polled, error": {
			Responses: []*http.Response{
				{StatusCode: 500, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectErr:       "500:",
			ExpectedUpdates: []*redfish.Task{},
			StatusChan:      true,
			EndTime:         startTime, // no wait until first poll
			TaskMonitor: &redfish.TaskMonitorInfo{
				TaskMonitor: testTaskMonitorUri,
			},
		},
		"done when first polled, success": {
			Responses: []*http.Response{
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*redfish.Task{},
			StatusChan:      true,
			EndTime:         startTime,
			TaskMonitor: &redfish.TaskMonitorInfo{
				TaskMonitor: testTaskMonitorUri,
			},
		},
		"two polls, then success": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*redfish.Task{nil, nil},
			StatusChan:      true,
			EndTime:         startTime.Add(20 * time.Second),
			TaskMonitor: &redfish.TaskMonitorInfo{
				TaskMonitor: testTaskMonitorUri,
			},
		},
		"two polls, then success, nil status channel": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			EndTime: startTime.Add(20 * time.Second),
			TaskMonitor: &redfish.TaskMonitorInfo{
				TaskMonitor: testTaskMonitorUri,
			},
		},
		"two polls, then fail": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 500, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*redfish.Task{nil, nil},
			StatusChan:      true,
			EndTime:         startTime.Add(20 * time.Second),
			TaskMonitor: &redfish.TaskMonitorInfo{
				TaskMonitor: testTaskMonitorUri,
			},
			ExpectErr: "500:",
		},
		"two polls, then success, initial wait time": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*redfish.Task{nil, nil},
			StatusChan:      true,
			EndTime:         startTime.Add(40 * time.Second),
			TaskMonitor: &redfish.TaskMonitorInfo{
				RetryAfter:  startTime.Add(20 * time.Second),
				TaskMonitor: testTaskMonitorUri,
			},
		},
		"two polls, with retry-after in response": {
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202,
					Header: http.Header{"Retry-After": []string{"30"}},
					Body:   io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*redfish.Task{nil, nil},
			StatusChan:      true,
			EndTime:         startTime.Add((30 + 20 + 10) * time.Second),
			TaskMonitor: &redfish.TaskMonitorInfo{
				RetryAfter:  startTime.Add(20 * time.Second),
				TaskMonitor: testTaskMonitorUri,
			},
		},
		"context timeout during initial wait": {
			CtxFunc: func(ctx context.Context) (context.Context, func()) {
				return context.WithTimeout(ctx, 10*time.Second)
			},
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202,
					Header: http.Header{"Retry-After": []string{"30"}},
					Body:   io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*redfish.Task{},
			StatusChan:      true,
			EndTime:         startTime.Add(10 * time.Second),
			TaskMonitor: &redfish.TaskMonitorInfo{
				RetryAfter:  startTime.Add(20 * time.Second),
				TaskMonitor: testTaskMonitorUri,
			},
			ExpectErr: context.DeadlineExceeded.Error(),
		},
		"context timeout during task poll": {
			CtxFunc: func(ctx context.Context) (context.Context, func()) {
				return context.WithTimeout(ctx, 10*time.Second)
			},
			Responses: []*http.Response{
				{StatusCode: 202, Body: io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 202,
					Header: http.Header{"Retry-After": []string{"30"}},
					Body:   io.NopCloser(bytes.NewBufferString(""))},
				{StatusCode: 200, Body: io.NopCloser(bytes.NewBufferString(""))},
			},
			ExpectedUpdates: []*redfish.Task{nil},
			StatusChan:      true,
			EndTime:         startTime.Add(10 * time.Second),
			TaskMonitor: &redfish.TaskMonitorInfo{
				RetryAfter:  startTime.Add(2 * time.Second),
				TaskMonitor: testTaskMonitorUri,
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
			ExpectedUpdates: []*redfish.Task{
				{TaskState: "Running", PercentComplete: 0},
				{TaskState: "Running", PercentComplete: 100},
			},
			StatusChan: true,
			EndTime:    startTime.Add(20 * time.Second),
			TaskMonitor: &redfish.TaskMonitorInfo{
				TaskMonitor: testTaskMonitorUri,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			synctest.Test(t, func(t *testing.T) {
				c := &common.TestClient{
					CustomReturnForActions: map[string][]interface{}{},
				}

				for _, resp := range test.Responses {
					c.CustomReturnForActions[http.MethodGet] = append(c.CustomReturnForActions[http.MethodGet], resp)
				}

				var statusChan chan *redfish.Task
				if test.StatusChan {
					statusChan = make(chan *redfish.Task)
					go func() {
						i := 0
						for status := range statusChan {
							require.Less(t, i, len(test.ExpectedUpdates), "test expects at least %d response", i+1)
							expectedStatus := test.ExpectedUpdates[i]
							if expectedStatus != nil {
								expectedStatus.SetClient(c)
							}

							assert.Equal(t, expectedStatus, status,
								"status update [%d], expected [%v], actual [%v]",
								i, expectedStatus, status)
							i++
						}
						assert.Equal(t, len(test.ExpectedUpdates), i)
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
					resp, err = redfish.WaitForTaskMonitor(ctx, c, test.PollRate, test.TaskMonitor, statusChan)
					if statusChan != nil {
						close(statusChan)
					}
					done <- true
				}()

				for {
					synctest.Wait()
					select {
					case <-done:
					default:
						time.Sleep(1 * time.Second) // advance time in the bubble
						continue
					}

					currTime := time.Now().UTC()
					assert.Equal(t, test.EndTime.UTC(), currTime)

					if test.ExpectErr != "" {
						require.ErrorContains(t, err, test.ExpectErr)
					} else {
						require.NoError(t, err)
						assert.Equal(t, test.Responses[len(test.Responses)-1], resp)
					}
					return
				}
			})
		})
	}
}

func TestWaitForTaskMonitorObject(t *testing.T) {
	raw := `{"PowerState": "On"}`
	c := &common.TestClient{
		CustomReturnForActions: map[string][]interface{}{
			http.MethodGet: {
				&http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(raw)),
				},
			},
		},
	}

	expected := &redfish.ComputerSystem{
		PowerState: "On",
		RawData:    []byte(raw),
	}
	expected.SetClient(c)

	object, _, err := redfish.WaitForTaskMonitorObject[redfish.ComputerSystem](
		t.Context(), c, 0, &redfish.TaskMonitorInfo{TaskMonitor: "/Monitor"}, nil,
	)
	require.NoError(t, err)
	assert.Equal(t, expected, object)
}
