package api

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	DefaultPort       = 443
	DefaultPathPrefix = "/api"
)

type Config struct {
	Hostname string

	Port       int
	PathPrefix string
}

func DefaultConfig(hostname string) *Config {
	return defaultConfig(hostname)
}

func defaultConfig(hostname string) *Config {
	return &Config{
		Hostname:   hostname,
		Port:       DefaultPort,
		PathPrefix: DefaultPathPrefix,
	}
}

type Client struct {
	config *Config
	client *http.Client

	auth Authenticator
}

func NewClient(conf *Config, authenticator Authenticator, client *http.Client) (*Client, error) {
	if authenticator == nil {
		return nil, errors.New("authenticator cannot be nil")
	}
	def := defaultConfig(conf.Hostname)
	if conf.PathPrefix != "" {
		def.PathPrefix = conf.PathPrefix
	}
	if conf.Port > 0 {
		def.Port = conf.Port
	}

	if client == nil {
		// shamelessly borrowed from https://github.com/hashicorp/go-cleanhttp/blob/master/cleanhttp.go
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     true,
				MaxIdleConnsPerHost:   -1,
			},
		}
	}

	c := &Client{
		config: def,
		client: client,
		auth:   authenticator,
	}

	return c, nil
}

func (c *Client) ClientConfig() Config {
	return *c.config
}

func (c *Client) Do(ctx context.Context, request *Request) (*http.Response, error) {
	_, httpResponse, err := c.do(ctx, request)
	return httpResponse, err
}

func (c *Client) Ensure(ctx context.Context, request *Request, successCode int, out interface{}) (*http.Response, []byte, error) {
	var httpResponse *http.Response
	var cas AuthCAS
	var buff []byte
	var err error

	// perform initial attempt
	if cas, httpResponse, err = c.do(ctx, request); err != nil {
		return nil, nil, err
	}

	// try to read out response body
	buff, err = ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		return httpResponse, nil, err
	}

	// if the status code is 401 and auth was required for this request...
	if httpResponse.StatusCode == http.StatusUnauthorized && request.RequiresToken() {
		// ...invalidate authenticator state
		if _, err = c.auth.Invalidate(ctx, cas); err != nil {
			return nil, nil, err
		}
		// ...attempt request again, this will kick off a Decorate -> Refresh cycle in response to Invalidate call
		if _, httpResponse, err = c.tryDo(ctx, request); err != nil {
			if httpResponse != nil {
				httpResponse.Body.Close()
			}
			return nil, nil, err
		}

		// if retry was successful, read new response body
		buff, err = ioutil.ReadAll(httpResponse.Body)
		httpResponse.Body.Close()
		if err != nil {
			return httpResponse, nil, err
		}
	}

	// if the response code is what was expected, attempt to unmarshal response and return
	if httpResponse.StatusCode == successCode {
		if out != nil {
			err = json.Unmarshal(buff, out)
		}
		return httpResponse, buff, err
	}

	// otherwise, attempt to unmarshal into api error
	// NOTE: We sometimes see an HTML response here despite having "Accept: application/json" header.  This is dumb.
	tmp := make(map[string]json.RawMessage)
	if decodeErr := json.Unmarshal(buff, &tmp); decodeErr == nil {
		// if decoding succeeds, attempt to extract actual error message
		if errData, ok := tmp["error"]; ok {
			apiErr := &APIError{}
			if decodeErr = json.Unmarshal(errData, apiErr); decodeErr == nil {
				err = apiErr
			} else if debug {
				log.Printf("[request-%d] Error decoding error: %s; Data: %s", request.ID(), decodeErr, string(errData))
			}
		} else if debug {
			log.Printf("[request-%d] Unable to locate \"error\" field in response: %+v", request.ID(), tmp)
		}
	} else if debug {
		log.Printf("[request-%d] Error decoding response: %s; Data: %s", request.ID(), decodeErr, string(buff))
	}

	// if we were unable to unmarshal into an error response, just build an error here.
	if err == nil {
		err = &APIError{
			Name:       "error",
			Status:     httpResponse.StatusCode,
			StatusCode: httpResponse.StatusCode,
			Message:    buff,
		}
	}

	return httpResponse, buff, err
}

func (c *Client) tryDo(ctx context.Context, request *Request) (AuthCAS, *http.Response, error) {
	var httpRequest *http.Request
	var httpResponse *http.Response
	var cas AuthCAS
	var err error

	if request.RequiresToken() {
		if cas, err = c.auth.Decorate(ctx, request); err != nil {
			if cas, err = c.auth.Refresh(ctx, c, cas); err != nil {
				return cas, nil, err
			} else if cas, err = c.auth.Decorate(ctx, request); err != nil {
				return cas, nil, err
			}
		}
	}

	if httpRequest, err = request.toHTTP(ctx, c.config); err != nil {
		return cas, nil, err
	}

	httpResponse, err = c.client.Do(httpRequest)
	return cas, httpResponse, err
}

func (c *Client) do(ctx context.Context, request *Request) (AuthCAS, *http.Response, error) {
	if ctx == nil {
		return 0, nil, errors.New("ctx must not be nil")
	}
	var httpResponse *http.Response
	var cas AuthCAS
	var err error
	if cas, httpResponse, err = c.tryDo(ctx, request); err != nil {
		if httpResponse != nil {
			httpResponse.Body.Close()
		}
		return cas, nil, err
	}
	return cas, httpResponse, err
}

type (
	UsersLoginPostRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UsersLoginPostResponse struct {
		ID      string `json:"id"`
		TTL     int    `json:"ttl"`
		Created string `json:"created"`
		UserID  int    `json:"userId"`
	}
)

func (c *Client) UsersLoginPost(ctx context.Context, requestModel *UsersLoginPostRequest) (*http.Response, *UsersLoginPostResponse, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/users/login", false)
	if err = request.SetBodyModel(requestModel); err != nil {
		return nil, nil, err
	}
	out := new(UsersLoginPostResponse)
	httpResponse, _, err := c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

func (c *Client) UsersLogoutPost(ctx context.Context) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	request := NewRequest("POST", "/users/logout", true)
	return c.Ensure(ctx, request, http.StatusNoContent, nil)
}

type (
	ReportGetAPMacListPostResponseDataAPMACsSlice []string

	ReportGetAPMacListPostResponseData struct {
		APMacs ReportGetAPMacListPostResponseDataAPMACsSlice `json:"apMacs"`
	}

	ReportGetAPMacListPostResponse struct {
		Data *ReportGetAPMacListPostResponseData `json:"data"`
	}
)

func (c *Client) ReportGetAPMacListPost(ctx context.Context, id int, query *Query) (*http.Response, *ReportGetAPMacListPostResponse, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/apmac", true)
	request.AddPathParameter("id", strconv.Itoa(id))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := new(ReportGetAPMacListPostResponse)
	httpResponse, _, err := c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

type (
	ReportGetSsidsPostResponseDataSlice []string

	ReportGetSsidsPostResponse struct {
		Data ReportGetSsidsPostResponseDataSlice `json:"data"`
	}
)

func (c *Client) ReportGetSsidsPost(ctx context.Context, id int, query *Query) (*http.Response, *ReportGetSsidsPostResponse, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/facets/ssid", true)
	request.AddPathParameter("id", strconv.Itoa(id))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	out := new(ReportGetSsidsPostResponse)
	httpResponse, _, err := c.Ensure(ctx, request, http.StatusOK, out)
	return httpResponse, out, err
}

// ReportGetDataPost will, on success, always return an object like:
//
//		{
//			"data": [],
//			"metadata": {}
//		}
//
// The exact structure of "data" and "metadata" vary greatly depending upon the report id and section queried, and
// therefore is not modeled here.  It is expected that the caller either doesn't care or knows which specific report
// they're interested in.
func (c *Client) ReportGetDataPost(ctx context.Context, id, sectionId int, query *Query) (*http.Response, []byte, error) {
	if ctx == nil {
		return nil, nil, errors.New("ctx cannot be nil")
	}
	var err error
	request := NewRequest("POST", "/reports/{id}/sections/{sectionId}/data", true)
	request.AddPathParameter("id", strconv.Itoa(id))
	request.AddPathParameter("sectionId", strconv.Itoa(sectionId))
	if err = request.SetBodyModel(query); err != nil {
		return nil, nil, err
	}
	return c.Ensure(ctx, request, http.StatusOK, nil)
}
