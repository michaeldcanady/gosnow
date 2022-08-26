package gosnow

import (
	"fmt"
	"net/url"

	"github.com/levigross/grequests"
)

type Batch Resource

func NewBatch(baseURL *url.URL, apiPath string, session *grequests.Session, chunkSize int) (B Batch) {

	B = Batch(NewResource(baseURL, apiPath, session, 8192))

	return
}

func (B Batch) Post(requests []BatchRequest) (Response, error) {

	batchRequestId := "1"

	args := map[string]interface{}{}
	args["batch_request_id"] = batchRequestId
	args["rest_requests"] = requests

	fmt.Println(requests)

	return Resource(B).Post(args)
}
