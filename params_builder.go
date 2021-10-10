package gosnow

import(
  "strings"
  "fmt"
  "github.com/levigross/grequests"
  "strconv"
)

type ParamsBuilder struct {
  _custom_params             map[string]interface{}
  _sysparms                  map[string]interface{}
}

func NewParamsBuilder() (P ParamsBuilder){
  P._sysparms = map[string]interface{}{
    "sysparm_query": "",
    "sysparm_limit": 10000,
    "sysparm_offset": 0,
    "sysparm_display_value": false,
    "sysparm_suppress_pagination_header": false,
    "sysparm_exclude_reference_link": false,
    "sysparm_view": "",
    "sysparm_fields": []string{},
  }
  return P
}

func (P ParamsBuilder) stringify_query(query map[string]interface{}) (string){
  var querySlice []string
  for k, v := range query {
    querySlice = append(querySlice,fmt.Sprintf("%v=%v", k, v))
  }
  return strings.Join(querySlice,"^")
}

func (P ParamsBuilder) limit(limit int) {
  if limit != 0 {
    P._sysparms["sysparm_limit"] = limit
  }else{
    P._sysparms["sysparm_limit"] = 10000
  }
}

func (P ParamsBuilder) getlimit() (int, error) {
  if  _, ok := P._sysparms["sysparm_limit"].(int); !ok {
    err := fmt.Errorf("'sysparm_limit' is %t not int like expected",P._sysparms["sysparm_limit"])
    logger.Println(err)
    return 0, err
  }
  return P._sysparms["sysparm_limit"].(int), nil
}

func (P ParamsBuilder) display_value(display_value bool) {
  P._sysparms["sysparm_display_value"] = display_value
}

func (P ParamsBuilder) getdisplay_value() (bool) {
  return P._sysparms["sysparm_display_value"].(bool)
}

func (P ParamsBuilder) exclude_reference_link(exclude_reference_link bool) {
  P._sysparms["sysparm_exclude_reference_link"] = exclude_reference_link
}

func (P ParamsBuilder) getexclude_reference_link() (bool) {
  return P._sysparms["sysparm_exclude_reference_link"].(bool)
}

func (P ParamsBuilder) suppress_pagination_header(suppress_pagination_header bool) {
  P._sysparms["sysparm_suppress_pagination_header"] = suppress_pagination_header
}

func (P ParamsBuilder) getsuppress_pagination_header() (bool) {
  return P._sysparms["sysparm_suppress_pagination_header"].(bool)
}

func (P ParamsBuilder) offset(offset int) {
  P._sysparms["sysparm_offset"] = offset
}

func (P ParamsBuilder) getoffset() (int) {
  return P._sysparms["sysparm_offset"].(int)
}

func (P ParamsBuilder) fields(fields ...interface{}) {
  max := len(fields)-1
  var f string
  for i, field := range fields {
    if i != max {
      f += fmt.Sprintf("%v, ",field)
    }else{
      f+= fmt.Sprintf("%v",field)
    }
  }
  P._sysparms["sysparm_fields"] = f
}

func (P ParamsBuilder) query(query map[string]interface{}) {
  P._sysparms["sysparm_query"] = P.stringify_query(query)
}

func (P ParamsBuilder) as_dict() (*grequests.RequestOptions){
  sysparms := grequests.RequestOptions{
    Params: make(map[string]string),
    }

  for k, v := range P._sysparms {
    switch v.(type){
    case int:
      sysparms.Params[k] = strconv.Itoa(v.(int))
    case string:
      if v.(string) == "<nil>" {
        sysparms.Params[k] = ""
      } else {
        sysparms.Params[k] = v.(string)
      }
    case bool:
      sysparms.Params[k] = fmt.Sprintf("%v",v.(bool))
    case []string:
      if v.([]string) == nil{
        sysparms.Params[k] = ""
      }else{
        sysparms.Params[k] = strings.Join(v.([]string),",")
      }
    default:
      panic("Not valid type")
    }
  }


  for k, v := range P._custom_params {
    switch v.(type){
    case int:
      sysparms.Params[k] = strconv.Itoa(v.(int))
    case string:
      sysparms.Params[k] = v.(string)
    case bool:
      sysparms.Params[k] = fmt.Sprintf("%v",v.(bool))
    default:
      panic("Not valid type")
    }
  }

  return &sysparms
}
