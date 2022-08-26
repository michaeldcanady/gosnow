package gosnow

import (
	"net/url"

	"github.com/levigross/grequests"
)

type Batch Resource

func NewBatch(baseURL *url.URL, apiPath string, session *grequests.Session, chunkSize int) (B Batch) {

	B = Batch(NewResource(baseURL, apiPath, session, 8192))

	return
}
