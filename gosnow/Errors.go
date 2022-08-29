package gosnow

import (
	"fmt"
	"net/http"
)

type ServiceCatalogError struct {
	msg string
	err string
}

func (err ServiceCatalogError) Error() string {
	return fmt.Sprintf("%s: %v", err.err, err.msg)
}

type MissingClientID struct {
}

type MissingParameter struct {
	msg string
	err string
}

func (err MissingParameter) Error() string {
	return fmt.Sprintf("%s: %v", err.err, err.msg)
}

func NewMissingParameter(msg string) error {
	return &MissingParameter{
		err: "Missing Parameter",
		msg: msg,
	}
}

type InvalidResource struct {
	msg string
	err string
}

func (err InvalidResource) Error() string {
	return fmt.Sprintf("%s: %v", err.err, err.msg)
}

func NewInvalidResource(msg string) error {
	return &MissingParameter{
		err: "Invalid Resource",
		msg: msg,
	}
}

type ReponseError struct {
	msg string
	err string
}

func (err ReponseError) Error() string {
	return fmt.Sprintf("%s: %v", err.err, err.msg)
}

type StatusError struct {
	Code    int
	Text    string
	Message string
}

func newStatusError(statusCode int, statusMessage string) *StatusError {
	return &StatusError{
		Code:    statusCode,
		Text:    http.StatusText(statusCode),
		Message: statusMessage,
	}
}

func (m *StatusError) Error() string {
	return fmt.Sprintf("Status Code: %v, Status Text: %v, Status Error: %v", m.Code, m.Text, m.Message)
}
