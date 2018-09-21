package api

import (
	"context"
	"errors"
	"net/http"
)

type Systems struct {
	c *Client
}

type (
	SystemsCount200Response struct {
		Count int `json:"count"`
	}
)

func (s Systems) Count(ctx context.Context) (*http.Response, *SystemsCount200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	out := new(SystemsCount200Response)
	request := NewRequest("GET", "/systems/count", true)
	httpResponse, _, err := s.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}
