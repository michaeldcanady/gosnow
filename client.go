package gosnow

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Client struct {
	Credential UsernamePasswordCredential
	BaseUrl    string
	Session    http.Client
	Decoder    json.Decoder
	Encoder    json.Encoder
}

func NewClient(credential UsernamePasswordCredential, baseUrl string) *Client {

	if !strings.HasSuffix(baseUrl, ".service-now.com") {
		baseUrl += ".service-now.com"
	}

	return &Client{
		Credential: credential,
		BaseUrl:    baseUrl,
		Session:    http.Client{},
	}
}

func (C *Client) Now() *NowRequestBuilder {
	return NewNowRequestBuilder(C.BaseUrl+"/now", C)
}

func (C *Client) Get(url string, target interface{}) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", C.Credential.GetAuthentication())
	resp, err := C.Session.Do(request)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
