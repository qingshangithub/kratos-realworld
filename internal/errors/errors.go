package errors

import (
	"errors"
	"fmt"
)

//创建一个新的Error类型，实现Error()方法，即可定制自己的Error
type HTTPError struct {
	Errors map[string][]string `json:"errors"`
	Code   int                 `json:"-"`
}

func NewHTTPError(code int, field string, detail string) *HTTPError {
	return &HTTPError{
		Code: code,
		Errors: map[string][]string{
			field: {detail},
		},
	}
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError: %d", e.Code)
}

func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}
	return &HTTPError{}
}
