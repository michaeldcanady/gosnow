package gosnow

import (
	"net/url"

	"github.com/levigross/grequests"
)

//Client used as main client for service-now
type Client struct {
	Username string `validate:"required"`
	Instance string
	Use_ssl  bool
	ready    bool
	Session  *grequests.Session
	BaseURL  *url.URL
}

// Creates a new Client struct using the provided username, password, and instance
func New(username, password, instance string) (C Client, err error) {

	if username == "" {
		err = NewMissingParameter("no username provided.")
		logger.Println(err)
		return C, err
	} else if password == "" {
		err = NewMissingParameter("no password provided.")
		logger.Println(err)
		return C, err
	} else if instance == "" {
		err = NewMissingParameter("no instance provided.")
		logger.Println(err)
		return C, err
	} else {
		C.Username = username
		C.BaseURL = &url.URL{
			Scheme: "https",
			Host:   instance + ".service-now.com",
		}
		C.Session = grequests.NewSession(&grequests.RequestOptions{Auth: []string{username, password}})
		C.ready = true
	}

	return C, nil
}

// Resource is used to create table resources
// Each new table that can be queried needs its own .Resource
func (C Client) Resource(apiPath string) (Resource, error) {
	basePath := "/api/now"

	if !C.ready {
		err := NewInvalidResource("failed to create resource, empty client.")
		logger.Println(err)
		return Resource{}, err
	}

	for _, path := range []string{apiPath, basePath} {
		if !validate_path(path) {
			err := NewInvalidResource("invalid web address")
			logger.Println(err)
			return Resource{}, err
		}
	}

	return NewResource(C.BaseURL, basePath, apiPath, C.Session, 8192), nil
}
