package gosnow

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strings"

	"github.com/levigross/grequests"
)

/* Request used to preform requests against the ServiceNow APIs. Contains the querying URL and the parameters
 */
type Request struct {
	Session        *grequests.Session
	url            string
	URLBuilder     *url.URL
	Chunk_size     int
	Resource       Resource
	ServiceCatalog ServiceCatalog
	Parameters     ParamsBuilder
}

// NewRequest used to create a new serviceNow request
func NewRequest(parameters ParamsBuilder, session *grequests.Session, url_builder *url.URL, chunk_size int, resource interface{}) (R Request) {
	R.Parameters = parameters
	R.Session = session
	R.URLBuilder = url_builder
	R.url = url_builder.String()
	R.Chunk_size = chunk_size

	return
}

func (R Request) get(requestType reflect.Type, query interface{}, limits int, offset int, stream bool, display_value, exclude_reference_link,
	suppress_pagination_header bool, fields ...interface{}) PreparedRequest {
	if _, ok := query.(string); ok {
		R.Parameters._sysparms["sysparm_query"] = query.(string)
	} else if _, ok := query.(map[string]interface{}); ok {
		R.Parameters.query(query.(map[string]interface{}))
	} else {
		log.Fatalf("%T is not a supported type for query. Please use string or map[string]interface{}", query)
	}
	R.Parameters.limit(limits)
	R.Parameters.offset(offset)
	R.Parameters.fields(fields...)
	R.Parameters.display_value(display_value)
	R.Parameters.exclude_reference_link(exclude_reference_link)
	R.Parameters.suppress_pagination_header(suppress_pagination_header)

	return NewPreparedRequest(requestType, R, GET, stream, grequests.RequestOptions{})
}

func (R Request) getResponse(method Method, stream bool, payload grequests.RequestOptions) (resp Response, err error) {
	var response *grequests.Response

	if R.url == "<nil>" {
		err = fmt.Errorf("URL error: URL is Empty")
		logger.Println(err)
		return Response{}, err
	}

	switch method {
	case GET:
		response, err = R.Session.Get(R.url, R.Parameters.as_dict())
	case POST:
		payload1 := (*R.Parameters.as_dict())
		payload1.Headers = payload.Headers
		payload1.JSON = payload.JSON

		response, err = R.Session.Post(R.url, &payload1)
	case PUT:
		response, err = R.Session.Put(R.url, &payload)
	case DELETE:
		response, err = R.Session.Delete(R.url, &payload)
	}

	if err != nil {
		err = fmt.Errorf("request Failed: %s, %v", method, err)
		log.Println(err)
		return Response{}, err
	}

	return NewResponse(response, R.Chunk_size, R.Resource, stream), nil
}

func (R Request) delete(requestType reflect.Type, query interface{}) PreparedRequest {
	offset := R.Parameters.getoffset()
	display_value := R.Parameters.getdisplay_value()
	exclude_reference_link := R.Parameters.getexclude_reference_link()
	suppress_pagination_header := R.Parameters.getsuppress_pagination_header()

	request := R.get(nil, query, 1, offset, false, display_value, exclude_reference_link, suppress_pagination_header, nil)

	resp, _ := request.Invoke()

	record, _ := resp.(Response).First()

	if len(record) == 0 {
		return PreparedRequest{} //, errors.New("no record retrieve, unable to complete delete request")
	}

	R.url = R.getCustomEndpoint(record["sys_id"].(string))

	return NewPreparedRequest(requestType, R, DELETE, false, grequests.RequestOptions{})
}

func (R Request) getCustomEndpoint(value string) string {
	fmt.Printf("%s'\n", value)
	if !strings.HasPrefix(value, "/") {
		value = fmt.Sprintf("/%s", value)
	}

	R.URLBuilder.Path = R.URLBuilder.Path + value

	return R.URLBuilder.String()
}

func (R Request) post(requestType reflect.Type, payload grequests.RequestOptions) PreparedRequest {
	R.getResponse(POST, false, payload)
	return NewPreparedRequest(requestType, R, GET, false, grequests.RequestOptions{})
}

func (R Request) update(requestType reflect.Type, query interface{}, payload grequests.RequestOptions) PreparedRequest {
	limits, err := R.Parameters.getlimit()
	if err != nil {
		err = fmt.Errorf("failed to get limit due to: %v", err)
		logger.Println(err)
		return PreparedRequest{}
	}
	offset := R.Parameters.getoffset()
	display_value := R.Parameters.getdisplay_value()
	exclude_reference_link := R.Parameters.getexclude_reference_link()
	suppress_pagination_header := R.Parameters.getsuppress_pagination_header()
	request := R.get(nil, query, limits, offset, false, display_value, exclude_reference_link, suppress_pagination_header, nil)

	record, err := request.Invoke()
	if err != nil {
		err = fmt.Errorf("get error: %v", err)
		return PreparedRequest{}
	}
	first_record, err := record.(Response).First()
	if err != nil {
		return PreparedRequest{} //, errors.New("could not update due to querying error")
	}
	R.url = R.getCustomEndpoint(first_record["sys_id"].(string))

	return NewPreparedRequest(requestType, R, PUT, false, payload)
}
