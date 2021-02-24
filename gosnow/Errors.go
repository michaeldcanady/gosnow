package gosnow

import (
	"errors"
	"fmt"
	"strings"
)

//Base Errors
var (
	BaseHashError          = errors.New("Hash Error")
	BaseQueryError         = errors.New("Query Error")
	BaseClientError        = errors.New("Client Error")
	BaseResourceError      = errors.New("Resource Error")
	BaseResponseError      = errors.New("Response Error")
	BaseURLBuilderError    = errors.New("URLBuilder Error")
	BaseSnowRequestError   = errors.New("SnowRequest Error")
	BaseParamsBuilderError = errors.New("ParamsBuilder Error")

	//All client based Errors
	EmptyClient          = CustomError{BaseClientError, "Client is empty"}
	FailedAuthError      = CustomError{BaseClientError, "Failed Authenication"}
	MissUsernameError    = CustomError{BaseClientError, "Missing username"}
	MissPasswordError    = CustomError{BaseClientError, "Missing password"}
	MissInstanceError    = CustomError{BaseClientError, "Missing instance"}
	InvalidInstanceError = CustomError{BaseClientError, "Invalid instance"}

	//All Hash based Errors
	FailedHashError        = CustomError{BaseHashError, "Failed to hash password"}
	IncorrectPasswordError = CustomError{BaseHashError, "The password entered was incorrect"}

	//All Query based errors

	//All Resource based errors

	//All Response based errors

	//All URLBuilder based errors

	//All SnowRequest based errors

	//All ParamsBuilder based errors

)

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

func (err CustomError) Unwrap() error {
	return err.err
}

func (err CustomError) wrap(inner error) error {
	return CustomError{msg: err.msg, err: inner}
}

func (err CustomError) Is(target error) bool {
	ts := target.Error()
	return ts == err.msg || strings.HasPrefix(ts, err.msg+": ")
}
