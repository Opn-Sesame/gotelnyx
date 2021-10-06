package telnyx

import (
	"net/http"
)

// Opts are the options to create the account.
type Opts struct {
	// mandatory options.
	AccountID, APIToken string
	//optional
	DefaultEndpoint string
	HTTPClient      *http.Client
	Verbose         bool
}

// Account is main API object
type Account struct {
	accountID, apiToken string
	DefaultEndpoint     string
	httpClient          *http.Client
	verbose             bool
}

// Telnyx error object
type ResponseErrorDetails struct {
	Code   string                 `json:"code"`
	Title  string                 `json:"title"`
	Detail string                 `json:"detail"`
	Source map[string]interface{} `json:"source"`
	Meta   map[string]interface{} `json:"meta"`
}
type ResponseError struct {
	Errors []ResponseErrorDetails `json:"errors"`
}
