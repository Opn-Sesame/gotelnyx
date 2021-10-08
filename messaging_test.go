package telnyx

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestCreateMessagingProfile(t *testing.T) {
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery: accountsPath,
		Method:       http.MethodPost,
		ResponseHeaders: map[string]string{
			"Location": fmt.Sprintf("https://api.telnyx.com/v2%s", accountsPath),
		},
		ResponseContent: FakeMessagingProfileResponse,
	}})
	defer server.Close()
	numberPoolOpts := NumberPoolSettings{false, 0, true, true, 0}
	data := MessagingProfileParameters{"Test123", true, numberPoolOpts, "https://app.opnsesame.com/"}
	result, err := api.CreateMessagingProfile(context.Background(), &data)
	if err != nil {
		t.Errorf("Failed call of CreateMessagingProfile(): %v", err)
		return
	}
	expect(t, result, "3fa85f64-5717-4562-b3fc-2c963f66afa6")
}
func TestSendMessage(t *testing.T) {

	sendMessageParameters := SendMessageParameters{
		From: "+18445550001",
		To:   "+18665550001",
		Text: "Hello, World!",
	}
	server, api := startMockServer(t, []RequestHandler{{
		PathAndQuery:            messagingPath,
		Method:                  http.MethodPost,
		EstimatedRequestContent: ExpectedSendMessagePayload,
		ResponseContent:         FakeSendMessageResponse}})
	defer server.Close()
	result, err := api.SendMessage(context.Background(), &sendMessageParameters)
	if err != nil {
		t.Errorf("Failed call of SendMessage(): %v", err)
		return
	}
	expect(t, result.Data.To[0].PhoneNumber, sendMessageParameters.To)
	expect(t, result.Data.From.PhoneNumber, sendMessageParameters.From)
	expect(t, result.Data.Text, sendMessageParameters.Text)
}
