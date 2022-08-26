package gosnow

import (
	"fmt"
	"net/url"

	"github.com/levigross/grequests"
)

type Table Resource

// NewTable creates a new instance of the ServiceNow Table API
func NewTable(baseURL *url.URL, basePath string, tableName string, session *grequests.Session, chunkSize int) (T Table) {

	apiPath := fmt.Sprintf("/table/%s", tableName)

	T = Table(NewResource(baseURL, basePath, apiPath, session, 8192))

	return
}

// String returns the string version of the path <[api/now/component/component]>
func (T Table) String() string {
	return Resource(T).String()
}

// Get used to fetch a record
func (T Table) Get(query interface{}, limits int, offset int, stream bool, fields ...interface{}) (resp Response, err error) {

	resp, err = Resource(T).Get(query, limits, offset, stream, fields...)

	return

}

// Delete used to remove a record
func (T Table) Delete(query interface{}) (Response, error) {
	return Resource(T).Delete(query)
}

// Create used to create a new record
func (T Table) Post(args map[string]string) (resp Response, err error) {

	resp, err = Resource(T).Post(args)

	return
}

// Update used to modify an existing record
func (T Table) Update(query interface{}, args map[string]string) (resp Response, err error) {

	resp, err = Resource(T).Update(query, args)

	return
}
