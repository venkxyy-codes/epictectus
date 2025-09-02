package error

import "errors"

type Error string

const (
	RuntimeError       Error = "runtime-error"
	ValidationError    Error = "validation-error"
	SystemError        Error = "system-error"
	PeerError          Error = "peer-app-error"
	ClientError        Error = "client-error"
	SerializationError Error = "serialization-error"
	HTTPError          Error = "http-error"
	NotImplemented     Error = "not-implemented"
	ResourceNotFound   Error = "resource-not-found"
	DatabaseError      Error = "database-error"

	NonCustom Error = "NonCustom"
)

var (
	ErrInvalidRequest = errors.New(string(ClientError))
)
