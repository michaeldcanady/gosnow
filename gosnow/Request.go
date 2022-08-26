package gosnow

import (
	"errors"
	"fmt"
	"log"
	"net/url"
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

func (R Request) get(query interface{}, limits int, offset int, stream bool, display_value, exclude_reference_link,
	suppress_pagination_header bool, fields ...interface{}) (Response, error) {
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

	return R.getResponse("GET", stream, grequests.RequestOptions{})
}

func (R Request) getResponse(method string, stream bool, payload grequests.RequestOptions) (Response, error) {
	var response *grequests.Response
	var err error
	if method == "GET" {
		response, err = R.Session.Get(R.url, R.Parameters.as_dict())
		if err != nil {
			err = fmt.Errorf("get request Failed: %v", err)
			log.Println(err)
			return Response{}, err
		}
	} else {

		switch method {
		case "POST":
			if R.url == "<nil>" {
				err := fmt.Errorf("URL error: URL is Empty")
				logger.Println(err)
				return Response{}, err
			}
			payload1 := (*R.Parameters.as_dict())
			payload1.Headers = payload.Headers
			payload1.JSON = payload.JSON

			response, err = R.Session.Post(R.url, &payload1)
			if err != nil {
				err = fmt.Errorf("post request Failed: %v", err)
				log.Println(err)
				return Response{}, err
			}
		case "PUT":
			response, err = R.Session.Put(R.url, &payload)
			if err != nil {
				err = fmt.Errorf("put request Failed: %v", err)
				log.Println(err)
				return Response{}, err
			}
		case "DELETE":
			fmt.Println(R.url)
			response, err = R.Session.Delete(R.url, &payload)
			if err != nil {
				err = fmt.Errorf("delete request Failed: %v", err)
				log.Println(err)
				return Response{}, err
			}
		}
	}

	return NewResponse(response, R.Chunk_size, R.Resource, stream), nil
}

func (R Request) delete(query interface{}) (Response, error) {
	offset := R.Parameters.getoffset()
	display_value := R.Parameters.getdisplay_value()
	exclude_reference_link := R.Parameters.getexclude_reference_link()
	suppress_pagination_header := R.Parameters.getsuppress_pagination_header()

	resp, _ := R.get(query, 1, offset, false, display_value, exclude_reference_link, suppress_pagination_header, nil)

	record, _ := resp.First()

	if len(record) == 0 {
		return Response{}, errors.New("no record retrieve, unable to complete delete request")
	}

	R.url = R.getCustomEndpoint(record["sys_id"].(string))

	resp, _ = R.getResponse("DELETE", false, grequests.RequestOptions{})

	return resp, nil
}

func (R Request) getCustomEndpoint(value string) string {
	fmt.Printf("%s'\n", value)
	if !strings.HasPrefix(value, "/") {
		value = fmt.Sprintf("/%s", value)
	}

	R.URLBuilder.Path = R.URLBuilder.Path + value

	return R.URLBuilder.String()
}

func (R Request) post(payload grequests.RequestOptions) (Response, error) {
	return R.getResponse("POST", false, payload)
}

func (R Request) update(query interface{}, payload grequests.RequestOptions) (Response, error) {
	limits, err := R.Parameters.getlimit()
	if err != nil {
		err = fmt.Errorf("failed to get limit due to: %v", err)
		logger.Println(err)
		return Response{}, err
	}
	offset := R.Parameters.getoffset()
	display_value := R.Parameters.getdisplay_value()
	exclude_reference_link := R.Parameters.getexclude_reference_link()
	suppress_pagination_header := R.Parameters.getsuppress_pagination_header()
	record, err := R.get(query, limits, offset, false, display_value, exclude_reference_link, suppress_pagination_header, nil)
	if err != nil {
		err = fmt.Errorf("get error: %v", err)
		return Response{}, err
	}
	first_record, err := record.First()
	if err != nil {
		return Response{}, errors.New("could not update due to querying error")
	}
	R.url = R.getCustomEndpoint(first_record["sys_id"].(string))

	return R.getResponse("PUT", false, payload)
}

func (R Request) custom(method string, pathAppend string, payload grequests.RequestOptions) (Response, error) {

	if pathAppend != "" {
		fmt.Printf("%s'\n", pathAppend)
		R.url = R.getCustomEndpoint(pathAppend)
	}

	return R.getResponse(method, false, payload)
}
