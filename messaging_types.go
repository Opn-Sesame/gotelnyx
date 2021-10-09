package telnyx

import "time"

// Messaging Profile Structs
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

// SendMessage Structs
type SendMessageParameters struct {
	To                 string   `json:"to"`
	AutoDetect         bool     `json:"auto_detect,omitempty"`
	From               string   `json:"from,omitempty"`
	MediaUrls          []string `json:"media_urls,omitempty"`
	MessagingProfileID string   `json:"messaging_profile_id,omitempty"`
	Subject            string   `json:"subject,omitempty"`
	Text               string   `json:"text,omitempty"`
	Type               string   `json:"type,omitempty"`
	UseProfileWebhooks bool     `json:"use_profile_webhooks,omitempty"`
	WebhookFailoverURL string   `json:"webhook_failover_url,omitempty"`
	WebhookURL         string   `json:"webhook_url,omitempty"`
}
type SendMessageResponseMedia struct {
	ContentType interface{} `json:"content_type"`
	Sha256      interface{} `json:"sha256"`
	Size        interface{} `json:"size"`
	URL         string      `json:"url"`
}
type SendMessageResponseTo struct {
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}
type SendMessageResponseFrom struct {
	Carrier     string `json:"carrier"`
	LineType    string `json:"line_type"`
	PhoneNumber string `json:"phone_number"`
}
type SendMessageResponseData struct {
	CompletedAt        time.Time                  `json:"completed_at"`
	Cost               string                     `json:"cost"`
	Direction          string                     `json:"direction"`
	Encoding           string                     `json:"encoding"`
	Errors             []ResponseError            `json:"errors"`
	From               SendMessageResponseFrom    `json:"from"`
	ID                 string                     `json:"id"`
	Media              []SendMessageResponseMedia `json:"media"`
	MessagingProfileID string                     `json:"messaging_profile_id"`
	OrganizationID     string                     `json:"organization_id"`
	Parts              int                        `json:"parts"`
	ReceivedAt         time.Time                  `json:"received_at"`
	RecordType         string                     `json:"record_type"`
	SentAt             time.Time                  `json:"sent_at"`
	Subject            string                     `json:"subject"`
	Tags               []string                   `json:"tags"`
	Text               string                     `json:"text"`
	To                 []SendMessageResponseTo    `json:"to"`
	Type               string                     `json:"type"`
	ValidUntil         time.Time                  `json:"valid_until"`
	WebhookFailoverURL string                     `json:"webhook_failover_url"`
	WebhookURL         string                     `json:"webhook_url"`
}
type SendMessageResponse struct {
	Data SendMessageResponseData `json:"data"`
}
