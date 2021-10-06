package telnyx

import "time"

// Webhook Structs
type DeliveryUpdateWebhookFrom struct {
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
}
type DeliveryUpdateWebhookTo struct {
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}
type DeliveryUpdateWebhookMedia struct {
	ContentType interface{} `json:"content_type"`
	Sha256      interface{} `json:"sha256"`
	Size        interface{} `json:"size"`
	URL         string      `json:"url"`
}
type DeliveryUpdateWebhookPayload struct {
	CompletedAt        interface{}                  `json:"completed_at"`
	Cost               interface{}                  `json:"cost"`
	Direction          string                       `json:"direction"`
	Encoding           string                       `json:"encoding"`
	Errors             []interface{}                `json:"errors"`
	From               DeliveryUpdateWebhookFrom    `json:"from"`
	ID                 string                       `json:"id"`
	Media              []DeliveryUpdateWebhookMedia `json:"media"`
	MessagingProfileID string                       `json:"messaging_profile_id"`
	OrganizationID     string                       `json:"organization_id"`
	Parts              int                          `json:"parts"`
	ReceivedAt         time.Time                    `json:"received_at"`
	RecordType         string                       `json:"record_type"`
	SentAt             interface{}                  `json:"sent_at"`
	Subject            string                       `json:"subject"`
	Tags               []string                     `json:"tags"`
	Text               string                       `json:"text"`
	To                 []DeliveryUpdateWebhookTo    `json:"to"`
	Type               string                       `json:"type"`
	ValidUntil         interface{}                  `json:"valid_until"`
	WebhookFailoverURL string                       `json:"webhook_failover_url"`
	WebhookURL         string                       `json:"webhook_url"`
}
type DeliveryUpdateWebhookData struct {
	EventType  string                       `json:"event_type"`
	ID         string                       `json:"id"`
	OccurredAt string                       `json:"occurred_at"`
	Payload    DeliveryUpdateWebhookPayload `json:"payload"`
	RecordType string                       `json:"record_type"`
}
type DeliveryUpdateWebhookMeta struct {
	Attempt     int    `json:"attempt"`
	DeliveredTo string `json:"delivered_to"`
}
type DeliveryUpdateWebhook struct {
	Data MessagingProfileData      `json:"data"`
	Meta DeliveryUpdateWebhookMeta `json:"meta"`
}

type InboundMessageWebhookToAndFrom struct {
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

type InboundMessageWebhookPayload struct {
	CompletedAt        interface{}                      `json:"completed_at"`
	Cost               interface{}                      `json:"cost"`
	Direction          string                           `json:"direction"`
	Encoding           string                           `json:"encoding"`
	Errors             []interface{}                    `json:"errors"`
	From               InboundMessageWebhookToAndFrom   `json:"from"`
	ID                 string                           `json:"id"`
	Media              []interface{}                    `json:"media"`
	MessagingProfileID string                           `json:"messaging_profile_id"`
	OrganizationID     string                           `json:"organization_id"`
	Parts              int                              `json:"parts"`
	ReceivedAt         time.Time                        `json:"received_at"`
	RecordType         string                           `json:"record_type"`
	SentAt             interface{}                      `json:"sent_at"`
	Subject            string                           `json:"subject"`
	Tags               []string                         `json:"tags"`
	Text               string                           `json:"text"`
	To                 []InboundMessageWebhookToAndFrom `json:"to"`
	Type               string                           `json:"type"`
	ValidUntil         interface{}                      `json:"valid_until"`
	WebhookFailoverURL string                           `json:"webhook_failover_url"`
	WebhookURL         string                           `json:"webhook_url"`
}
type InboundMessageWebhookData struct {
	EventType  string                       `json:"event_type"`
	ID         string                       `json:"id"`
	OccurredAt string                       `json:"occurred_at"`
	Payload    InboundMessageWebhookPayload `json:"payload"`
	RecordType string                       `json:"record_type"`
}

type InboundMessageWebhook struct {
	Data InboundMessageWebhookData `json:"data"`
}

type ReplacedClickWebhookData struct {
	MessageID   string `json:"message_id"`
	RecordType  string `json:"record_type"`
	TimeClicked string `json:"time_clicked"`
	To          string `json:"to"`
	URL         string `json:"url"`
}
type ReplacedClickWebhook struct {
	Data ReplacedClickWebhookData `json:"data"`
}

type NumberOrderStatusUpdateWebhookPayload struct {
	BillingGroupID     string    `json:"billing_group_id"`
	ConnectionID       string    `json:"connection_id"`
	CreatedAt          time.Time `json:"created_at"`
	CustomerReference  string    `json:"customer_reference"`
	ID                 string    `json:"id"`
	MessagingProfileID string    `json:"messaging_profile_id"`
	PhoneNumbersCount  int       `json:"phone_numbers_count"`
	RecordType         string    `json:"record_type"`
	RequirementsMet    bool      `json:"requirements_met"`
	Status             string    `json:"status"`
	UpdatedAt          time.Time `json:"updated_at"`
}
type NumberOrderStatusUpdateWebhookData struct {
	EventType  string                                `json:"event_type"`
	ID         string                                `json:"id"`
	OccurredAt string                                `json:"occurred_at"`
	Payload    NumberOrderStatusUpdateWebhookPayload `json:"payload"`
	RecordType string                                `json:"record_type"`
}
type NumberOrderStatusUpdateWebhook struct {
	Data NumberOrderStatusUpdateWebhookData `json:"data"`
}
