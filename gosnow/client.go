package gosnow

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/levigross/grequests"
)

const pathFormat string = "^/(?:[._a-zA-Z0-9-]/?)+[^/]$"

// isValidatePath assesses whether the given path is valid
//
// expected format is /<component>[/component]
func isValidatePath(path string) bool {
	if match, _ := regexp.MatchString(pathFormat, path); !match {
		logger.Printf("Path validation failed - Expected: '/<component>[/component], got: %s\n", path)
		return false
	}
	return true
}

// Client used as main client for service-now
type Client struct {
	Username string `validate:"required"`
	Instance string
	Use_ssl  bool
	ready    bool
	Session  *grequests.Session
	BaseURL  *url.URL
}

const sharedBase = "/api"
const sharedBaseNow = sharedBase + "/now"

// New Creates a new Client struct using the provided username, password, and instance
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

	apiPath = sharedBaseNow + apiPath

	if !C.ready {
		err := NewInvalidResource("failed to create resource, empty client.")
		logger.Println(err)
		return Resource{}, err
	}

	for _, path := range []string{apiPath} {
		if !isValidatePath(path) {
			err := NewInvalidResource("invalid web address")
			logger.Println(err)
			return Resource{}, err
		}
	}

	return NewResource(C.BaseURL, apiPath, C.Session, 8192), nil
}

// Table returns a new instance of the Table API
func (C Client) Table(tableName string) (Table, error) {

	apiPath := fmt.Sprintf("%s/table/%s", sharedBaseNow, tableName)

	if !C.ready {
		err := NewInvalidResource("failed to create resource, empty client.")
		logger.Println(err)
		return Table{}, err
	}

	return NewTable(C.BaseURL, apiPath, C.Session, 8192), nil
}

// Attachments returns a new instance of the Attachments API
func (C Client) Attachments() (Attachment, error) {

	ApiPath := sharedBaseNow + "/attachment"

	if !C.ready {
		err := NewInvalidResource("failed to create resource, empty client.")
		logger.Println(err)
		return Attachment{}, err
	}

	return NewAttachment(C.BaseURL, ApiPath, C.Session, 8192), nil
}

// ServiceCatalog returns a new instance of the Service Catalog API
func (C Client) ServiceCatalog(apiPath string) (ServiceCatalog, error) {

	apiPath = sharedBase + "/sn_sc/servicecatalog" + apiPath

	if !C.ready {
		err := errors.New("failed to create service catalog, empty client")
		logger.Println(err)
		return ServiceCatalog{}, err
	}

	for _, path := range []string{apiPath} {
		if !isValidatePath(path) {
			err := errors.New("invalid web address")
			logger.Println(err)
			return ServiceCatalog{}, err
		}
	}

	return NewServiceCatalog(C.BaseURL, apiPath, C.Session, 8192), nil
}

func (C Client) Batch() (Batch, error) {

	apiPath := sharedBase + "/now/v1/batch"

	if !C.ready {
		err := errors.New("failed to create service catalog, empty client")
		logger.Println(err)
		return Batch{}, err
	}

	return NewBatch(C.BaseURL, apiPath, C.Session, 8192), nil
}
