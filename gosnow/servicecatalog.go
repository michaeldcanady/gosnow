package gosnow

import (
	"net/url"
	"reflect"

	"github.com/levigross/grequests"
)

// ServiceCatalog the service catalog API
type ServiceCatalog Resource

// NewTable creates a new instance of the ServiceNow Table API
func NewServiceCatalog(baseURL *url.URL, apiPath string, session *grequests.Session, chunkSize int) (S ServiceCatalog) {

	S = ServiceCatalog(NewResource(baseURL, apiPath, session, 8192))

	return
}

// String returns the string version of the path <[api/now/component/component]>
func (S ServiceCatalog) String() string {
	return Resource(S).String()
}

// Get used to fetch a record
func (S ServiceCatalog) Get(query interface{}) PreparedRequest {
	return Resource(S).Get(reflect.TypeOf(S), query, 0, 0, false, nil)
}

// Delete used to remove a record
func (S ServiceCatalog) Delete(query interface{}) PreparedRequest {
	return Resource(S).Delete(query)
}

// Create used to create a new record
func (S ServiceCatalog) Post(requestType reflect.Type, args map[string]interface{}) PreparedRequest {
	return Resource(S).Post(requestType, args)
}

// Update used to modify an existing record
func (S ServiceCatalog) Update(query interface{}, args map[string]interface{}) PreparedRequest {

	return Resource(S).Update(reflect.TypeOf(S), query, args)
}
