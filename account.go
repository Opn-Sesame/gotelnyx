package telnyx

import (
	"context"
	"fmt"
	"net/http"
)

// CreateMessagingProfile creates the Messaging Profile and returns its ID.
func (c *Client) CreateMessagingProfile(ctx context.Context, params MessagingProfileParameters) (string, error) {
	path := c.DefaultEndpoint + accountsPath
	data := params
	result, _, err := c.makeRequest(ctx, http.MethodPost, path, &MessagingProfileResponse{}, data)
	if err != nil {
		return "", err
	}

	id := result.(*MessagingProfileResponse).Data.ID
	if id == "" {
		return "", fmt.Errorf("error retrieving messaging profile id")
	}
	return id, nil
}

// GetNumbers returns the toll-free numbers associated with the site.
func (c *Client) GetNumbers(ctx context.Context, searchOpts NumberSearchOptions) (*NumberSearchResponse, error) {
	path := c.DefaultEndpoint + numberSearchPath
	areaCodeFilter := ""
	if searchOpts.AreaCode != "" {
		areaCodeFilter = fmt.Sprintf("&filter[phone_number][starts_with]=%s", searchOpts.AreaCode)
	}
	limit := 1
	if searchOpts.SearchLimit != 0 {
		limit = searchOpts.SearchLimit
	}
	queryParams := fmt.Sprintf(`?filter[limit][]=%v%s`, limit, areaCodeFilter)
	result, _, err := c.makeRequest(ctx, http.MethodGet, path+queryParams, &NumberSearchResponse{})
	if err != nil {
		return nil, err
	}
	return result.(*NumberSearchResponse), nil
}

// // OrderNumbersByAreaCode purchases n numbers given area-code.
// func (c *Client) OrderNumbersByAreaCode(ctx context.Context, siteID, peerID, areaCode string, n int) (*OrderResponse, error) {
// 	path := c.AccountsEndpoint + "/orders"
// 	req := AreaCodeRequest{
// 		AreaCodeOrder: AreaCodeOrder{
// 			SiteID: siteID,
// 			PeerID: peerID,
// 			AreaCodeSearchAndOrderType: AreaCodeSearchAndOrderType{
// 				Quantity: n,
// 				AreaCode: areaCode,
// 			},
// 		},
// 	}
// 	result, _, err := c.makeRequest(ctx, http.MethodPost, path, &OrderResponse{}, &req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result.(*OrderResponse), nil
// }

// // SearchNumbersByAreaCode finds n numbers given area-code.
// func (c *Client) SearchNumbersByAreaCode(ctx context.Context, areaCode string, n int) (*SearchResult, error) {
// 	path := c.AccountsEndpoint + "/availableNumbers"
// 	params := map[string]string{
// 		"areaCode": areaCode,
// 		"quantity": strconv.Itoa(n),
// 	}
// 	result, _, err := c.makeRequest(ctx, http.MethodGet, path, &SearchResult{}, params)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result.(*SearchResult), nil
// }

// // OrderTollFreeNumbers purchases n numbers given toll-free mask.
// func (c *Client) OrderTollFreeNumbers(ctx context.Context, siteID, peerID, mask string, n int) (*OrderResponse, error) {
// 	path := c.AccountsEndpoint + "/orders"
// 	req := TollFreeOrderRequest{
// 		TollFreeOrder: TollFreeOrder{
// 			SiteID: siteID,
// 			PeerID: peerID,
// 			TollFreeWildCharSearchAndOrderType: TollFreeWildCharSearchAndOrderType{
// 				Quantity:                n,
// 				TollFreeWildCardPattern: mask,
// 			},
// 		},
// 	}
// 	result, _, err := c.makeRequest(ctx, http.MethodPost, path, &OrderResponse{}, &req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result.(*OrderResponse), nil
// }

// // SearchTollFreeNumbers finds n numbers given tollfree mask.
// func (c *Client) SearchTollFreeNumbers(ctx context.Context, mask string, n int) (*SearchResult, error) {
// 	path := c.AccountsEndpoint + "/availableNumbers"
// 	params := map[string]string{
// 		"tollFreeWildCardPattern": mask,
// 		"quantity":                strconv.Itoa(n),
// 	}
// 	result, _, err := c.makeRequest(ctx, http.MethodGet, path, &SearchResult{}, params)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result.(*SearchResult), nil
// }

// // Disconnect disconnects the given phone numbers.
// func (c *Client) Disconnect(ctx context.Context, numbers []string) (*DisconnectTelephoneNumberOrderResponse, error) {
// 	path := c.AccountsEndpoint + "/disconnects"
// 	req := DisconnectTelephoneNumberOrder{
// 		DisconnectTelephoneNumberOrderType: DisconnectTelephoneNumberOrderType{
// 			TelephoneNumberList: TelephoneNumberList{
// 				TelephoneNumber: numbers,
// 			},
// 		},
// 	}
// 	result, _, err := c.makeRequest(ctx, http.MethodPost, path, &DisconnectTelephoneNumberOrderResponse{}, &req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result.(*DisconnectTelephoneNumberOrderResponse), nil
// }

// // GetDisconnect returns information regarding the given disconnect by ID.
// func (c *Client) GetDisconnect(ctx context.Context, id string) (*DisconnectTelephoneNumberOrderResponse, error) {
// 	path := c.AccountsEndpoint + "/disconnects/" + id
// 	result, _, err := c.makeRequest(ctx, http.MethodGet, path, &DisconnectTelephoneNumberOrderResponse{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result.(*DisconnectTelephoneNumberOrderResponse), nil
// }

// // GetOrder returns information regarding the given order.
// func (c *Client) GetOrder(ctx context.Context, id string) (*OrderResponse, error) {
// 	path := c.AccountsEndpoint + "/orders/" + id
// 	result, _, err := c.makeRequest(ctx, http.MethodGet, path, &OrderResponse{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result.(*OrderResponse), nil
// }
