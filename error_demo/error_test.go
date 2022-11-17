package error_demo

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Err struct {
	Code int
	Msg  string
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func New(code int, msg string) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}

func TestError(t *testing.T) {
	fmt.Println(New(401, "无此权限"))
}
