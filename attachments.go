package gosnow

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/levigross/grequests"
)

var (
	HASMAGIC = false
)

type Attachment struct {
	resource  Resource
	tableName string
}

func NewAttachment(resource Resource, tableName string) (A Attachment) {
	A.resource = resource
	A.tableName = tableName
	return
}

func (A Attachment) Get(sys_id string, limit int) (Response, error) {
	if sys_id == "" {
		query := map[string]interface{}{"table_name": A.tableName}
		return A.resource.Get(query, 1, 0, true, nil)
	}
	//return A.resource.Get(map[string]interface{}{"table_sys_id": sys_id, "table_name": A.tableName}, limit, 0, true, nil)
	return A.resource.Get(map[string]interface{}{"sys_id": sys_id}, limit, 0, true, nil)
}

func (A Attachment) GetTicket(sys_id string, limit int) (Response, error) {
	if sys_id == "" {
		query := map[string]interface{}{"table_name": A.tableName}
		return A.resource.Get(query, 1, 0, true, nil)
	}
	return A.resource.Get(map[string]interface{}{"table_sys_id": sys_id, "table_name": A.tableName}, limit, 0, true, nil)
}

func (A Attachment) Upload(sys_id, file_path string, multipart bool) (Response, error) {
	var payload grequests.RequestOptions

	payload.Headers = make(map[string]string)

	resource := A.resource
	name := filepath.Base(file_path)
	resource.Parameters.AddCustom(map[string]interface{}{"table_name": A.tableName, "table_sys_id": sys_id, "file_name": name})
	payload.JSON, _ = os.ReadFile(file_path)

	var path_append string

	if multipart {
		payload.Headers["Content-Type"] = "multipart/form-data"
		path_append = "/upload"
	} else {
		//TODO Add ability to read magic from files

		if HASMAGIC {
			//magic.from_file(file_path, mime=True)
		} else {
			payload.Headers["Content-Type"] = ("text/plain")
		}
		path_append = "/file"
	}

	return resource.request("POST", path_append, payload)
}

func (A Attachment) Delete(sys_id string) (Response, error) {
	query := map[string]interface{}{"sys_id": sys_id}
	return A.resource.Delete(query)
}

func (A Attachment) Download(sys_id string, destinationPath string) (Response, error) {
	response, err := A.Get(sys_id, 1)

	if err != nil {
		return Response{}, err
	}

	attachment, err := response.First()

	if err != nil {
		return Response{}, err
	}
	downloadLink := attachment["download_link"].(string)

	request := A.resource._request()
	request._url = downloadLink
	resp, err := request.Session.Get(downloadLink, nil)

	downloadPath := destinationPath + "\\" + attachment["file_name"].(string)

	err = ioutil.WriteFile(downloadPath, resp.Bytes(), 0777)

	return Response{}, nil
}
