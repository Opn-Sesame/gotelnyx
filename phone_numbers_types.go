package telnyx

import "time"

// Number Search Structs
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

type NumberSearchResponseList struct {
	PhoneNumbers []string
}

// Number Order Structs
type NumberOrderResponsePhoneNumbers struct {
	ID                     string `json:"id"`
	PhoneNumber            string `json:"phone_number"`
	RecordType             string `json:"record_type"`
	RegulatoryGroupID      string `json:"regulatory_group_id"`
	RegulatoryRequirements []struct {
	} `json:"regulatory_requirements"`
	RequirementsMet bool   `json:"requirements_met"`
	Status          string `json:"status"`
}
type NumberOrderResponseData struct {
	BillingGroupID     string                            `json:"billing_group_id"`
	ConnectionID       string                            `json:"connection_id"`
	CreatedAt          time.Time                         `json:"created_at"`
	CustomerReference  string                            `json:"customer_reference"`
	ID                 string                            `json:"id"`
	MessagingProfileID string                            `json:"messaging_profile_id"`
	PhoneNumbersCount  int                               `json:"phone_numbers_count"`
	RecordType         string                            `json:"record_type"`
	RequirementsMet    bool                              `json:"requirements_met"`
	Status             string                            `json:"status"`
	UpdatedAt          time.Time                         `json:"updated_at"`
	PhoneNumbers       []NumberOrderResponsePhoneNumbers `json:"phone_numbers"`
}
type NumberOrderResponse struct {
	Data NumberOrderResponseData `json:"data"`
}

type NumberOrderPhoneNumbers struct {
	PhoneNumber            string `json:"phone_number"`
	RegulatoryRequirements []*struct {
	} `json:"regulatory_requirements,omitempty"`
}
type NumberOrderParameters struct {
	BillingGroupID     string                     `json:"billing_group_id,omitempty"`
	ConnectionID       string                     `json:"connection_id,omitempty"`
	CustomerReference  string                     `json:"customer_reference,omitempty"`
	MessagingProfileID string                     `json:"messaging_profile_id,omitempty"`
	PhoneNumbers       *[]NumberOrderPhoneNumbers `json:"phone_numbers,omitempty"`
}

type RetreiveNumberOrderParameters struct {
	ID string `json:"number_order_id"`
}
