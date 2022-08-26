package gosnow

import (
	"fmt"
	"net/url"
	"os"

	"github.com/levigross/grequests"
)

type Attachment Resource

func NewAttachment(BaseURL *url.URL, BasePath string, session *grequests.Session, chunkSize int) (A Attachment) {

	ApiPath := "/attachment"

	A = Attachment(NewResource(BaseURL, BasePath, ApiPath, session, chunkSize))

	return
}

// String returns the string version of the path <[api/now/component/component]>
func (A Attachment) String() string {
	return Resource(A).String()
}

// Get used to fetch a record
func (A Attachment) Get(query interface{}, limits int, offset int, stream bool, fields ...interface{}) (resp Response, err error) {

	resp, err = Resource(A).Get(query, limits, offset, stream, fields...)

	return
}

// Delete used to remove a record
func (A Attachment) Delete(query interface{}) (Response, error) {
	return Resource(A).Delete(query)
}

func (A Attachment) Upload(fileData, tableName, tableSysId, fileName string) (resp Response, err error) {

	args := make(map[string]string)
	args["data"] = fileData

	oldPath := A.url.Path

	A.url.Path += "/file"

	parameters := make(map[string]interface{})
	parameters["table_name"] = tableName
	parameters["table_sys_id"] = tableSysId
	parameters["file_name"] = fileName

	A.Parameters.AddCustom(parameters)

	resp, err = Resource(A).Post(args)

	// reset path to original
	A.url.Path = oldPath

	return
}

func (A Attachment) Download(sysId, destinationPath string) (err error) {

	oldPath := A.url.Path

	query := map[string]interface{}{"sys_id": sysId}

	response, err := Resource(A).Get(query, 0, 0, false, nil)

	if err != nil {
		return err
	}

	attachment, err := response.First()

	if err != nil {
		return err
	}

	fmt.Println(attachment)

	downloadLink := attachment["download_link"].(string)

	request := Resource(A)._request()
	request.url = downloadLink
	resp, err := request.Session.Get(downloadLink, nil)

	if err != nil {
		return err
	}

	downloadPath := destinationPath + "\\" + attachment["file_name"].(string)

	err = os.WriteFile(downloadPath, resp.Bytes(), 0777)

	if err != nil {
		return err
	}

	// reset path to original
	A.url.Path = oldPath

	return
}
