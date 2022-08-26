package gosnow

import (
	"net/url"
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
func (A Attachment) Get(query interface{}, limits int, offset int, stream bool, fields ...interface{}) PreparedRequest {

	return Resource(A).Get(reflect.TypeOf(A), query, limits, offset, stream, fields...)
}

// Delete used to remove a record
func (A Attachment) Delete(query interface{}) PreparedRequest {
	return Resource(A).Delete(query)
}

func (A Attachment) Upload(fileData, tableName, tableSysId, fileName string) PreparedRequest {

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

//func (A Attachment) Download(sysId, destinationPath string) (err error) {

//	oldPath := A.url.Path

//	query := map[string]interface{}{"sys_id": sysId}

//	response := Resource(A).Get(query, 0, 0, false, nil)

//	if err != nil {
//		return err
//	}

//	response.Invoke()

//	attachment, err := response.First()

//	if err != nil {
//		return err
//	}

//	fmt.Println(attachment)

//	downloadLink := attachment["download_link"].(string)

//	request := Resource(A)._request()
//	request.url = downloadLink
//	resp, err := request.Session.Get(downloadLink, nil)

//	if err != nil {
//		return err
//	}

//	downloadPath := destinationPath + "\\" + attachment["file_name"].(string)

//	err = os.WriteFile(downloadPath, resp.Bytes(), 0777)

//	if err != nil {
//		return err
//	}

// reset path to original
//	A.url.Path = oldPath

//	return
//}
