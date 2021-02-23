package gosnow

import(
  "regexp"
  "fmt"
  "net/url"
)

type URLBuilder struct {
  Base_url  *url.URL
  Base_path string
  Api_path  string
  Full_path string
  _resource_url string
}

func URLBuilderNew(base_url *url.URL, base_path, api_path string) (U URLBuilder){
  U.Base_url  = base_url
  U.Base_path = base_path
  U.Api_path  = api_path
  U.Full_path = fmt.Sprintf("%s",base_path + api_path)
  U._resource_url = fmt.Sprintf("%s",base_url.String() + U.Full_path)

  return U
}

func validate_path(path string) (bool, error) {
  if match, _ := regexp.MatchString("^/(?:[._a-zA-Z0-9-]/?)+[^/]$",path); !match {
    return false, fmt.Errorf("Path validation failed - Expected: '/<component>[/component], got: %s", path)
  }
  return true,nil
}

func (U URLBuilder) get_appended_custom(path_component string) (string){
  if _, err := validate_path(path_component); err != nil {
    panic(err)
  }

  return U._resource_url + path_component
}

func (U URLBuilder) getURL() string{
  return fmt.Sprintf("%s%s", U.Base_url, U.Full_path)
}
