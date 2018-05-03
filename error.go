package api

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	// Name seems to be "Error" or "error", not sure if it can be anything else.
	Name string `json:"name"`

	// Status and StatusCode appear to always be the same value.
	Status     int `json:"status"`
	StatusCode int `json:"statusCode"`

	// Code's usefulness is probably limited, as it seems to be missing from quite a few upstream error responses.
	Code string `json:"code,omitempty"`

	// Message, from observation, will sometimes be a string and will sometimes be an object.
	Message json.RawMessage `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Status: %d; Code: %s; Message: %s", e.Status, e.Code, e.Message)
}

func (e APIError) String() string {
	return e.Error()
}
