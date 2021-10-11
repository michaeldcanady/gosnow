package gosnow

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/levigross/grequests"
)

type Resource struct {
	Base_url    *url.URL
	Base_path   string
	Api_path    string
	Session     *grequests.Session
	ChunkSize   int
	Url_builder URLBuilder
	Parameters  ParamsBuilder
}

func NewResource(base_url *url.URL, base_path, api_path string, session *grequests.Session, chunkSize int) (R Resource) {

	R.Base_url = base_url
	R.Base_path = base_path
	R.Api_path = api_path
	R.Session = session
	R.ChunkSize = chunkSize
	R.Url_builder = URLBuilderNew(base_url, base_path, api_path)
	R.Parameters = NewParamsBuilder()

	return R
}

func (R Resource) String() string {
	return fmt.Sprintf("<[%s]>", R.path())
}

func (R Resource) path() string {
	return fmt.Sprintf("%s", R.Base_path+R.Api_path)
}

func (R Resource) _request() SnowRequest {

	return SnowRequestNew(R.Parameters, R.Session, R.Url_builder, 0, R)
}

func (R Resource) request(method string, path_append string, headers map[string]string, args map[string]string) (Response, error) {
	return R._request().custom(method, path_append, headers, args)
}

func (R Resource) Get(query interface{}, limits int, offset int, stream bool, fields ...interface{}) (resp Response, err error) {

	if R.Base_path == "" {
		err = errors.New("Failed 'Get': Resource is nil")
		logger.Println(err)
		return resp, err
	}
	display_value := R.Parameters._sysparms["sysparm_display_value"].(bool)
	exclude_reference_link := R.Parameters._sysparms["sysparm_exclude_reference_link"].(bool)
	suppress_pagination_header := R.Parameters._sysparms["sysparm_suppress_pagination_header"].(bool)
	return R._request().get(query, limits, offset, stream, display_value, exclude_reference_link, suppress_pagination_header, fields...)
}

func (R Resource) Delete(query interface{}) (map[string]interface{}, error) {

	return R._request().delete(query)
}

func (R Resource) Create(payload map[string]string) (Response, error) {

	return R._request().create(payload)
}

func (R Resource) Update(query interface{}, payload map[string]string) (Response, error) {
	return R._request().update(query, payload)
}
