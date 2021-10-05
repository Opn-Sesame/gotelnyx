package telnyx

import (
	"context"
	"fmt"
	"net/http"
	"net/textproto"
	"testing"
)

func TestNew(t *testing.T) {
	api, _ := New(Opts{AccountID: testAccountID, APIToken: "YXBpVG9rZW46YXBpU2VjcmV0"})
	expect(t, api.accountID, testAccountID)
	expect(t, api.apiToken, "YXBpVG9rZW46YXBpU2VjcmV0")
	expect(t, api.DefaultEndpoint+accountsPath, "https://api.telnyx.com/v2/messaging_profiles")
	expect(t, api.DefaultEndpoint+messagingPath, "https://api.telnyx.com/v2/messages")
}

func TestNewFail(t *testing.T) {
	shouldFail(t, func() (interface{}, error) {
		return New(Opts{APIToken: "YXBpVG9rZW46YXBpU2VjcmV0"})
	})
	shouldFail(t, func() (interface{}, error) {
		return New(Opts{AccountID: testAccountID})
	})
}

func TestCreateRequest(t *testing.T) {
	endpoint := "https://localhost"
	api := getAPI(endpoint)
	req, err := api.createRequest(context.Background(), http.MethodGet, endpoint+"/v2/test")
	if err != nil {
		t.Fatal(err)
	}
	expect(t, req.URL.String(), endpoint+"/v2/test")
	expect(t, req.Method, http.MethodGet)
	expect(t, req.Header.Get("Accept"), "application/json")
	expect(t, req.Header.Get("User-Agent"), "go-telnyx/v2")
	expect(t, req.Header.Get("Authorization"), "Bearer YXBpVG9rZW46YXBpU2VjcmV0")
}

func TestCreateRequestFail(t *testing.T) {
	endpoint := "https://localhost"
	api := getAPI(endpoint)
	shouldFail(t, func() (interface{}, error) {
		return api.createRequest(context.Background(), "invalid\n\r\tmethod", "invalid:/\n/ url = ")
	})
}

func TestCheckJSONResponse(t *testing.T) {
	type Test struct {
		Test string `json:"test"`
	}
	endpoint := "https://localhost"
	api := getAPI(endpoint)
	data, _, _ := api.checkJSONResponse(createFakeResponse(`{"test": "test"}`, 200), map[string]interface{}{})
	result := data.(map[string]interface{})
	expect(t, result["test"].(string), "test")
	data, _, _ = api.checkJSONResponse(createFakeResponse(`{"test": "test"}`, 200), nil)
	result = data.(map[string]interface{})
	expect(t, result["test"].(string), "test")
	data, _, _ = api.checkJSONResponse(createFakeResponse(`{"test": "test"}`, 200), &Test{})
	testResult := data.(*Test)
	expect(t, testResult.Test, "test")
}

func TestCheckResponseFail(t *testing.T) {
	endpoint := "https://localhost"
	api := getAPI(endpoint)
	fail := func(action func() (interface{}, http.Header, error)) error {
		_, _, err := action()
		if err == nil {
			t.Error("Should fail here")
		}
		return err
	}
	err := fail(func() (interface{}, http.Header, error) {
		return api.checkJSONResponse(createFakeResponse(FakeErrorResponse, 400), nil)
	})
	expect(t, err.Error(), "The provided resource ID was invalid")
	err = fail(func() (interface{}, http.Header, error) {
		return api.checkJSONResponse(createFakeResponse(`{"errors":[{"code": "400"}]}`, 400), nil)
	})
	expect(t, err.Error(), "400")
	err = fail(func() (interface{}, http.Header, error) {
		return api.checkJSONResponse(createFakeResponse(`{"errors":[]}`, 400), nil)
	})
	expect(t, err.Error(), "http code 400")
	fail(func() (interface{}, http.Header, error) {
		return api.checkJSONResponse(createFakeResponse("invalid\njson", 400), nil)
	})
	err = fail(func() (interface{}, http.Header, error) {
		resp := createFakeResponse("", 429)
		resp.Header = map[string][]string{textproto.CanonicalMIMEHeaderKey("X-RateLimit-Reset"): {"1479308598680"}}
		return api.checkJSONResponse(resp, nil)
	})
	e := err.(*RateLimitError)
	expect(t, e.Reset.Unix(), int64(1479308599))
}

func TestMakeRequest(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery:  messagingPath,
		ContentToSend: `{"test": "test"}`}})
	defer server.Close()
	result, _, _ := api.makeRequest(context.Background(), http.MethodGet, api.DefaultEndpoint+messagingPath, map[string]interface{}{})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithQuery(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery:  fmt.Sprintf("%s?field1=value1&field2=value+with+space", messagingPath),
		ContentToSend: `{"test": "test"}`}})
	defer server.Close()
	result, _, _ := api.makeRequest(context.Background(), http.MethodGet, api.DefaultEndpoint+messagingPath, nil, map[string]string{
		"field1": "value1",
		"field2": "value with space"})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithBody(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery:     messagingPath,
		Method:           http.MethodPost,
		EstimatedHeaders: map[string]string{"Content-Type": "application/json"},
		EstimatedContent: `{"field1":"value1","field2":"value with space"}`,
		ContentToSend:    `{"test": "test"}`}})
	defer server.Close()
	result, _, _ := api.makeRequest(context.Background(), http.MethodPost, api.DefaultEndpoint+messagingPath, nil, map[string]interface{}{
		"field1": "value1",
		"field2": "value with space"})
	expect(t, result.(map[string]interface{})["test"], "test")
}

func TestMakeRequestWithEmptyResponse(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery:  messagingPath,
		Method:        http.MethodGet,
		ContentToSend: ""}})
	defer server.Close()
	result, _, _ := api.makeRequest(context.Background(), http.MethodGet, api.DefaultEndpoint+messagingPath, &[]interface{}{})
	expect(t, len(*result.(*[]interface{})), 0)
}
