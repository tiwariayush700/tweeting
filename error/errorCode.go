package userError

type ErrorCode string

// Errors
const (
	ErrorNotFound      = ErrorCode(`not_found`)
	ErrorInternal      = ErrorCode(`internal`)
	ErrorBadRequest    = ErrorCode(`bad request`)
	ErrorForbidden     = ErrorCode(`forbidden`)
	ErrorTokenExpected = ErrorCode(`token_expected`)
	ErrorTokenInvalid  = ErrorCode(`token_invalid`)
)

// Error implements the error interface.
func (e ErrorCode) Error() string {
	return string(e)
}
