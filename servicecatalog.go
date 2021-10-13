package gosnow

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/levigross/grequests"
)

type ServiceCatalog struct {
	Base_url    *url.URL
	Base_path   string
	Api_path    string
	Session     *grequests.Session
	ChunkSize   int
	Url_builder URLBuilder
	Parameters  ParamsBuilder
}

func NewServiceCatalog(base_url *url.URL, base_path, api_path string, session *grequests.Session, chunkSize int) (S ServiceCatalog) {
	S.Base_url = base_url
	S.Base_path = base_path
	S.Api_path = api_path
	S.Session = session
	S.ChunkSize = chunkSize
	S.Url_builder = URLBuilderNew(base_url, base_path, api_path)
	S.Parameters = NewParamsBuilder()

	return
}

// String returns a string representation of the path
func (S ServiceCatalog) String() string {
	return fmt.Sprintf("<[%s]>", S.path())
}

// path returns a string representation of the path
func (S ServiceCatalog) path() string {
	return fmt.Sprintln(S.Base_path + S.Api_path)
}

func (S ServiceCatalog) _request() SnowRequest {
	return SnowRequestNew(S.Parameters, S.Session, S.Url_builder, 0, S)
}

func (S ServiceCatalog) Get(query interface{}) (resp Response, err error) {
	if S.Base_path == "" {
		err = errors.New("failed 'Get': ServiceCatalog is nil")
		logger.Println(err)
		return resp, err
	}

	return S._request().get(query, 0, 0, false, false, false, false)
}
