package gosnow

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/levigross/grequests"
)

type SnowRequest struct {
	Session        *grequests.Session
	_url           string
	Chunk_size     int
	Resource       Resource
	ServiceCatalog ServiceCatalog
	Parameters     ParamsBuilder
	URLBuilder     URLBuilder
}

func SnowRequestNew(parameters ParamsBuilder, session *grequests.Session, url_builder URLBuilder, chunk_size int, resource interface{}) (S SnowRequest) {
	S.Parameters = parameters
	S.Session = session
	S._url = url_builder.getURL()
	S.URLBuilder = url_builder
	S.Chunk_size = chunk_size

	switch resource := resource.(type) {
	case Resource:
		S.Resource = resource
	case ServiceCatalog:
		S.ServiceCatalog = resource
	}

	return S
}

func (S SnowRequest) get(query interface{}, limits int, offset int, stream bool, display_value, exclude_reference_link,
	suppress_pagination_header bool, fields ...interface{}) (Response, error) {
	if _, ok := query.(string); ok {
		S.Parameters._sysparms["sysparm_query"] = query.(string)
	} else if _, ok := query.(map[string]interface{}); ok {
		S.Parameters.query(query.(map[string]interface{}))
	} else {
		log.Fatalf("%T is not a supported type for query. Please use string or map[string]interface{}", query)
	}
	S.Parameters.limit(limits)
	S.Parameters.offset(offset)
	S.Parameters.fields(fields...)
	S.Parameters.display_value(display_value)
	S.Parameters.exclude_reference_link(exclude_reference_link)
	S.Parameters.suppress_pagination_header(suppress_pagination_header)

	return S._get_response("GET", stream, grequests.RequestOptions{})
}

func (S SnowRequest) create(payload grequests.RequestOptions) (Response, error) {

	response, err := S._get_response("POST", false, payload)
	if err != nil {
		err = fmt.Errorf("response Error: %v", err)
		logger.Println(err)
		return Response{}, err
	}
	fmt.Println(response._response.StatusCode)
	return S._get_response("POST", false, payload)
}

func (S SnowRequest) _get_response(method string, stream bool, payload grequests.RequestOptions) (Response, error) {
	var response *grequests.Response
	var err error
	if method == "GET" {
		response, err = S.Session.Get(S._url, S.Parameters.as_dict())
		if err != nil {
			err = fmt.Errorf("get request Failed: %v", err)
			log.Println(err)
			return Response{}, err
		}
	} else {

		switch method {
		case "POST":
			if S._url == "<nil>" {
				err := fmt.Errorf("URL error: URL is Empty")
				logger.Println(err)
				return Response{}, err
			}
			payload1 := (*S.Parameters.as_dict())
			payload1.Headers = payload.Headers
			payload1.JSON = payload.JSON

			fmt.Println(payload1)

			response, err = S.Session.Post(S._url, &payload1)
			if err != nil {
				err = fmt.Errorf("post request Failed: %v", err)
				log.Println(err)
				return Response{}, err
			}
		case "PUT":
			response, err = S.Session.Put(S._url, &payload)
			if err != nil {
				err = fmt.Errorf("Put request Failed: %v", err)
				log.Println(err)
				return Response{}, err
			}
		case "DELETE":
			response, err = S.Session.Delete(S._url, &payload)
			if err != nil {
				err = fmt.Errorf("delete request Failed: %v", err)
				log.Println(err)
				return Response{}, err
			}
		}
	}

	return NewResponse(response, S.Chunk_size, S.Resource, stream), nil
}

func (S SnowRequest) _get_custom_endpoint(value string) string {
	fmt.Printf("%s'\n", value)
	if !strings.HasPrefix(value, "/") {
		value = fmt.Sprintf("/%s", value)
	}
	return S.URLBuilder.get_appended_custom(value)
}

func (S SnowRequest) update(query interface{}, payload grequests.RequestOptions) (Response, error) {
	limits, err := S.Parameters.getlimit()
	if err != nil {
		err = fmt.Errorf("failed to get limit due to: %v", err)
		logger.Println(err)
		return Response{}, err
	}
	offset := S.Parameters.getoffset()
	display_value := S.Parameters.getdisplay_value()
	exclude_reference_link := S.Parameters.getexclude_reference_link()
	suppress_pagination_header := S.Parameters.getsuppress_pagination_header()
	record, err := S.get(query, limits, offset, false, display_value, exclude_reference_link, suppress_pagination_header, nil)
	if err != nil {
		err = fmt.Errorf("get error: %v", err)
	}
	first_record, err := record.First()
	if err != nil {
		return Response{}, errors.New("could not update due to querying error")
	}
	S._url = S._get_custom_endpoint(first_record["sys_id"].(string))

	return S._get_response("PUT", false, payload)
}

func (S SnowRequest) delete(query interface{}) (map[string]interface{}, error) {
	offset := S.Parameters.getoffset()
	display_value := S.Parameters.getdisplay_value()
	exclude_reference_link := S.Parameters.getexclude_reference_link()
	suppress_pagination_header := S.Parameters.getsuppress_pagination_header()
	resp, _ := S.get(query, 1, offset, false, display_value, exclude_reference_link, suppress_pagination_header, nil)
	record, _ := resp.First()
	S._url = S._get_custom_endpoint(record["sys_id"].(string))
	resp, _ = S._get_response("DELETE", false, grequests.RequestOptions{})

	return resp.First()
}

func (S SnowRequest) custom(method string, pathAppend string, payload grequests.RequestOptions) (Response, error) {

	if pathAppend != "" {
		fmt.Printf("%s'\n", pathAppend)
		S._url = S._get_custom_endpoint(pathAppend)
	}

	return S._get_response(method, false, payload)
}
