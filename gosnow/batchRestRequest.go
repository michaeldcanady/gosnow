package gosnow

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type BatchRestRequest struct {
	Id      string
	Headers []map[string]string
	Url     string
	Method  string
}
