package zincsearch

// ZincSaerchErrorReponse is the response of ZincSearch when an error occurs
type ZincSearchErrorReponse struct {
	Error string `json:"error"`
}

// IndexDocumentsResponse is the response of ZincSearch when a document is created
type IndexDocumentsResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}

// SearchDocumentsRsponse is the response of ZincSearch when search for documents
type SearchDocumentsRsponse struct {
	Took     int         `json:"took"`
	TimedOut bool        `json:"timed_out"`
	MaxScore float64     `json:"max_score"`
	Hits     Hits        `json:"hits"`
	Buckets  interface{} `json:"buckets"`
	Error    string      `json:"error"`
}

// Simplified SearchResponse structure
type Hits struct {
	Total Total `json:"total"`
	Hits  []Hit `json:"hits"`
}

type Hit struct {
	Index     string                 `json:"_index"`
	Type      string                 `json:"_type"`
	ID        string                 `json:"_id"`
	Score     float64                `json:"_score"`
	Timestamp string                 `json:"@timestamp"`
	Source    map[string]interface{} `json:"_source"`
}

type Total struct {
	Value int `json:"value"`
}

//ZincSearch search response
// {
// 	"took": 73,
// 	"timed_out": false,
// 	"max_score": 0,
// 	"hits": {
// 			"total": {
// 					"value": 132
// 			},
// 			"hits": [
// 					{
// 							"_index": "emails",
// 							"_type": "_doc",
// 							"_id": "28eNnVcUlTZ",
// 							"_score": 7.220565265795022,
// 							"@timestamp": "2024-05-13T21:04:56.670025472Z",
// 							"_source": {
// 									"@timestamp": "2024-05-13T21:04:56.670025472Z",
// 									"bcc": [
// 											"psellers@pacbell.net",
// 											"psellers@haas.berkeley.edu"
// 									],
// 									"body": "My calendar shows free, but you'll have to check with ms prentice.\n\n\n\n\tNancy Sellers <Nancy.Sellers@RobertMondavi.com>\n\t09/25/2000 10:49 AM\n\t\t \n\t\t To: \"'Jeff.Dasovich@enron.com'\" <Jeff.Dasovich@enron.com>\n\t\t cc: \"'Prentice @ Berkeley'\" <PSellers@haas.berkeley.edu>, \"'Prentice \nSellers'\" <PSellers@pacbell.net>\n\t\t Subject: RE: Dome Driveway\n\nHi there - long time no see!!  I honestly do need to know if you and PP want \nto join us for the October 7 dinner cause I honestly am ordering the foie \ngras and will need to get two if you will be with us (cuz you eat so much!)  \nI know that Scotty and Cameron will be in Europe\n-----Original Message----- \nFrom:   Jeff.Dasovich@enron.com [SMTP:Jeff.Dasovich@enron.com] \nSent:   Monday, September 25, 2000 8:36 AM \nTo:     eldon@interx.net; Nancy.Sellers@RobertMondavi.com \nSubject:        Dome Driveway \n\nGreetings (or Hola!): \nTalked to Jeff Green.  He said he'd do the road, but can't do it with the \nwood that's in the driveway.  Eldon, are you and Bob planning to go up this \nweek?  If so, any chance of getting the wood moved?  Otherwise, going to be \ntough to do the road.  Jeff starts on my road on Wednesday and I'll be \ngoing up this weekend to monitor (and start organizing some of the chopped \nup wood that's lying around the yard at the dome. \nWelcome back.  Hope you had a good time in Mexico. \nBest, \nJeff \n",
// 									"cc": [
// 											"psellers@pacbell.net",
// 											"psellers@haas.berkeley.edu"
// 									],
// 									"content_type": "text/plain; charset=us-ascii",
// 									"date": "2000-09-25T03:55:00-07:00",
// 									"from": "jeff.dasovich@enron.com",
// 									"message_id": "<30873727.1075843192679.JavaMail.evans@thyme>",
// 									"subject": "RE: Dome Driveway",
// 									"to": [
// 											"nancy.sellers@robertmondavi.com"
// 									]
// 							}
// 					},
// 					{
// 							"_index": "emails",
// 							"_type": "_doc",
// 							"_id": "28eNnVcCKLZ",
// 							"_score": 7.863911737922981,
// 							"@timestamp": "2024-05-13T21:04:56.669403648Z",
// 							"_source": {
// 									"@timestamp": "2024-05-13T21:04:56.669403648Z",
// 									"bcc": null,
// 									"body": "Greetings (or Hola!):\n\nTalked to Jeff Green.  He said he'd do the road, but can't do it with the \nwood that's in the driveway.  Eldon, are you and Bob planning to go up this \nweek?  If so, any chance of getting the wood moved?  Otherwise, going to be \ntough to do the road.  Jeff starts on my road on Wednesday and I'll be going \nup this weekend to monitor (and start organizing some of the chopped up wood \nthat's lying around the yard at the dome.\n\nWelcome back.  Hope you had a good time in Mexico.  \n\nBest,\nJeff",
// 									"cc": null,
// 									"content_type": "text/plain; charset=us-ascii",
// 									"date": "2000-09-25T03:35:00-07:00",
// 									"from": "jeff.dasovich@enron.com",
// 									"message_id": "<9748533.1075843192656.JavaMail.evans@thyme>",
// 									"subject": "Dome Driveway",
// 									"to": [
// 											"eldon@interx.net",
// 											"nancy.sellers@robertmondavi.com"
// 									]
// 							}
// 					}
// 			]
// 	},
// 	"buckets": null,
// 	"error": ""
// }
