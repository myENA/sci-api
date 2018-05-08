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
	UsersLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UsersLoginResponse struct {
		ID      string `json:"id"`
		TTL     int    `json:"ttl"`
		Created string `json:"created"`
		UserID  int    `json:"userId"`
	}
)

func (u *Users) Login(ctx context.Context, requestModel *UsersLoginRequest) (*http.Response, *UsersLoginResponse, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/users/login", false)
	if err = request.SetBodyModel(requestModel); err != nil {
		return nil, nil, err
	}
	out := new(UsersLoginResponse)
	httpResponse, _, err := u.c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

func (u *Users) Logout(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	request := NewRequest("POST", "/users/logout", true)
	return u.c.Ensure(ctx, request, http.StatusNoContent, nil)
}
