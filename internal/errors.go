package internal

import "fmt"

type DevApiError struct {
	CodeError string
	Cause     error
}

func (d DevApiError) Error() string {
	return fmt.Sprintf("error_code:%s - cause:%s", d.CodeError, d.Cause.Error())
}
