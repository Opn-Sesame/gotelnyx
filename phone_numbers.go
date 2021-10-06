package telnyx

import (
	"context"
	"fmt"
	"net/http"
)

// NumberSearch returns a response of available phone numbers.
func (a *Account) NumberSearch(ctx context.Context, searchOpts NumberSearchOptions) (*NumberSearchResponse, error) {
	path := a.DefaultEndpoint + numberSearchPath
	areaCodeFilter := ""
	if searchOpts.AreaCode != "" {
		areaCodeFilter = fmt.Sprintf("&filter[phone_number][starts_with]=%s", searchOpts.AreaCode)
	}
	limit := 1
	if searchOpts.SearchLimit != 0 {
		limit = searchOpts.SearchLimit
	}
	queryParams := fmt.Sprintf(`?filter[limit][]=%v%s`, limit, areaCodeFilter)
	result, _, err := a.makeRequest(ctx, http.MethodGet, path+queryParams, &NumberSearchResponse{})
	if err != nil {
		return nil, err
	}
	return result.(*NumberSearchResponse), nil
}

// CreateNumberOrder creates a phone number order.
func (a *Account) CreateNumberOrder(ctx context.Context, NumberOrderParams NumberOrderParameters) (*NumberOrderResponse, error) {
	path := a.DefaultEndpoint + numberOrderPath
	result, _, err := a.makeRequest(ctx, http.MethodPost, path, &NumberOrderResponse{}, &NumberOrderParams)
	if err != nil {
		return nil, err
	}
	return result.(*NumberOrderResponse), nil
}

// RetreiveNumberOrder creates a phone number order.
func (a *Account) RetreiveNumberOrder(ctx context.Context, retreiveNumberOrderParams RetreiveNumberOrderParameters) (*NumberOrderResponse, error) {
	path := fmt.Sprintf("%s%s/%s", a.DefaultEndpoint, numberOrderPath, retreiveNumberOrderParams.ID)
	result, _, err := a.makeRequest(ctx, http.MethodGet, path, &NumberOrderResponse{}, nil)
	if err != nil {
		return nil, err
	}
	return result.(*NumberOrderResponse), nil
}
