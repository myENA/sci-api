package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	AccessTokenKey = "access_token"
)

type (
	AuthCAS uint64

	Authenticator interface {
		Decorate(context.Context, *Request) (AuthCAS, error)
		Refresh(context.Context, *Client, AuthCAS) (AuthCAS, error)
		Invalidate(context.Context, AuthCAS) (AuthCAS, error)
	}
)

type PasswordAuthenticator struct {
	mu sync.RWMutex

	username string
	password string

	cas       uint64
	refreshed time.Time

	token  string
	cookie *http.Cookie
}

func NewPasswordAuthenticator(username, password string) *PasswordAuthenticator {
	pa := &PasswordAuthenticator{
		username: username,
		password: password,
	}
	return pa
}

func (pa *PasswordAuthenticator) Username() string {
	return pa.username
}

func (pa *PasswordAuthenticator) Password() string {
	return pa.password
}

func (pa *PasswordAuthenticator) Decorate(ctx context.Context, request *Request) (AuthCAS, error) {
	if debug {
		log.Printf("[pw-auth-%s] Decorate called for request %d", pa.username, request.ID())
	}

	pa.mu.RLock()
	cas := pa.cas

	if request == nil {
		pa.mu.RUnlock()
		return AuthCAS(cas), errors.New("request cannot be nil")
	}
	if err := ctx.Err(); err != nil {
		pa.mu.RUnlock()
		return AuthCAS(cas), err
	}

	token := pa.token
	cookie := pa.cookie

	if token != "" && cookie != nil {
		request.AddQueryParameter(AccessTokenKey, token)
		request.SetCookies([]*http.Cookie{cookie})
		pa.mu.RUnlock()
		return AuthCAS(cas), nil
	}

	pa.mu.RUnlock()
	return AuthCAS(cas), errors.New("token requires refresh")
}

func (pa *PasswordAuthenticator) Refresh(ctx context.Context, client *Client, cas AuthCAS) (AuthCAS, error) {
	if debug {
		log.Printf("[pw-auth-%s] Refresh called", pa.username)
	}

	pa.mu.Lock()
	ccas := pa.cas

	if client == nil {
		pa.mu.Unlock()
		return AuthCAS(ccas), errors.New("client cannot be nil")
	}
	if err := ctx.Err(); err != nil {
		pa.mu.Unlock()
		return AuthCAS(ccas), err
	}
	if ccas < uint64(cas) {
		pa.mu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value is greater than possible")
	}
	if ccas > uint64(cas) {
		pa.mu.Unlock()
		return AuthCAS(ccas), nil
	}

	// build new login request
	request := NewRequest("POST", "/users/login", false)
	err := request.SetBodyModel(&UsersLoginRequest{Username: pa.username, Password: pa.password})
	if err != nil {
		pa.token = ""
		pa.cookie = nil
		pa.cas++
		ncas := pa.cas
		pa.mu.Unlock()
		return AuthCAS(ncas), err
	}

	// response model
	out := new(UsersLoginResponse)

	// execute login request
	httpResponse, _, err := client.Ensure(ctx, request, 200, out)
	if err != nil {
		pa.token = ""
		pa.cookie = nil
		pa.cas++
		ncas := pa.cas
		pa.mu.Unlock()
		return AuthCAS(ncas), err
	}

	// attempt to locate appropriate cookie...
	if cookie := TryExtractAccessTokenCookie(httpResponse); cookie != nil {
		// cookie found, update state and cas
		pa.token = out.ID
		pa.cookie = cookie
		if debug {
			log.Printf("[pw-auth-%s] Token %s; Cookie: %+v", pa.username, pa.token, pa.cookie)
		}
		// update internal cas
		pa.cas++
		ncas := pa.cas
		pa.mu.Unlock()
		return AuthCAS(ncas), nil
	}

	// if we were unable to find a cookie, reset state, iterate cas, and return error
	pa.token = ""
	pa.cookie = nil
	pa.cas++
	ncas := pa.cas
	log.Printf("[pw-auth-%s] Unable to locate cookie \"%s\" in response", pa.username, AccessTokenKey)
	pa.mu.Unlock()
	return AuthCAS(ncas), fmt.Errorf("unable to locate cookie \"%s\" in response", AccessTokenKey)
}

func (pa *PasswordAuthenticator) Invalidate(ctx context.Context, cas AuthCAS) (AuthCAS, error) {
	if debug {
		log.Printf("[pw-auth-%s] Invalidate called", pa.username)
	}

	pa.mu.Lock()
	ccas := pa.cas

	if err := ctx.Err(); err != nil {
		pa.mu.Unlock()
		return AuthCAS(ccas), err
	}
	if ccas < uint64(cas) {
		pa.mu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value is greater than possible")
	}
	if ccas > uint64(cas) {
		pa.mu.Unlock()
		return AuthCAS(ccas), nil
	}

	pa.cas++
	ncas := pa.cas
	pa.token = ""
	pa.cookie = nil
	pa.mu.Unlock()
	return AuthCAS(ncas), nil
}
