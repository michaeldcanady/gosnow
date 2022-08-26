package gosnow

import (
	"net/url"
	"reflect"

	"github.com/levigross/grequests"
)

type Table Resource

// NewTable creates a new instance of the ServiceNow Table API
func NewTable(baseURL *url.URL, apiPath string, session *grequests.Session, chunkSize int) (T Table) {

	T = Table(NewResource(baseURL, apiPath, session, 8192))

	return
}

// String returns the string version of the path <[api/now/component/component]>
func (T Table) String() string {
	return Resource(T).String()
}

// Get used to fetch a record
func (T Table) Get(query interface{}, limits int, offset int, stream bool, fields ...interface{}) PreparedRequest {

	return Resource(T).Get(reflect.TypeOf(T), query, limits, offset, stream, fields...)

}

// Delete used to remove a record
func (T Table) Delete(query interface{}) PreparedRequest {
	return Resource(T).Delete(query)
}

// Post used to create a new record
func (T Table) Post(args map[string]interface{}) (resp Response, err error) {

	prop, err := Resource(T).Post(reflect.TypeOf(T), args).Invoke()

	return prop.(Response), err
}

// Update used to modify an existing record
func (T Table) Update(query interface{}, args map[string]interface{}) PreparedRequest {

	return Resource(T).Update(reflect.TypeOf(T), query, args)
}
