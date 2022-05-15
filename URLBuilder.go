package gosnow

import (
	"fmt"
	"net/url"
	"regexp"
)

const pathFormat string = "^/(?:[._a-zA-Z0-9-]/?)+[^/]$"

type URLBuilder struct {
	BaseURL       *url.URL
	BasePath      string
	ApiPath       string
	Full_path     string
	_resource_url string
}

func URLBuilderNew(BaseURL *url.URL, BasePath, ApiPath string) (U URLBuilder) {
	U.BaseURL = BaseURL
	U.BasePath = BasePath
	U.ApiPath = ApiPath
	U.Full_path = fmt.Sprintf("%s", BasePath+ApiPath)
	U._resource_url = fmt.Sprintf("%s", BaseURL.String()+U.Full_path)

	return U
}

//isValidatePath assesses whether the given path is valid
//
//expected format is /<component>[/component]
func isValidatePath(path string) bool {
	if match, _ := regexp.MatchString(pathFormat, path); !match {
		logger.Printf("Path validation failed - Expected: '/<component>[/component], got: %s\n", path)
		return false
	}
	return true
}

func (U URLBuilder) getAppendedCustom(path_component string) string {
	if !isValidatePath(path_component) {
		return ""
	}

	return U._resource_url + path_component
}

//getURL returns string format of URL
func (U URLBuilder) getURL() string {
	return fmt.Sprintf("%s%s", U.BaseURL, U.Full_path)
}
