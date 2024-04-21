package gonsul

import "errors"

var (
	errUnableToCreateHttpRequest  = errors.New("GONSUL_UNABLE_TO_CREATE_HTTP_REQUEST")
	errUnableToCallConsulAPI      = errors.New("GONSUL_UNABLE_TO_CALL_CONSUL_API")
	errUnableToReadConsulResponse = errors.New("GONSUL_UNABLE_TO_READ_CONSUL_RESPONSE")
	errConsulKeyNotFound          = errors.New("GONSUL_CONSUL_KEY_NOT_FOUND")
	errConsulUpdateValueFailed    = errors.New("GONSUL_CONSUL_UPDATE_VALUE_FAILED")
	errUnableToReadInputFile      = errors.New("GONSUL_UNABLE_TO_READ_INPUT_FILE")
	errFileNotFound               = errors.New("GONSUL_FILE_NOT_FOUND")
	errValueNotFound              = errors.New("GONSUL_VALUE_NOT_FOUND")
)
