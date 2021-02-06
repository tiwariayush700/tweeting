package userError

import "strings"

// Error represents an error.
type Error struct {
	Code        string `json:"code"`
	Message     string `json:"message,omitempty"`
	ActualError error  `json:"error,omitempty"`
}

// ErrorNotFound returns a not found error.
func NewErrorNotFound(err error, messages ...string) Error {
	return Error{Code: ErrorNotFound.Error(), Message: strings.Join(messages, " "), ActualError: err}
}

// ErrorInternal returns an internal error.
func NewErrorInternal(err error, messages ...string) Error {
	return Error{Code: ErrorInternal.Error(), Message: strings.Join(messages, " "), ActualError: err}
}

// ErrorBadRequest returns an bad request error.
func NewErrorBadRequest(err error, messages ...string) Error {
	return Error{Code: ErrorBadRequest.Error(), Message: strings.Join(messages, " "), ActualError: err}
}

// ErrorForbidden returns a forbidden request error.
func NewErrorForbidden(err error, messages ...string) Error {
	return Error{Code: ErrorForbidden.Error(), Message: strings.Join(messages, " "), ActualError: err}
}

// Error implements the error interface.
func (e Error) Error() string {
	return e.Code
}
