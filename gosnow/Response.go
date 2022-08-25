package gosnow

import (
	"errors"

	"github.com/levigross/grequests"
)

// Response a ServiceNow API response
type Response struct {
	_response   *grequests.Response
	_chunk_size int
	_count      int
	resource    Resource
	stream      bool
}

// NewResponse generates a response struct
func NewResponse(response *grequests.Response, chunk_size int, resource Resource, stream bool) (R Response) {
	if chunk_size == 0 {
		chunk_size = 8192
	}
	R._response = response
	R._chunk_size = chunk_size
	R._count = 0
	R.resource = resource
	R.stream = stream

	return R
}

// Sanatizes the response for the user
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

// Buffers the reponse recieved to make usable by the user
func (R Response) _get_buffered_response() ([]map[string]interface{}, int, error) {
	response, err := R.getResponse()
	if err != nil {
		//err := errors.New("could not buffer error due to response error")
		return []map[string]interface{}{}, 0, err
	}
	if response.StatusCode == 204 {
		deleted := map[string]interface{}{"status": "record deleted"}
		return []map[string]interface{}{deleted}, 1, nil
	}

	sanitized_response := _sanitize(response)
	return sanitized_response, len(sanitized_response), nil
}

func (R Response) getResponse() (*grequests.Response, error) {
	response := R._response

	if response == nil {
		err := errors.New("Error: Response is empty")
		return nil, err
	}

	if response.Ok {
		switch code := response.StatusCode; {
		case code == 200:
			logger.Println("request completed successfully.")
		case code == 201:
			logger.Println("record created successfully.")
		case code == 204:
			logger.Println("record deleted successfully.")
		case code == 202:
			logger.Println("serviceNow responded with 202 code! Error found.")
		default:
			logger.Printf("ServiceNow reponded with %v.", code)
		}
		return response, nil
	} else {
		//Used for the JSON recieved
		reponseStruct := struct {
			Error  map[string]interface{}
			Status interface{}
		}{}
		//Used for the error code given to the user
		var error ReponseError
		//unmartial error into JSON
		_ = response.JSON(&reponseStruct)
		logger.Printf("serviceNow responded with %v code!", response.StatusCode)
		error.msg = reponseStruct.Error["message"].(string)
		switch code := response.StatusCode; {
		case code <= 409 && code >= 400:
			logger.Println("client Side error detected.")
			error.err = "client Side"
		case code <= 509 && code >= 500:
			logger.Println("server Side error detected.")
			error.err = "server Side"
		default:
			logger.Printf("unknown error code %v returned. info: ", code)
			error.err = "unknown ServiceNow"
		}
		return nil, error
	}
}

// First returns the first record in the map
func (R Response) First() (map[string]interface{}, error) {
	content, _, err := R.All()
	if err != nil {
		//err = fmt.Errorf("could not retrieve first record because of upstream error")
		return map[string]interface{}{}, err
	}
	if len(content) != 0 {
		logger.Println(content[0])
		return content[0], nil
	} else {
		return map[string]interface{}{}, nil
	}
}

// All returns all found serviceNow records in a map slice
func (R Response) All() ([]map[string]interface{}, int, error) {
	return R._get_buffered_response()
}

/*
//Upload is used to attach an image to a request already made
func (R Response) Upload(filePath string, multipart bool) (resp Response, err error) {

	attachments, err := R.resource.attachments()
	if err != nil {
		return
	}

	response := _sanitize(R._response)

	sysID := response[0]["sys_id"].(string)

	return attachments.Upload(sysID, filePath, multipart)
}
*/
/*
func (R Response) Get(limit int) (resp Response, err error) {

	attachments, err := R.resource.attachments()
	if err != nil {
		return
	}

	response := _sanitize(R._response)

	sysID := response[0]["sys_id"].(string)

	return attachments.Get(sysID, limit)
}
*/
