package gosnow

import(
  "github.com/levigross/grequests"
  //"fmt"
  //"strings"
)

type Response struct {
  _response   *grequests.Response
  _chunk_size int
  _count      int
  _resource   Resource
  _stream     bool
}

func NewResponse(response *grequests.Response, chunk_size int, resource Resource, stream bool) (R Response){
  if chunk_size == 0{
    chunk_size = 8192
  }
  R._response = response
  R._chunk_size = chunk_size
  R._count = 0
  R._resource = resource
  R._stream = stream

  return R
}

type Test interface{}

func (R Response) _get_buffered_response() ([]map[string]interface{}, int) {
  response := R._get_response()
  if response.StatusCode == 204 {
    deleted := map[string]interface{}{"status": "record deleted"}
    return []map[string]interface{}{deleted}, 1
  }

  var dT = make(map[string]interface{})

  err := R._response.JSON(&dT)
  if err != nil {
    panic("response error "+err.Error())
  }

  var returnValue = make([]map[string]interface{}, 0)

  for _, r := range dT{
    for _, r := range r.([]interface{}) {
      returnValue = append(returnValue, r.(map[string]interface{}))
    }
  }
  return returnValue, 0
}

func (R Response) _get_response() (*grequests.Response) {
  response := R._response
  //fmt.Printf("%T",response)
  //log.Fatal(response)

  if response.StatusCode == 202 {
    panic("NO Return value")
  }

  return response
}

func (R Response) First() (map[string]interface{}) {
  content, _ := R.All()
  return content[0]
}

func (R Response) Created() (map[string]interface{}, int) {
  response := R._get_response()
  if response.StatusCode == 204 {
    deleted := map[string]interface{}{"status": "record deleted"}
    return deleted, 1
  }

  var dT = make(map[string]interface{})

  err := R._response.JSON(&dT)
  if err != nil {
    panic("response error "+err.Error())
  }

  return dT, 0
}

func (R Response) All() ([]map[string]interface{}, int) {
  return R._get_buffered_response()
}
