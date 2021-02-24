package gosnow

import(
  "github.com/levigross/grequests"
  "fmt"
  "strings"
  "encoding/json"
  "log"
)

type SnowRequest struct {
  Session     *grequests.Session
  _url        string
  Chunk_size  int
  Resource    Resource
  Parameters  ParamsBuilder
  URLBuilder  URLBuilder
}

func SnowRequestNew(parameters ParamsBuilder, session *grequests.Session, url_builder URLBuilder, chunk_size int, resource Resource) (S SnowRequest){
  S.Parameters = parameters
  S.Session    = session
  S._url       = url_builder.getURL()
  S.URLBuilder = url_builder
  S.Chunk_size = chunk_size
  S.Resource   = resource

  return S
}

func (S SnowRequest) get(query interface{}, limits int, offset int, stream bool, display_value, exclude_reference_link,
  suppress_pagination_header bool, fields ...interface{}) (Response, error) {
  if _, ok := query.(string); ok {
    S.Parameters._sysparms["sysparm_query"] = query.(string)
  }else if _, ok := query.(map[string]interface{}); ok{
    S.Parameters.query(query.(map[string]interface{}))
  }else{
    log.Fatalf("%T is not a supported type for query. Please use string or map[string]interface{}",query)
  }
  S.Parameters.limit(limits)
  S.Parameters.offset(offset)
  S.Parameters.fields(fields...)
  S.Parameters.display_value(display_value)
  S.Parameters.exclude_reference_link(exclude_reference_link)
  S.Parameters.suppress_pagination_header(suppress_pagination_header)

  return S._get_response("GET",stream)
}

func (S SnowRequest) create(payload map[string]string) (Response, error) {
  response, err := S._get_response("POST",false,payload)
  if err != nil {
    err = fmt.Errorf("Response Error: %v", err)
    logger.Println(err)
    return Response{}, err
  }
  fmt.Println(response._response.StatusCode)
  return S._get_response("POST",false,payload)
}

func (S SnowRequest) _get_response(method string, stream bool, payloadSlice ...map[string]string) (Response, error){
  var response *grequests.Response
  var err error
  if method == "GET"{
    response, err = S.Session.Get(S._url, S.Parameters.as_dict())
    if err != nil{
      err = fmt.Errorf("Get request Failed: %v\n",err)
      log.Println(err)
      return Response{}, err
    }
  } else{
    pay1 := payloadSlice[0]

    jsonString, err := json.Marshal(pay1)
    if err != nil {
      err = fmt.Errorf("Issue marshalling payload into Javascript: %v\n",err)
      log.Println(err)
      return Response{}, err
    }

    payload := grequests.RequestOptions{
      JSON: jsonString,
      }

    switch method {
    case "POST":
      if S._url == "<nil>" {
        err := fmt.Errorf("URL error: URL is Empty")
        logger.Println(err)
        return Response{}, err
      }
      response, err = S.Session.Post(S._url, &payload)
      if err != nil{
        err = fmt.Errorf("Post request Failed: %v\n",err)
        log.Println(err)
        return Response{},err
      }
    case "PUT":
      response, err = S.Session.Put(S._url, &payload)
      if err != nil{
        err = fmt.Errorf("Put request Failed: %v\n",err)
        log.Println(err)
        return Response{},err
      }
    }
  }

  return NewResponse(response,S.Chunk_size, S.Resource, stream), nil
}

func (S SnowRequest) _get_custom_endpoint(value string) (string){
  var segment string
  if !strings.HasPrefix(value,"/"){
    segment = fmt.Sprintf("/%s", value)
  }
  return S.URLBuilder.get_appended_custom(segment)
}

func (S SnowRequest) update(query interface{}, payload map[string]string) (Response, error) {
  limits, err := S.Parameters.getlimit()
  if err != nil {
    err = fmt.Errorf("Failed to get limit due to: %v",err)
    logger.Println(err)
    return Response{}, err
  }
  offset := S.Parameters.getoffset()
  display_value := S.Parameters.getdisplay_value()
  exclude_reference_link := S.Parameters.getexclude_reference_link()
  suppress_pagination_header := S.Parameters.getsuppress_pagination_header()
  record, err := S.get(query, limits, offset, false, display_value, exclude_reference_link, suppress_pagination_header, nil)
  if err != nil {
    err = fmt.Errorf("Get error: %v", err)
  }
  first_record, err := record.First()
  if err != nil {
    return Response{}, fmt.Errorf("Could not update due to querying error.")
  }
  S._url = S._get_custom_endpoint(first_record["sys_id"].(string))
  return S._get_response("PUT", false, payload)
}
