package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Data       string
	StatusCode int
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, context.DeadlineExceeded
		}
		return nil, err
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return &APIResponse{
		Data:       string(body),
		StatusCode: r.StatusCode,
	}, nil
}
