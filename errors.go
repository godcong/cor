package cor

import "errors"

var (
	NIL_TARGET = errors.New("target could not found")
)

var (
	ERROR_HEADER_FLAG_SET_ERROR = errors.New("could not set header flag")
	ERROR_SERVER_READ_HEADER    = errors.New("server could not read header")
	ERROR_CLIENT_WRITE_HEADER   = errors.New("client could not write header")
	ERROR_BITS_SET_OVERFLOW     = errors.New("bits set overflow")
	ERROR_PARSE_TO_UINT8        = errors.New("parse uint8 error")
	ERROR_PARSE_TO_UINT64       = errors.New("parse uint64 error")
)
