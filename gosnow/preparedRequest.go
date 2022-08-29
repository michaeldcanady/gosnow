package gosnow

import (
	"net/url"
	"reflect"

	"github.com/levigross/grequests"
)

type PreparedRequest struct {
	requestType    reflect.Type
	Method         Method
	Stream         bool
	RequestOptions grequests.RequestOptions
	Request        *Request
}

func NewPreparedRequest(requestType reflect.Type, request Request, method Method, stream bool, requestOptions grequests.RequestOptions) (P PreparedRequest) {

	P.requestType = requestType
	P.Method = method
	P.Stream = stream
	P.RequestOptions = requestOptions
	P.Request = &request

	return
}

// Invoke runs the prepared request
func (P PreparedRequest) Invoke() (interface{}, error) {
	switch P.requestType.String() {
	case "gosnow.Table":
		response, err := P.Request.getResponse(P.Method, P.Stream, P.RequestOptions)
		if err != nil {
			return nil, err
		}
		return TableResponse(response), nil
	default:
		return P.Request.getResponse(P.Method, P.Stream, P.RequestOptions)
	}
}

func (P PreparedRequest) AsBatchRequest(id string, excludeResponseHeaders bool) BatchRequest {

	parsedQuery := url.Values{}

	requestType := P.requestType

	Body := ""
	headers := []map[string]string{}

	params := P.Request.Parameters.as_dict().Params

	for key, value := range params {
		parsedQuery.Set(key, value)
	}

	URI := P.Request.URLBuilder.Path + "?" + parsedQuery.Encode()

	return NewBatchRequest(requestType, Body, excludeResponseHeaders, headers, id, P.Method, URI)
}
