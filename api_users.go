package api

import (
	"context"
	"errors"
	"net/http"
)

type Users struct {
	c *Client
}

type (
	UsersLoginPostRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UsersLoginPost200Response struct {
		ID      string `json:"id"`
		TTL     int    `json:"ttl"`
		Created string `json:"created"`
		UserID  int    `json:"userId"`
	}
)

func (u *Users) LoginPost(ctx context.Context, requestModel *UsersLoginPostRequest) (*http.Response, *UsersLoginPost200Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/users/login", false)
	if err = request.SetBodyModel(requestModel); err != nil {
		return nil, nil, err
	}
	out := new(UsersLoginPost200Response)
	httpResponse, _, err := u.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

func (u *Users) LogoutPost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	request := NewRequest("POST", "/users/logout", true)
	return u.c.Ensure(ctx, request, http.StatusNoContent, nil)
}
