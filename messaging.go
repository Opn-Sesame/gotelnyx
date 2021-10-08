package telnyx

import (
	"context"
	"fmt"
	"net/http"
)

// CreateMessagingProfile creates the Messaging Profile and returns its ID.
func (a *Account) CreateMessagingProfile(ctx context.Context, params *MessagingProfileParameters) (string, error) {
	path := a.DefaultEndpoint + accountsPath
	data := params
	result, _, err := a.makeRequest(ctx, http.MethodPost, path, &MessagingProfileResponse{}, data)
	if err != nil {
		return "", err
	}

	id := result.(*MessagingProfileResponse).Data.ID
	if id == "" {
		return "", fmt.Errorf("error retrieving messaging profile id")
	}
	return id, nil
}

// SendMessage allows you to send a message with any messaging resource. Current messaging resources include: long-code, short-code, number-pool, and alphanumeric-sender-id.
func (a *Account) SendMessage(ctx context.Context, params *SendMessageParameters) (*SendMessageResponse, error) {
	path := a.DefaultEndpoint + messagingPath
	if params.From == "" && params.MessagingProfileID == "" {
		return nil, fmt.Errorf("error: from number or number pool not found")
	}
	result, _, err := a.makeRequest(ctx, http.MethodPost, path, &SendMessageResponse{}, params)
	if err != nil {
		return nil, err
	}

	id := result.(*SendMessageResponse).Data.ID
	if id == "" {
		return nil, fmt.Errorf("error retrieving message id")
	}
	return result.(*SendMessageResponse), nil
}
