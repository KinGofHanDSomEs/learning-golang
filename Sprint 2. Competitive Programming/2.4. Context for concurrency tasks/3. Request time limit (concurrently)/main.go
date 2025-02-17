package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	URL        string
	Data       string
	StatusCode int
	Err        error
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	var responses []*APIResponse
	ch := make(chan *APIResponse)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	for _, url := range urls {
		go func(u string) {
			resp := &APIResponse{URL: u}
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				resp.Err = err
				ch <- resp
				return
			}
			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				resp.Err = err
				ch <- resp
				return
			}
			defer response.Body.Close()
			resp.StatusCode = response.StatusCode
			body, err := io.ReadAll(response.Body)
			if err != nil {
				resp.Err = err
			} else {
				resp.Data = string(body)
			}
			ch <- resp
		}(url)
	}
	for range urls {
		select {
		case r := <-ch:
			responses = append(responses, r)
		case <-ctx.Done():
			responses = append(responses, &APIResponse{Err: context.DeadlineExceeded})
			return responses
		}
	}
	return responses
}
