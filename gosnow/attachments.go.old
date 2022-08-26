package gosnow

import (
	"os"
	"path/filepath"

	"github.com/levigross/grequests"
)

// HASMAGIC if the attachment has magic
// currently not in use
var (
	HASMAGIC = false
)

// Attachment the ServiceNow Attachments API
type Attachment struct {
	resource  Resource
	TableName string
}

// NewAttachment returns new instance of the attachments API
func NewAttachment(resource Resource, TableName string) (A Attachment) {
	A.resource = resource
	A.TableName = TableName
	return
}

// Get used to query a specific attachment
func (A Attachment) Get(sys_id string, limit int) (Response, error) {
	if sys_id == "" {
		query := map[string]interface{}{"table_name": A.TableName}
		return A.resource.Get(query, 1, 0, true, nil)
	}
	//return A.resource.Get(map[string]interface{}{"table_sys_id": sys_id, "table_name": A.TableName}, limit, 0, true, nil)
	return A.resource.Get(map[string]interface{}{"sys_id": sys_id}, limit, 0, true, nil)
}

func (A Attachment) GetTicket(sys_id string, limit int) (Response, error) {
	if sys_id == "" {
		query := map[string]interface{}{"table_name": A.TableName}
		return A.resource.Get(query, 1, 0, true, nil)
	}
	return A.resource.Get(map[string]interface{}{"table_sys_id": sys_id, "table_name": A.TableName}, limit, 0, true, nil)
}

// Upload new attachment to table
func (A Attachment) Upload(sys_id, file_path string, multipart bool) (Response, error) {
	var payload grequests.RequestOptions

	payload.Headers = make(map[string]string)

	resource := A.resource
	name := filepath.Base(file_path)
	resource.Parameters.AddCustom(map[string]interface{}{"table_name": A.TableName, "table_sys_id": sys_id, "file_name": name})
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

// Delete delete a specific attachment by sys_id
func (A Attachment) Delete(sys_id string) (Response, error) {
	query := map[string]interface{}{"sys_id": sys_id}
	return A.resource.Delete(query)
}

// Download download specified attachment to desintationPath from ServiceNow
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
	request.url = downloadLink
	resp, err := request.Session.Get(downloadLink, nil)

	if err != nil {
		return Response{}, err
	}

	downloadPath := destinationPath + "\\" + attachment["file_name"].(string)

	err = os.WriteFile(downloadPath, resp.Bytes(), 0777)

	if err != nil {
		return Response{}, err
	}

	return Response{}, nil
}
