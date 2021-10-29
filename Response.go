package gosnow

import (
	"fmt"

	"github.com/levigross/grequests"
)

type Response struct {
	_response   *grequests.Response
	_chunk_size int
	_count      int
	_resource   Resource
	_stream     bool
}

type ErrorWrapper struct {
	error  ErrorMessage
	status string
}
type ErrorMessage struct {
	message string
	detail  string
}

func NewResponse(response *grequests.Response, chunk_size int, resource Resource, stream bool) (R Response) {
	if chunk_size == 0 {
		chunk_size = 8192
	}
	R._response = response
	R._chunk_size = chunk_size
	R._count = 0
	R._resource = resource
	R._stream = stream

	return R
}

type Test interface{}

func _sanitize(response *grequests.Response) []map[string]interface{} {
	var dT = make(map[string]interface{})

	err := response.JSON(&dT)
	if err != nil {
		logger.Fatal("response error " + err.Error())
	}

	var returnValue = make([]map[string]interface{}, 0)
sanitize:
	for _, r := range dT {
		if _, ok := r.(map[string]interface{}); ok {
			returnValue = append(returnValue, r.(map[string]interface{}))
			break sanitize
		} else if _, ok := r.([]interface{}); ok {
			for _, r := range r.([]interface{}) {
				returnValue = append(returnValue, r.(map[string]interface{}))
			}
		} else {
			logger.Println(r)
		}
	}
	return returnValue
}

func (R Response) _get_buffered_response() ([]map[string]interface{}, int, error) {
	response, err := R._get_response()
	if err != nil {
		err := fmt.Errorf("Could not buffer error due to response error")
		return []map[string]interface{}{}, 0, err
	}
	if response.StatusCode == 204 {
		deleted := map[string]interface{}{"status": "record deleted"}
		return []map[string]interface{}{deleted}, 1, nil
	}

	sanitized_response := _sanitize(response)
	return sanitized_response, len(sanitized_response), nil
}

func (R Response) _get_response() (*grequests.Response, error) {
	response := R._response

	if response == nil {
		err := fmt.Errorf("Error: Response is empty")
		return nil, err
	}

	if response.Ok {
		switch code := response.StatusCode; {
		case code == 200:
			logger.Println("Request completed successfully.")
		case code == 201:
			logger.Println("Record created successfully.")
		case code == 204:
			logger.Println("Record deleted successfully.")
		case code == 202:
			logger.Printf("ServiceNow responded with 202 code! Error found.\n")
		}
		return response, nil
	} else {
		//TODO find a way to get the error code from the response and return it
		logger.Printf("ServiceNow responded with %v code!", response.StatusCode)
		switch code := response.StatusCode; {
		case code <= 409 && code >= 400:
			logger.Println("Client Side error detected.")
		case code <= 509 && code >= 500:
			logger.Println("Server Side error detected.")
		default:
			logger.Printf("Unknown error code %v returned. info: ", code)
		}
		return response, nil
	}
}

func (R Response) First() (map[string]interface{}, error) {
	content, _, err := R.All()
	if err != nil {
		err = fmt.Errorf("could not retrieve first record because of upstream error")
		return map[string]interface{}{}, err
	}
	logger.Println(content[0])
	return content[0], nil
}

func (R Response) All() ([]map[string]interface{}, int, error) {
	return R._get_buffered_response()
}

func (R Response) Upload(filePath string, multipart bool) (resp Response, err error) {

	attachments, err := R._resource.attachments()
	if err != nil {
		return
	}

	response := _sanitize(R._response)

	sysID := response[0]["sys_id"].(string)

	return attachments.Upload(sysID, filePath, multipart)
}
