package httpclient

import "fmt"

type HTTPError struct {
	path   string
	msg    string
	status int
}

func NewHTTPError(path, msg string, status int) HTTPError {
	return HTTPError{
		path:   path,
		msg:    msg,
		status: status,
	}
}

func (h HTTPError) String() string {
	return fmt.Sprintf("error response at %s with status %d: %s", h.path, h.status, h.msg)
}

func (h HTTPError) Error() string {
	return h.String()
}
