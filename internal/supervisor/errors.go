package supervisor

import "fmt"

type Error struct {
	CodeError string
	Cause     error
}

func (d Error) Error() string {
	return fmt.Sprintf("error_code:%s - cause:%s", d.CodeError, d.Cause.Error())
}
