package gosnow

import (
	"fmt"
	"net/url"

	"github.com/levigross/grequests"
)

type ServiceCatalog struct {
	url         *url.URL
	Session     *grequests.Session
	ChunkSize   int
	Url_builder URLBuilder
	Parameters  ParamsBuilder
}

func NewServiceCatalog(BaseURL *url.URL, BasePath, ApiPath string, session *grequests.Session, chunkSize int) (S ServiceCatalog) {
	S.url = BaseURL
	S.url.Path = fmt.Sprintf("%s%s", BasePath, ApiPath)
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

// path returns a string representation of the api path
func (S ServiceCatalog) path() string {
	return S.url.Path
}

// _request returns a new SNow Request
func (S ServiceCatalog) _request() Request {
	return NewRequest(S.Parameters, S.Session, S.url, 0, S)
}

// Get returns a response and an error
func (S ServiceCatalog) Get(query interface{}) (resp Response, err error) {
	return S._request().get(query, 0, 0, false, false, false, false)
}

//func (S ServiceCatalog) Post() (resp Response, err error) {
//	return S._request().create(grequests.RequestOptions{})
//}
