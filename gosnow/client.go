package gosnow

import(
  "net/url"
  "github.com/levigross/grequests"
)

//Client used as main client for service-now
type Client struct {
	Username string
  Password string
  Instance string
  Use_ssl  bool
  Session  *grequests.Session
  BaseURL  *url.URL
}

func New(username, password, instance string) (C Client, err error) {
  err = nil
  C.Username = username
  C.Password = password
  C.BaseURL = &url.URL{
    Scheme: "https",
	  Host: instance+".service-now.com",
  }
  C.Session = C.getSession()
  // hashes password for security purposes
  C.Password = hashPassword(C.Password)

  return C, err
}

func (C Client) getSession() (*grequests.Session){
  session := grequests.NewSession(&grequests.RequestOptions{Auth: []string{C.Username,C.Password}})
  return session
}

func (C Client) Resource(api_path string) (Resource, error) {
  base_path := "/api/now"

  for _, path := range []string{api_path, base_path} {
    if _, err := validate_path(path); err != nil {
      return Resource{}, err
    }
  }

  return NewResource(C.BaseURL, base_path, api_path, C.Session, 8192), nil

}
