package gosnow

import "github.com/levigross/grequests"

type PreparedRequest struct {
	Method         Method
	Stream         bool
	RequestOptions grequests.RequestOptions
	Request        Request
}

func NewPreparedRequest(request Request, method Method, stream bool, requestOptions grequests.RequestOptions) (P PreparedRequest) {

	P.Method = method
	P.Stream = stream
	P.RequestOptions = requestOptions
	P.Request = request

	return
}

func (P PreparedRequest) Invoke() (Response, error) {
	return P.Request.getResponse(P.Method, P.Stream, P.RequestOptions)
}

func (P PreparedRequest) AsBatchRequest(id string, excludeResponseHeaders bool) BatchRequest {

	Body := ""
	headers := []map[string]string{} //P.Request.Session.RequestOptions.Headers

	return NewBatchRequest(Body, excludeResponseHeaders, headers, id, P.Method, P.Request.URLBuilder.Path)
}
