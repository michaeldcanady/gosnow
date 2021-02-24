package gosnow

import (
	"errors"
	"fmt"
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

	if username == "" {
		err = errors.New("No username provided. Connection denied")
		logger.Println(err)
		return C, err
	} else if password == "" {
		err = errors.New("No password provided. Connection denied")
		logger.Println(err)
		return C, err
	} else if instance == "" {
		logger.Println("No password provided. Connection denied")
		logger.Println(err)
		return C, err
	} else {
		C.Username = username
		C.Password = password
		C.BaseURL = &url.URL{
			Scheme: "https",
			Host:   instance + ".service-now.com",
		}
		C.Session = C.getSession()
		// hashes password for security purposes
		C.Password = hashPassword(C.Password)
	}

	return C, nil
}

func (C Client) getSession() *grequests.Session {
	session := grequests.NewSession(&grequests.RequestOptions{Auth: []string{C.Username, C.Password}})
	return session
}

// Resource is used to create table resources
// Each new table that can be queried needs its own .Resource

func (C Client) Resource(apiPath string) (Resource, error) {
	var clientCheck Client
	if C == clientCheck {
		err := errors.New("Failed to created resource. Error: Client is nil")
		logger.Println(err)
		return Resource{}, err
	}
	basePath := "/api/now"

	for _, path := range []string{apiPath, basePath} {
		if !validate_path(path) {
			err := fmt.Errorf("Invalid web address: %s", path)
			logger.Println(err)
			return Resource{}, err
		}
	}

	return NewResource(C.BaseURL, basePath, apiPath, C.Session, 8192), nil

}
