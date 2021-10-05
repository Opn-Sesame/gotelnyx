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
		"toll_free_weight": 10
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
	  "total_results": 100
	}
  }`
