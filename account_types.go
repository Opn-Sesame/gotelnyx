package telnyx

import "time"

type NumberPoolSettings struct {
	Geomatch       bool `json:"geomatch"`
	LongCodeWeight int  `json:"long_code_weight"`
	SkipUnhealthy  bool `json:"skip_unhealthy"`
	StickySender   bool `json:"sticky_sender"`
	TollFreeWeight int  `json:"toll_free_weight"`
}

type MessagingProfileParameters struct {
	ProfileName    string             `json:"name"`
	IsEnabled      bool               `json:"enabled"`
	NumberPoolOpts NumberPoolSettings `json:"number_pool_settings"`
	WebhookUrl     string             `json:"webhook_url"`
}
type URLShortenerSettings struct {
	Domain               string `json:"domain"`
	Prefix               string `json:"prefix"`
	ReplaceBlacklistOnly bool   `json:"replace_blacklist_only"`
	SendWebhooks         bool   `json:"send_webhooks"`
}
type MessagingProfileData struct {
	CreatedAt               time.Time            `json:"created_at"`
	Enabled                 bool                 `json:"enabled"`
	ID                      string               `json:"id"`
	Name                    string               `json:"name"`
	NumberPoolSettings      NumberPoolSettings   `json:"number_pool_settings"`
	RecordType              string               `json:"record_type"`
	UpdatedAt               time.Time            `json:"updated_at"`
	URLShortenerSettings    URLShortenerSettings `json:"url_shortener_settings"`
	V1Secret                string               `json:"v1_secret"`
	WebhookAPIVersion       string               `json:"webhook_api_version"`
	WebhookFailoverURL      string               `json:"webhook_failover_url"`
	WebhookURL              string               `json:"webhook_url"`
	WhitelistedDestinations []string             `json:"whitelisted_destinations"`
}
type MessagingProfileResponse struct {
	Data MessagingProfileData `json:"data"`
}
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

type NumberSearchOptions struct {
	AreaCode    string
	SearchLimit int
}

type NumberSearchResponseCostInfo struct {
	Currency    string `json:"currency"`
	MonthlyCost string `json:"monthly_cost"`
	UpfrontCost string `json:"upfront_cost"`
}
type NumberSearchResponseFeatures struct {
	Name string `json:"name"`
}
type NumberSearchResponseRegionInfo struct {
	RegionName string `json:"region_name"`
	RegionType string `json:"region_type"`
}
type NumberSearchResponseData struct {
	BestEffort        bool                             `json:"best_effort"`
	CostInformation   NumberSearchResponseCostInfo     `json:"cost_information"`
	Features          []NumberSearchResponseFeatures   `json:"features"`
	PhoneNumber       string                           `json:"phone_number"`
	Quickship         bool                             `json:"quickship"`
	RecordType        string                           `json:"record_type"`
	RegionInformation []NumberSearchResponseRegionInfo `json:"region_information"`
	Reservable        bool                             `json:"reservable"`
	VanityFormat      string                           `json:"vanity_format"`
}

type NumberSearchResponseMeta struct {
	BestEffortResults int `json:"best_effort_results"`
	TotalResults      int `json:"total_results"`
}
type NumberSearchResponse struct {
	Data []NumberSearchResponseData `json:"data"`
	Meta NumberSearchResponseMeta   `json:"meta"`
}
