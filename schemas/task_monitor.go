//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TaskMonitorInfo struct {
	TaskMonitor string    // URI to poll for task state
	Task        *Task     // optional representation of the task
	RetryAfter  time.Time // optional time after which we should poll
}

func PostObject[T any,
	PT GenericSchemaObjectPointer[T],
](c Client, uri string, payload any, headers map[string]string) (*T, *TaskMonitorInfo, error) {
	return postObjectWithTask[T, PT](c, uri, payload, headers, false)
}

func PostObjectMultipart[T any,
	PT GenericSchemaObjectPointer[T],
](c Client, uri string, payload map[string]io.Reader, headers map[string]string) (*T, *TaskMonitorInfo, error) {
	return postObjectWithTask[T, PT](c, uri, payload, headers, true)
}

func PostWithTask(c Client, uri string, payload any, headers map[string]string, isMultipart bool) (*http.Response, *TaskMonitorInfo, error) {
	return sendWithTask(c, headers, func() (*http.Response, error) {
		if isMultipart {
			return c.PostMultipartWithHeaders(uri, payload.(map[string]io.Reader), headers)
		}
		return c.PostWithHeaders(uri, payload, headers)
	})
}

// sendWithTask runs a mutating request via send and interprets the result as a
// possibly-asynchronous operation. On a 412 Precondition Failed it drops the
// If-Match header and retries once; some BMC implementations (e.g. Dell
// iDRAC10) reject If-Match on these endpoints. A 202 Accepted response is
// returned as a TaskMonitorInfo; any other success is returned as the raw
// response for the caller to decode.
func sendWithTask(c Client, headers map[string]string, send func() (*http.Response, error)) (*http.Response, *TaskMonitorInfo, error) {
	resp, err := send()

	if err != nil && Is412PreconditionFailed(err) {
		DeferredCleanupHTTPResponse(resp)
		delete(headers, "If-Match")
		resp, err = send()
	}

	if err != nil {
		defer DeferredCleanupHTTPResponse(resp)
		return nil, nil, err
	}

	if resp.StatusCode == http.StatusAccepted {
		defer DeferredCleanupHTTPResponse(resp)
		return nil, ParseTaskMonitorInfo(c, resp), nil
	}
	return resp, nil, nil
}

func postObjectWithTask[T any,
	PT GenericSchemaObjectPointer[T],
](c Client, uri string, payload any, headers map[string]string, isMultipart bool) (*T, *TaskMonitorInfo, error) {
	resp, taskMonitor, err := PostWithTask(c, uri, payload, headers, isMultipart)
	defer DeferredCleanupHTTPResponse(resp)
	if taskMonitor != nil {
		return nil, taskMonitor, err
	}

	entity, err := DecodeGenericEntity[T, PT](c, resp)
	return entity, nil, err
}

func PatchObject[T any,
	PT GenericSchemaObjectPointer[T],
](c Client, uri string, payload any, headers map[string]string) (*T, *TaskMonitorInfo, error) {
	return patchObjectWithTask[T, PT](c, uri, payload, headers)
}

func PatchWithTask(c Client, uri string, payload any, headers map[string]string) (*http.Response, *TaskMonitorInfo, error) {
	return sendWithTask(c, headers, func() (*http.Response, error) {
		return c.PatchWithHeaders(uri, payload, headers)
	})
}

func patchObjectWithTask[T any,
	PT GenericSchemaObjectPointer[T],
](c Client, uri string, payload any, headers map[string]string) (*T, *TaskMonitorInfo, error) {
	resp, taskMonitor, err := PatchWithTask(c, uri, payload, headers)
	defer DeferredCleanupHTTPResponse(resp)
	if taskMonitor != nil {
		return nil, taskMonitor, err
	}

	entity, err := DecodeGenericEntity[T, PT](c, resp)
	return entity, nil, err
}

func ParseTaskMonitorInfo(c Client, resp *http.Response) *TaskMonitorInfo {
	// https://www.dmtf.org/sites/default/files/standards/documents/DSP0266_1.23.0.html#asynchronous-operations
	// When a client issues a request that results in a long-running operation,
	// the service returns the HTTP 202 Accepted status code
	// and a Location header that contains a task monitor URI and,
	// optionally, the Retry-After header that defines the amount of time
	// that the client should wait before querying the status of the operation.
	//
	// The 202 Accepted response __should__ include a response body.
	// If a response body is provided,
	// it __shall__ contain a representation of the Task resource that represents the state of the operation.
	taskMonitorInfo := &TaskMonitorInfo{
		TaskMonitor: resp.Header.Get("Location"),
	}

	retryAfter, err := ParseRetryAfter(resp.Header.Get("Retry-After"))
	if err == nil {
		taskMonitorInfo.RetryAfter = retryAfter
	}

	// task isn't guaranteed to be returned, so mostly ignore the error
	task := &Task{}
	if err := json.NewDecoder(resp.Body).Decode(task); err == nil || task.ODataID != "" {
		task.SetClient(c)
		taskMonitorInfo.Task = task
	}

	return taskMonitorInfo
}

func WaitForTaskMonitor(ctx context.Context, c Client, defaultPollRate time.Duration, taskMonitor *TaskMonitorInfo, taskChan chan<- *Task) (*http.Response, error) {
	if defaultPollRate == 0 {
		defaultPollRate = 10 * time.Second
	}

	if taskMonitor == nil {
		return nil, fmt.Errorf("task monitor is nil")
	} else if taskMonitor.TaskMonitor == "" {
		return nil, fmt.Errorf("task monitor URI is missing")
	}

	if time.Now().Before(taskMonitor.RetryAfter) {
		select {
		case <-time.After(time.Until(taskMonitor.RetryAfter)):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		resp, err := c.Get(taskMonitor.TaskMonitor)
		if err != nil {
			DeferredCleanupHTTPResponse(resp)
			return resp, err
		}

		if resp.StatusCode == http.StatusAccepted {
			if taskChan != nil {
				task := &Task{}
				if err := json.NewDecoder(resp.Body).Decode(task); err == nil || task.ODataID != "" {
					task.SetClient(c)
					taskChan <- task
				} else {
					taskChan <- nil // indicate that we're still getting a response
				}
			}

			waitTime := defaultPollRate
			if retryAfter, err := ParseRetryAfter(resp.Header.Get("Retry-After")); err == nil {
				waitTime = time.Until(retryAfter)
			}

			DeferredCleanupHTTPResponse(resp)

			select {
			case <-time.After(waitTime):
				continue
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		} else {
			// the task monitor is complete, return the raw response
			return resp, nil
		}
	}
}

func WaitForTaskMonitorObject[T any,
	PT GenericSchemaObjectPointer[T],
](ctx context.Context, c Client, defaultPollRate time.Duration, taskMonitor *TaskMonitorInfo, taskChan chan<- *Task) (*T, http.Header, error) {
	resp, err := WaitForTaskMonitor(ctx, c, defaultPollRate, taskMonitor, taskChan)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, nil, err
	}

	headers := resp.Header
	entity, err := DecodeGenericEntity[T, PT](c, resp)
	return entity, headers, err
}
