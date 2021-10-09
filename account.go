package telnyx

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// New creates new instances of api
// It returns Account instance. Use it to make API calls.
func New(opts Opts) (*Account, error) {
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

	c := &Account{
		accountID:       opts.AccountID,
		apiToken:        opts.APIToken,
		DefaultEndpoint: endpoint,
		httpClient:      client,
		verbose:         opts.Verbose,
	}
	return c, nil
}

func (a *Account) CreateRequest(ctx context.Context, method, path string) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, a.apiToken))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "go-telnyx/v2")
	return request, nil
}

func (a *Account) checkJSONResponse(response *http.Response, responseBody interface{}) (interface{}, http.Header, error) {
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

func (a *Account) makeRequest(ctx context.Context, method, path string, data ...interface{}) (interface{}, http.Header, error) {
	request, err := a.CreateRequest(ctx, method, path)
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
	if a.verbose {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, nil, err
		}
		fmt.Printf("%q\n", dump)
	}
	response, err := a.httpClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	if a.verbose {
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			return nil, nil, err
		}
		fmt.Printf("%q\n", dump)
	}

	return a.checkJSONResponse(response, responseBody)

}
