package xendit

import (
	"encoding/json"
)

const (
	// APIValidationErrCode error code for parameters validation
	APIValidationErrCode string = "API_VALIDATION_ERROR"
	// GoErrCode error code for errors happen inside Go code
	GoErrCode string = "GO_ERROR"
)

// Error is the conventional Xendit error
type Error struct {
	Status    int    `json:"status,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
}

// Error returns error message.
// This enables xendit.Error to comply with Go error interface
func (e *Error) Error() string {
	return e.Message
}

// GetErrorCode returns error code coming from xendit backend
func (e *Error) GetErrorCode() string {
	return e.ErrorCode
}

// GetStatus returns http status code
func (e *Error) GetStatus() int {
	return e.Status
}

// FromGoErr generates xendit.Error from generic go errors
func FromGoErr(status int, err error) *Error {
	return &Error{
		Status:    status,
		ErrorCode: GoErrCode,
		Message:   err.Error(),
	}
}

// FromHTTPErr generates xendit.Error from http errors with non 2xx status
func FromHTTPErr(status int, respBody []byte) *Error {
	var httpError *Error
	if err := json.Unmarshal(respBody, &httpError); err != nil {
		return FromGoErr(status, err)
	}
	httpError.Status = status

	return httpError
}
