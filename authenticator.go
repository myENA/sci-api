package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
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
	username string
	password string

	cas       uint64
	refreshed time.Time

	authMu sync.RWMutex
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
	if request == nil {
		return AuthCAS(atomic.LoadUint64(&pa.cas)), errors.New("request cannot be nil")
	}
	if err := ctx.Err(); err != nil {
		return AuthCAS(atomic.LoadUint64(&pa.cas)), err
	}
	pa.authMu.RLock()
	cas := atomic.LoadUint64(&pa.cas)
	token := pa.token
	cookie := pa.cookie
	if token != "" && cookie != nil {
		request.AddQueryParameter(AccessTokenKey, token)
		request.SetCookies([]*http.Cookie{cookie})
		pa.authMu.RUnlock()
		return AuthCAS(cas), nil
	}
	pa.authMu.RUnlock()
	return AuthCAS(cas), errors.New("token requires refresh")
}

func (pa *PasswordAuthenticator) Refresh(ctx context.Context, client *Client, cas AuthCAS) (AuthCAS, error) {
	if debug {
		log.Printf("[pw-auth-%s] Refresh called", pa.username)
	}
	if client == nil {
		return AuthCAS(atomic.LoadUint64(&pa.cas)), errors.New("client cannot be nil")
	}
	if err := ctx.Err(); err != nil {
		return AuthCAS(atomic.LoadUint64(&pa.cas)), err
	}
	pa.authMu.Lock()
	ccas := atomic.LoadUint64(&pa.cas)
	if ccas < uint64(cas) {
		pa.authMu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value is greater than possible")
	}
	if ccas > uint64(cas) {
		pa.authMu.Unlock()
		return AuthCAS(ccas), nil
	}

	// build new login request
	request := NewRequest("POST", "/users/login", false)
	err := request.SetBodyModel(&UsersLoginPostRequest{
		Username: pa.username,
		Password: pa.password,
	})
	if err != nil {
		pa.token = ""
		pa.cookie = nil
		ncas := atomic.AddUint64(&pa.cas, 1)
		pa.authMu.Unlock()
		return AuthCAS(ncas), err
	}

	// response model
	out := new(UsersLoginPostResponse)

	// execute login request
	httpResponse, _, err := client.Ensure(ctx, request, 200, out)
	if err != nil {
		pa.token = ""
		pa.cookie = nil
		ncas := atomic.AddUint64(&pa.cas, 1)
		pa.authMu.Unlock()
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
		ncas := atomic.AddUint64(&pa.cas, 1)
		pa.authMu.Unlock()
		return AuthCAS(ncas), nil
	}

	// if we were unable to find a cookie, reset state, iterate cas, and return error
	pa.token = ""
	pa.cookie = nil
	ncas := atomic.AddUint64(&pa.cas, 1)
	log.Printf("[pw-auth-%s] Unable to locate cookie \"%s\" in response", pa.username, AccessTokenKey)
	pa.authMu.Unlock()
	return AuthCAS(ncas), fmt.Errorf("unable to locate cookie \"%s\" in response", AccessTokenKey)
}

func (pa *PasswordAuthenticator) Invalidate(ctx context.Context, cas AuthCAS) (AuthCAS, error) {
	if debug {
		log.Printf("[pw-auth-%s] Invalidate called", pa.username)
	}
	if err := ctx.Err(); err != nil {
		return AuthCAS(atomic.LoadUint64(&pa.cas)), err
	}
	pa.authMu.Lock()
	ccas := atomic.LoadUint64(&pa.cas)
	if ccas < uint64(cas) {
		pa.authMu.Unlock()
		return AuthCAS(ccas), errors.New("provided cas value is greater than possible")
	}
	if ccas > uint64(cas) {
		pa.authMu.Unlock()
		return AuthCAS(ccas), nil
	}
	ncas := atomic.AddUint64(&pa.cas, 1)
	pa.token = ""
	pa.cookie = nil
	pa.authMu.Unlock()
	return AuthCAS(ncas), nil
}
