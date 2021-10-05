package telnyx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	defaultEndpoint  = "https://api.telnyx.com/v2"
	accountsPath     = "/messaging_profiles"
	messagingPath    = "/messages"
	numberSearchPath = "/available_phone_numbers"
)

type name interface {
}

// RateLimitError is error for 429 http error
type RateLimitError struct {
	Reset time.Time
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("RateLimitError: reset at %v", e.Reset)
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

// Opts are the options to create the client.
type Opts struct {
	// mandatory options.
	AccountID, APIToken string
	//optional
	DefaultEndpoint string
	HTTPClient      *http.Client
	Verbose         bool
}

// Client is main API object
type Client struct {
	accountID, apiToken string
	DefaultEndpoint     string
	httpClient          *http.Client
	verbose             bool
}

// New creates new instances of api
// It returns Client instance. Use it to make API calls.
func New(opts Opts) (*Client, error) {
	if opts.AccountID == "" || opts.APIToken == "" {
		return nil, errors.New("missing auth data")
	}
	endpoint := defaultEndpoint
	if opts.DefaultEndpoint != "" {
		endpoint = opts.DefaultEndpoint
	}
	client := http.DefaultClient
	if opts.HTTPClient != nil {
		client = opts.HTTPClient
	}

	c := &Client{
		accountID:       opts.AccountID,
		apiToken:        opts.APIToken,
		DefaultEndpoint: endpoint,
		httpClient:      client,
		verbose:         opts.Verbose,
	}
	return c, nil
}

func (c *Client) createRequest(ctx context.Context, method, path string) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, c.apiToken))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "go-telnyx/v2")
	return request, nil
}

func (c *Client) checkJSONResponse(response *http.Response, responseBody interface{}) (interface{}, http.Header, error) {
	defer response.Body.Close()
	body := responseBody
	if body == nil {
		body = map[string]interface{}{}
	}
	rawJSON, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode >= 200 && response.StatusCode < 400 {
		if len(rawJSON) > 0 {
			err = json.Unmarshal([]byte(rawJSON), &body)
			if err != nil {
				return nil, nil, err
			}
		}
		return body, response.Header, nil
	}
	if response.StatusCode == 429 {
		reset, _ := strconv.ParseInt(response.Header.Get("X-RateLimit-Reset"), 10, 64)
		return nil, nil, &RateLimitError{Reset: time.Unix(int64((reset/1000)+1), 0)}
	}
	errorBody := ResponseError{}
	if len(rawJSON) > 0 {
		err = json.Unmarshal([]byte(rawJSON), &errorBody)

		if err != nil {
			return nil, nil, err
		}
	}
	message := ""
	if len(errorBody.Errors) > 0 {
		message = errorBody.Errors[0].Detail
		if message == "" {
			message = errorBody.Errors[0].Code
		}
	}

	if message == "" {
		return nil, nil, fmt.Errorf("http code %d", response.StatusCode)
	}
	return nil, nil, errors.New(message)
}

func (c *Client) makeRequest(ctx context.Context, method, path string, data ...interface{}) (interface{}, http.Header, error) {
	request, err := c.createRequest(ctx, method, path)
	var responseBody interface{}
	if err != nil {
		return nil, nil, err
	}
	if len(data) > 0 {
		responseBody = data[0]
	}
	if len(data) > 1 {
		if method == "GET" {
			var item map[string]string
			if data[1] == nil {
				item = make(map[string]string)
			} else {
				var ok bool
				item, ok = data[1].(map[string]string)
				if !ok {
					item = make(map[string]string)
					structType := reflect.TypeOf(data[1]).Elem()
					structValue := reflect.ValueOf(data[1])
					if !structValue.IsNil() {
						structValue = structValue.Elem()
						fieldCount := structType.NumField()
						for i := 0; i < fieldCount; i++ {
							fieldName := structType.Field(i).Name
							fieldValue := structValue.Field(i).Interface()
							if fieldValue == reflect.Zero(structType.Field(i).Type).Interface() {
								//ignore fields with default values
								continue
							}
							item[strings.Replace(strings.ToLower(string(fieldName[0]))+fieldName[1:], "ID", "Id", -1)] = fmt.Sprintf("%v", fieldValue)
						}
					}
				}
			}
			query := make(url.Values)
			for key, value := range item {
				query[key] = []string{value}
			}
			request.URL.RawQuery = query.Encode()
		} else {
			var body []byte
			var err error

			request.Header.Set("Content-Type", "application/json")
			body, err = json.Marshal(data[1])

			if err != nil {
				return nil, nil, err
			}
			request.Body = nopCloser{bytes.NewReader(body)}

		}
	}
	if c.verbose {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, nil, err
		}
		fmt.Printf("%q\n", dump)
	}
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	if c.verbose {
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			return nil, nil, err
		}
		fmt.Printf("%q\n", dump)
	}

	return c.checkJSONResponse(response, responseBody)

}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
