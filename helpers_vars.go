package telnyx

var FakeErrorResponse = `{
	"errors":[
	  {
		"code":"10006",
		"title":"Invalid resource ID",
		"detail":"The provided resource ID was invalid",
		"source":{
		  "pointer":"/id"
		},
		"meta":{
		  "url":"https://developers.telnyx.com/docs/api/v2/overview/errors/10006"
		}
	  }
	]
  }`

var FakeMessagingProfileResponse = `{
	"data": {
	  "created_at": "2019-01-23T18:10:02.574Z",
	  "enabled": true,
	  "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
	  "name": "Profile for Messages",
	  "number_pool_settings": {
		"geomatch": false,
		"long_code_weight": 2,
		"skip_unhealthy": false,
		"sticky_sender": true,
		"toll_free_weight": 10.0,
		"tl" :399
	  },
	  "record_type": "messaging_profile",
	  "updated_at": "2019-01-23T18:10:02.574Z",
	  "url_shortener_settings": {
		"domain": "example.ex",
		"prefix": "cmpny",
		"replace_blacklist_only": true,
		"send_webhooks": false
	  },
	  "v1_secret": "rP1VamejkU2v0qIUxntqLW2c",
	  "webhook_api_version": "2",
	  "webhook_failover_url": "https://backup.example.com/hooks",
	  "webhook_url": "https://www.example.com/hooks",
	  "whitelisted_destinations": [
		"US"
	  ]
	}
  }`

var FakeNumberSearchResponse = `{
	"data": [
	  {
		"best_effort": false,
		"cost_information": {
		  "currency": "USD",
		  "monthly_cost": "6.54",
		  "upfront_cost": "3.21"
		},
		"features": [
		  {
			"name": "sms"
		  },
		  {
			"name": "voice"
		  }
		],
		"phone_number": "+19705555098",
		"quickship": true,
		"record_type": "available_phone_number",
		"region_information": [
		  {
			"region_name": "US",
			"region_type": "country_code"
		  }
		],
		"reservable": true,
		"vanity_format": ""
	  }
	],
	"meta": {
	  "best_effort_results": 50,
	  "total_results": 1
	}
  }`

var FakeNumberOrderResponse = `{
	"data": {
	  "billing_group_id": "abc85f64-5717-4562-b3fc-2c9600",
	  "connection_id": "346789098765567",
	  "created_at": "2018-01-01T00:00:00.000000Z",
	  "customer_reference": "MY REF 001",
	  "id": "12ade33a-21c0-473b-b055-b3c836e1c292",
	  "messaging_profile_id": "abc85f64-5717-4562-b3fc-2c9600",
	  "phone_numbers_count": 1,
	  "record_type": "number_order",
	  "requirements_met": true,
	  "status": "pending",
	  "updated_at": "2018-01-01T00:00:00.000000Z",
	  "phone_numbers": [
		{
		  "id": "dc8e4d67-33a0-4cbb-af74-7b58f05bd494",
		  "phone_number": "+19705555098",
		  "record_type": "number_order_phone_number",
		  "regulatory_group_id": "dc8e4d67-33a0-4cbb-af74-7b58f05bd494",
		  "regulatory_requirements": [
			{}
		  ],
		  "requirements_met": true,
		  "status": "pending"
		}
	  ]
	}
  }`

var ExpectedCreateNumberOrderPayload = `{"phone_numbers":[{"phone_number":"+19705555098"}]}`

var FakeSendMessageResponse = `{
	"data": {
	  "completed_at": null,
	  "cost": null,
	  "direction": "outbound",
	  "encoding": "GSM-7",
	  "errors": [],
	  "from": {
		"carrier": "TELNYX LLC",
		"line_type": "VoIP",
		"phone_number": "+18445550001"
	  },
	  "id": "40385f64-5717-4562-b3fc-2c963f66afa6",
	  "media": [
		{
		  "content_type": null,
		  "sha256": null,
		  "size": null,
		  "url": "https://pbs.twimg.com/profile_images/1142168442042118144/AW3F4fFD_400x400.png"
		}
	  ],
	  "messaging_profile_id": "4000eba1-a0c0-4563-9925-b25e842a7cb6",
	  "organization_id": "b448f9cc-a842-4784-98e9-03c1a5872950",
	  "parts": 1,
	  "received_at": "2019-01-23T18:10:02.574Z",
	  "record_type": "message",
	  "sent_at": null,
	  "subject": "From Telnyx!",
	  "tags": [
		"Greetings"
	  ],
	  "text": "Hello, World!",
	  "to": [
		{
		  "carrier": "T-MOBILE USA, INC.",
		  "line_type": "Wireless",
		  "phone_number": "+18665550001",
		  "status": "queued"
		}
	  ],
	  "type": "MMS",
	  "valid_until": null,
	  "webhook_failover_url": "https://backup.example.com/hooks",
	  "webhook_url": "https://www.example.com/hooks"
	}
  }
`
var ExpectedSendMessagePayload = `{"to":"+18665550001","from":"+18445550001","text":"Hello, World!"}`
