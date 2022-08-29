package gosnow

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/levigross/grequests"
)

type Batch Resource

func NewBatch(baseURL *url.URL, apiPath string, session *grequests.Session, chunkSize int) (B Batch) {

	B = Batch(NewResource(baseURL, apiPath, session, 8192))

	return
}

func (B Batch) Post(requests []BatchRequest) ([]interface{}, error) {

	batchRequestId := "2"

	args := map[string]interface{}{}
	args["batch_request_id"] = batchRequestId
	args["rest_requests"] = requests

	resp, err := Resource(B).Post(reflect.TypeOf(B), args).Invoke()
	fmt.Println(err)
	if err != nil {
		return []interface{}{}, err
	}

	//_, _, err = resp.(Response).All()
	//fmt.Println(err)

	batchedResp, err := resp.(Response).toBatchedResponse()

	if err != nil {
		return []interface{}{}, err
	}

	goodResults := []interface{}{}

	fmt.Println(batchedResp)

	for _, response := range batchedResp.ServicedRequests {
		for _, request := range requests {
			if response.Id == request.Id {
				// Convert each list of responses to thier appropriate types
				switch request.requestType.String() {
				case "gosnow.Table":
					results := batchedResp.response._sanitize(response.Body)
					for _, v := range results {
						goodResults = append(goodResults, TableEntry(v))
					}
				case "gosnow.Attachment":
					results := batchedResp.response._sanitize(response.Body)
					for _, v := range results {
						goodResults = append(goodResults, AttachmentEntry(v))
					}
				}
			}
		}
	}

	return goodResults, nil
}
