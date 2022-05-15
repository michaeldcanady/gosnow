package gosnow

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/levigross/grequests"
)

type ServiceCatalog struct {
	BaseURL     *url.URL
	BasePath    string
	ApiPath     string
	Session     *grequests.Session
	ChunkSize   int
	Url_builder URLBuilder
	Parameters  ParamsBuilder
}

func NewServiceCatalog(BaseURL *url.URL, BasePath, ApiPath string, session *grequests.Session, chunkSize int) (S ServiceCatalog) {
	S.BaseURL = BaseURL
	S.BasePath = BasePath
	S.ApiPath = ApiPath
	S.Session = session
	S.ChunkSize = chunkSize
	S.Url_builder = URLBuilderNew(BaseURL, BasePath, ApiPath)
	S.Parameters = NewParamsBuilder()

	return
}

// String returns a string representation of the path
func (S ServiceCatalog) String() string {
	return fmt.Sprintf("<[%s]>", S.path())
}

// path returns a string representation of the path
func (S ServiceCatalog) path() string {
	return fmt.Sprintln(S.BasePath + S.ApiPath)
}

func (S ServiceCatalog) _request() SnowRequest {
	return SnowRequestNew(S.Parameters, S.Session, S.Url_builder, 0, S)
}

func (S ServiceCatalog) Get(query interface{}) (resp Response, err error) {
	if S.BasePath == "" {
		err = errors.New("failed 'Get': ServiceCatalog is nil")
		logger.Println(err)
		return resp, err
	}

	return S._request().get(query, 0, 0, false, false, false, false)
}

func (S ServiceCatalog) Post() (resp Response, err error) {
	return S._request().create(grequests.RequestOptions{})
}
