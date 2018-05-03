package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
)

var requestID uint64

type Request struct {
	mu sync.RWMutex

	id            uint64
	method        string
	uri           string
	requiresToken bool

	queryParameters map[string]string
	pathParameters  map[string]string
	headers         url.Values
	cookies         []*http.Cookie
	body            []byte
}

func NewRequest(method, uri string, requiresToken bool) *Request {
	r := &Request{
		id:              atomic.AddUint64(&requestID, 1),
		method:          strings.ToUpper(method),
		uri:             uri,
		requiresToken:   requiresToken,
		queryParameters: make(map[string]string),
		pathParameters:  make(map[string]string),
		headers:         make(url.Values),
		cookies:         make([]*http.Cookie, 0),
	}

	return r
}

func (r *Request) ID() uint64 {
	return r.id
}

func (r *Request) Method() string {
	return r.method
}

func (r *Request) URI() string {
	return r.uri
}

func (r *Request) RequiresToken() bool {
	return r.requiresToken
}

func (r *Request) SetHeaders(headers url.Values) {
	var l int
	r.mu.Lock()
	r.headers = make(url.Values, len(headers))
	for name, values := range headers {
		l = len(values)
		r.headers[name] = make([]string, l, l)
		copy(r.headers[name], values)
	}
	r.mu.Unlock()
}

func (r *Request) AddHeader(name, value string) {
	r.mu.Lock()
	r.headers.Add(name, value)
	r.mu.Unlock()
}

// SetHeader will attempt to overwrite an existing header with the same name, simply adding it if one is not found.
func (r *Request) SetHeader(name, value string) {
	r.mu.Lock()
	r.headers.Set(name, value)
	r.mu.Unlock()
}

// RemoveHeader will attempt to remove a header from this request, returning the value being removed.
func (r *Request) RemoveHeader(name string) (string, bool) {
	r.mu.Lock()
	current := r.headers.Get(name)
	if current == "" {
		r.mu.Unlock()
		return "", false
	}
	r.headers.Del(name)
	r.mu.Unlock()
	return current, true
}

// Headers will return a copy of current headers on this request
func (r *Request) Headers() url.Values {
	var l int
	r.mu.RLock()
	l = len(r.headers)
	if l == 0 {
		r.mu.RUnlock()
		return nil
	}
	headers := make(url.Values, l)
	for name, values := range r.headers {
		l = len(values)
		headers[name] = make([]string, l, l)
		copy(headers[name], values)
	}
	r.mu.RUnlock()
	return headers
}

func (r *Request) SetCookies(cookies []*http.Cookie) {
	r.mu.Lock()
	l := len(cookies)
	r.cookies = make([]*http.Cookie, l, l)
	copy(r.cookies, cookies)
	r.mu.Unlock()
}

func (r *Request) AddCookie(cookie *http.Cookie) {
	r.mu.Lock()
	r.cookies = append(r.cookies, cookie)
	r.mu.Unlock()
}

// SetCookie will attempt to locate and overwrite a cookie with the same name, simply appending it to the list if one is
// not found
func (r *Request) SetCookie(cookie *http.Cookie) {
	r.mu.Lock()
	for i, cc := range r.cookies {
		if cc.Name == cookie.Name {
			r.cookies[i] = cookie
			r.mu.Unlock()
			return
		}
	}
	r.cookies = append(r.cookies, cookie)
	r.mu.Unlock()
}

func (r *Request) RemoveCookie(name string) (*http.Cookie, bool) {
	r.mu.Lock()
	for _, cookie := range r.cookies {
		if cookie.Name == name {
			r.mu.Unlock()
			return cookie, true
		}
	}
	r.mu.Unlock()
	return nil, false
}

// Cookies will return a copy of the list of cookies to be used with this request
// NOTE: The cookies are pointers.  Be aware.
func (r *Request) Cookies() []*http.Cookie {
	r.mu.RLock()
	l := len(r.cookies)
	if l == 0 {
		r.mu.RUnlock()
		return nil
	}
	cookies := make([]*http.Cookie, l, l)
	copy(cookies, r.cookies)
	r.mu.RUnlock()
	return cookies
}

func (r *Request) SetQueryParameters(params map[string]string) {
	r.mu.Lock()
	r.queryParameters = make(map[string]string, len(params))
	for k, v := range params {
		r.queryParameters[k] = v
	}
	r.mu.Unlock()
}

func (r *Request) AddQueryParameter(param, value string) {
	r.mu.Lock()
	r.queryParameters[param] = value
	r.mu.Unlock()
}

func (r *Request) SetQueryParameter(param, value string) {
	r.mu.Lock()
	r.queryParameters[param] = value
	r.mu.Unlock()
}

func (r *Request) RemoveQueryParameter(param string) {
	r.mu.Lock()
	delete(r.queryParameters, param)
	r.mu.Unlock()
}

func (r *Request) SetFilterQueryParameter(filter *Filter) error {
	if filter == nil {
		r.RemoveQueryParameter("filter")
		return nil
	}
	if b, err := json.Marshal(filter); err != nil {
		return err
	} else {
		r.SetQueryParameter("filter", string(b))
		return nil
	}
}

func (r *Request) QueryParameters() map[string]string {
	r.mu.RLock()
	params := make(map[string]string, len(r.queryParameters))
	for k, v := range r.queryParameters {
		params[k] = v
	}
	r.mu.RUnlock()
	return params
}

func (r *Request) SetPathParameters(params map[string]string) {
	r.mu.Lock()
	r.pathParameters = make(map[string]string, len(params))
	for k, v := range params {
		r.pathParameters[k] = v
	}
	r.mu.Unlock()
}

func (r *Request) SetPathParameter(param, value string) {
	r.mu.Lock()
	r.pathParameters[param] = value
	r.mu.Unlock()
}

func (r *Request) PathParameters() map[string]string {
	r.mu.RLock()
	params := make(map[string]string, len(r.pathParameters))
	for k, v := range r.pathParameters {
		params[k] = v
	}
	r.mu.RUnlock()
	return params
}

func (r *Request) SetBody(body []byte) {
	r.mu.Lock()
	l := len(body)
	if l == 0 {
		r.body = nil
		r.mu.Unlock()
		return
	}
	r.body = make([]byte, l, l)
	copy(r.body, body)
	r.mu.Unlock()
}

func (r *Request) SetBodyModel(model interface{}) error {
	if model == nil {
		r.SetBody(nil)
		return nil
	}
	b, err := json.Marshal(model)
	if err != nil {
		return err
	}
	r.SetBody(b)
	return nil
}

func (r *Request) Body() []byte {
	r.mu.RLock()
	l := len(r.body)
	if l == 0 {
		r.mu.RUnlock()
		return nil
	}
	tmp := make([]byte, l, l)
	copy(tmp, r.body)
	r.mu.RUnlock()
	return tmp
}

func (r *Request) compileURI() string {
	pathParams := r.PathParameters()
	queryParams := r.QueryParameters()

	uri := r.uri

	if len(pathParams) > 0 {
		for k, v := range pathParams {
			uri = strings.Replace(uri, fmt.Sprintf("{%s}", k), v, -1)
		}
	}
	if len(queryParams) > 0 {
		uri = fmt.Sprintf("%s%s", uri, buildQueryParamString(queryParams))
	}
	return uri
}

// toHTTP will attempt to construct an executable http.request
func (r *Request) toHTTP(ctx context.Context, conf *Config) (*http.Request, error) {
	var err error
	var httpRequest *http.Request
	var value string

	body := r.Body()
	bodyLen := len(body)
	compiledURL := fmt.Sprintf("https://%s:%d%s%s", conf.Hostname, conf.Port, conf.PathPrefix, r.compileURI())

	// if debug mode is enabled, prepare a big'ol log statement.
	if debug {
		logMsg := fmt.Sprintf("[request-%d] Preparing request \"%s %s\"", r.id, r.method, compiledURL)

		if bodyLen == 0 {
			logMsg = fmt.Sprintf("%s without body", logMsg)
		} else {
			logMsg = fmt.Sprintf("%s with body: %s", logMsg, string(body))
		}

		log.Print(logMsg)
	}

	if bodyLen == 0 {
		httpRequest, err = http.NewRequest(r.method, compiledURL, nil)
	} else {
		httpRequest, err = http.NewRequest(r.method, compiledURL, bytes.NewBuffer(body))
	}

	if err != nil {
		return nil, err
	}

	for _, cookie := range r.Cookies() {
		httpRequest.AddCookie(cookie)
	}

	for name, values := range r.Headers() {
		for _, value = range values {
			httpRequest.Header.Set(name, value)
		}
	}

	if bodyLen != 0 {
		httpRequest.Header.Set("Content-Type", "application/json")
	}

	httpRequest.Header.Set("Accept", "application/json")

	return httpRequest.WithContext(ctx), nil
}

func buildQueryParamString(queryParams map[string]string) string {
	// TODO: Make more efficient?
	paramParts := make([]string, 0)

	for name, value := range queryParams {
		value = strings.TrimSpace(value)
		if "" != value {
			paramParts = append(paramParts, fmt.Sprintf("%s=%s", name, value))
		}
	}

	if 0 < len(paramParts) {
		return fmt.Sprintf("?%s", strings.Join(paramParts, "&"))
	} else {
		return ""
	}
}
