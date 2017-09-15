package error

import (
	"fmt"
)

// HTTPError info
type HTTPError struct {
	StatusCode uint
	Message    string
}

// J is JSON data struct
type J map[string]interface{}

func (err *HTTPError) Error() string {
	return fmt.Sprintf("%s", err.Message)
}

// JSON format
func (err *HTTPError) JSON() J {
	return J{
		"message": err.Message,
	}
}
