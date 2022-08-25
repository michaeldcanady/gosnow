package gosnow

import (
	"fmt"
	"net/url"

	"github.com/levigross/grequests"
)

// ServiceCatalog the service catalog API
type ServiceCatalog Resource

// NewTable creates a new instance of the ServiceNow Table API
func NewServiceCatalog(baseURL *url.URL, basePath string, apiPath string, session *grequests.Session, chunkSize int) (S ServiceCatalog) {

	basePath += "/sn_sc/servicecatalog"

	S = ServiceCatalog(NewResource(baseURL, basePath, apiPath, session, 8192))

	fmt.Println(S.url)

	return
}

// String returns the string version of the path <[api/now/component/component]>
func (S ServiceCatalog) String() string {
	return Resource(S).String()
}

// Get used to fetch a record
func (S ServiceCatalog) Get(query interface{}) (resp Response, err error) {

	resp, err = Resource(S).Get(query, 0, 0, false, nil)

	return

}

// Delete used to remove a record
func (S ServiceCatalog) Delete(query interface{}) (Response, error) {
	return Resource(S).Delete(query)
}

// Create used to create a new record
func (S ServiceCatalog) Create(args map[string]string) (resp Response, err error) {

	resp, err = Resource(S).Create(args)

	return
}

// Update used to modify an existing record
func (S ServiceCatalog) Update(query interface{}, args map[string]string) (resp Response, err error) {

	resp, err = Resource(S).Update(query, args)

	return
}