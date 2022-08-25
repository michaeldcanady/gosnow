package gosnow

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

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

//NewResource returns a new serviceNow API resource
func NewResource(BaseURL *url.URL, BasePath, ApiPath string, session *grequests.Session, chunkSize int) (R Resource) {
	R.url = BaseURL
	R.url.Path = fmt.Sprintf("%s%s", BasePath, ApiPath)
	R.Session = session
	R.tableName = ApiPath
	R.ChunkSize = chunkSize
	R.Parameters = NewParamsBuilder()

	return
}

func (R Resource) toJSON(args map[string]string) (JSON []byte, err error) {
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

func (R Resource) attachments() (A Attachment, err error) {
	copyResource := R

	copyResource.url.Path = "/api/now/attachment"

	return NewAttachment(copyResource, R.tableName), err
}

func (R Resource) attachment() (A Attachment, err error) {
	copyResource := R

	copyResource.url.Path = "/api/now/attachment"

	return NewAttachment(copyResource, R.tableName), err
}

func (R Resource) request(method string, path_append string, payload grequests.RequestOptions) (resp Response, err error) {
	return R._request().custom(method, path_append, payload)
}

//Get used to fetch a record
func (R Resource) Get(query interface{}, limits int, offset int, stream bool, fields ...interface{}) (resp Response, err error) {
	display_value := R.Parameters._sysparms["sysparm_display_value"].(bool)
	exclude_reference_link := R.Parameters._sysparms["sysparm_exclude_reference_link"].(bool)
	suppress_pagination_header := R.Parameters._sysparms["sysparm_suppress_pagination_header"].(bool)

	return R._request().get(query, limits, offset, stream, display_value, exclude_reference_link, suppress_pagination_header, fields...)
}

//Delete used to remove a record
func (R Resource) Delete(query interface{}) (Response, error) {
	return R._request().delete(query)
}

//Create used to create a new record
func (R Resource) Create(args map[string]string) (resp Response, err error) {

	var payload grequests.RequestOptions

	payload.JSON, err = R.toJSON(args)
	if err != nil {
		return
	}

	return R._request().create(payload)
}

//Update used to modify an existing record
func (R Resource) Update(query interface{}, args map[string]string) (resp Response, err error) {

	var payload grequests.RequestOptions

	payload.JSON, err = R.toJSON(args)
	if err != nil {
		return
	}

	return R._request().update(query, payload)
}
