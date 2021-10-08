package telnyx

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestNumberSearch(t *testing.T) {
	number := "+19705555098"
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery:    numberSearchPath + `?filter[limit][]=1&filter[phone_number][starts_with]=970`,
		Method:          http.MethodGet,
		ResponseContent: FakeNumberSearchResponse}})
	defer server.Close()
	queryParams := NumberSearchOptions{
		AreaCode:    "970",
		SearchLimit: 1,
	}
	result, err := api.NumberSearch(context.Background(), &queryParams)
	if err != nil {
		t.Errorf("Failed call of NumberSearch(): %v", err)
		return
	}
	expect(t, len(result.Data), 1)
	expect(t, result.Data[0].PhoneNumber, number)
}

func TestNumberSearchFail(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery:       numberSearchPath,
		Method:             http.MethodGet,
		ResponseStatusCode: http.StatusBadRequest}})
	defer server.Close()
	queryParams := NumberSearchOptions{
		AreaCode:    "970",
		SearchLimit: 1,
	}
	shouldFail(t, func() (interface{}, error) { return api.NumberSearch(context.Background(), &queryParams) })
}

func TestCreateNumberOrder(t *testing.T) {

	phoneNumbers := []NumberOrderPhoneNumbers{
		{PhoneNumber: "+19705555098"},
	}
	numberOrderParams := NumberOrderParameters{
		PhoneNumbers: &phoneNumbers,
	}
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery:            numberOrderPath,
		Method:                  http.MethodPost,
		EstimatedRequestContent: ExpectedCreateNumberOrderPayload,
		ResponseContent:         FakeNumberOrderResponse}})
	defer server.Close()
	result, err := api.CreateNumberOrder(context.Background(), &numberOrderParams)
	if err != nil {
		t.Errorf("Failed call of NumberOrder(): %v", err)
		return
	}
	expect(t, result.Data.Status, "pending")
}

func TestRetreiveNumberOrder(t *testing.T) {

	retreiveNumberOrderParameters := RetreiveNumberOrderParameters{
		ID: "12ade33a-21c0-473b-b055-b3c836e1c292",
	}
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery:    fmt.Sprintf("%s/%s", numberOrderPath, retreiveNumberOrderParameters.ID),
		Method:          http.MethodGet,
		ResponseContent: FakeNumberOrderResponse}})
	defer server.Close()
	result, err := api.RetreiveNumberOrder(context.Background(), &retreiveNumberOrderParameters)
	if err != nil {
		t.Errorf("Failed call of NumberOrder(): %v", err)
		return
	}
	expect(t, result.Data.ID, "12ade33a-21c0-473b-b055-b3c836e1c292")
}
