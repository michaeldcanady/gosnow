package gosnow

import (
	"net/url"
	"os"
	"reflect"

	"github.com/levigross/grequests"
)

type Attachment Resource

func NewAttachment(BaseURL *url.URL, ApiPath string, session *grequests.Session, chunkSize int) (A Attachment) {

	A = Attachment(NewResource(BaseURL, ApiPath, session, chunkSize))

	return
}

// String returns the string version of the path <[api/now/component/component]>
func (A Attachment) String() string {
	return Resource(A).String()
}

// Get used to fetch a record
func (A Attachment) Get(query interface{}, limits int, offset int, stream bool, fields ...interface{}) Request {

	return Resource(A).Get(reflect.TypeOf(A), query, limits, offset, stream, fields...)
}

// Delete used to remove a record
func (A Attachment) Delete(query interface{}) Request {
	return Resource(A).Delete(reflect.TypeOf(A), query)
}

func (A Attachment) Upload(fileData, tableName, tableSysId, fileName string) Request {

	args := make(map[string]interface{})
	args["data"] = fileData

	oldPath := A.url.Path

	A.url.Path += "/file"

	parameters := make(map[string]interface{})
	parameters["table_name"] = tableName
	parameters["table_sys_id"] = tableSysId
	parameters["file_name"] = fileName

	A.Parameters.AddCustom(parameters)

	resp := Resource(A).Post(reflect.TypeOf(A), args)

	// reset path to original
	A.url.Path = oldPath

	return resp
}

func (A Attachment) Download(sysId, destinationPath string) (err error) {

	oldPath := A.url.Path

	query := map[string]interface{}{"sys_id": sysId}

	preparedRequest := Resource(A).Get(reflect.TypeOf(A), query, 0, 0, false, nil)

	if err != nil {
		return err
	}

	response, err := preparedRequest.Invoke()

	if err != nil {
		return err
	}

	attachment, err := response.(Response).First()

	if err != nil {
		return err
	}

	downloadLink := attachment.Entry["download_link"].(string)

	request := Resource(A)._request()
	request.url = downloadLink
	resp, err := request.Session.Get(downloadLink, nil)

	if err != nil {
		return err
	}

	downloadPath := destinationPath + "\\" + attachment.Entry["file_name"].(string)

	err = os.WriteFile(downloadPath, resp.Bytes(), 0777)

	if err != nil {
		return err
	}

	// reset path to original
	A.url.Path = oldPath

	return
}
