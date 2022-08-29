package gosnow

import "reflect"

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type BatchRequest struct {
	requestType            reflect.Type        `json:"-"`
	Body                   string              `json:",omitempty"`
	ExcludeResponseHeaders bool                `json:"exclude_response_headers"`
	Headers                []map[string]string `json:"headers"`
	Id                     string              `json:"id"`
	Method                 Method              `json:"method"`
	Url                    string              `json:"url"`
}

func NewBatchRequest(requestType reflect.Type, body string, excludeResponseHeaders bool, headers []map[string]string, id string, method Method, url string) (B BatchRequest) {

	if len(headers) == 0 {
		headers = []map[string]string{0: {"name": "Content-Type", "value": "application/json"}, 1: {"name": "Accept", "value": "application/json"}}
	}

	B.requestType = requestType
	B.Body = body
	B.ExcludeResponseHeaders = excludeResponseHeaders
	B.Headers = headers
	B.Id = id
	B.Method = method
	B.Url = url

	return
}
