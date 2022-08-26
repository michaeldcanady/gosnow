package gosnow

type Method string

type BatchRestRequest struct {
	Id      string
	Headers []map[string]string
	Url     string
	Method  string
}
