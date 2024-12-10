package io

type Error interface {
	error
	GetCode() string
	GetMessage() string
}

type ApiError struct {
	Code    string `bson:"code"  json:"code"`
	Message string `bson:"message" json:"message"`
}

func NewApiError(code, message string) *ApiError {
	return &ApiError{
		Code:    code,
		Message: message,
	}
}

func (err *ApiError) Error() string {
	return err.Error()
}

func (err *ApiError) GetCode() string {
	return err.Code
}

func (err *ApiError) GetMessage() string {
	return err.Message
}
