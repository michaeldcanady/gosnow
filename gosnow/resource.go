package gosnow

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"reflect"

	"github.com/levigross/grequests"
)

// Resource representations of the vague service now API api/now
type Resource struct {
	url        *url.URL
	Session    *grequests.Session
	tableName  string
	ChunkSize  int
	Parameters ParamsBuilder
}

// NewResource returns a new serviceNow API resource
func NewResource(BaseURL *url.URL, ApiPath string, session *grequests.Session, chunkSize int) (R Resource) {
	R.url = BaseURL
	R.url.Path = ApiPath
	R.Session = session
	R.tableName = ApiPath
	R.ChunkSize = chunkSize
	R.Parameters = NewParamsBuilder()

	return
}

func (R Resource) toJSON(args map[string]interface{}) (JSON []byte, err error) {
	JSON, err = json.Marshal(args)
	if err != nil {
		err = fmt.Errorf("issue marshalling args into Javascript: %v", err)
		log.Println(err)
	}
	return
}

// returns the string version of the path <[api/now/component/component]>
func (R Resource) String() string {
	return fmt.Sprintf("<[%s]>", R.path())
}

func (R Resource) path() string {
	return fmt.Sprintln(R.url.Path)
}

func (R Resource) _request() Request {
	return NewRequest(R.Parameters, R.Session, R.url, 0, R)
}

// Get used to generate a prepared request
func (R Resource) Get(requestType reflect.Type, query interface{}, limits int, offset int, stream bool, fields ...interface{}) Request {
	display_value := R.Parameters._sysparms["sysparm_display_value"].(bool)
	exclude_reference_link := R.Parameters._sysparms["sysparm_exclude_reference_link"].(bool)
	suppress_pagination_header := R.Parameters._sysparms["sysparm_suppress_pagination_header"].(bool)

	return R._request().get(requestType, query, limits, offset, stream, display_value, exclude_reference_link, suppress_pagination_header, fields...)
}

// Delete used to remove a record
func (R Resource) Delete(requestType reflect.Type, query interface{}) Request {
	return R._request().delete(requestType, query)
}

// Create used to create a new record
func (R Resource) Post(requestType reflect.Type, args map[string]interface{}) Request {

	var payload grequests.RequestOptions
	var err error

	payload.JSON, err = R.toJSON(args)
	if err != nil {
		return Request{}
	}

	return R._request().post(requestType, payload)
}

// Update used to modify an existing record
func (R Resource) Update(requestType reflect.Type, query interface{}, args map[string]interface{}) Request {

	var payload grequests.RequestOptions
	var err error

	payload.JSON, err = R.toJSON(args)
	if err != nil {
		return Request{}
	}

	return R._request().update(requestType, query, payload)
}
