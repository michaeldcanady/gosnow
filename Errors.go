package gosnow

import (
	"fmt"
)

//Base Errors
var ()

type CustomError struct {
	err error
	msg string
}

func (err CustomError) Error() string {
	if err.err != nil {
		return fmt.Sprintf("%s: %v", err.err, err.msg)
	}
	return err.msg
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
