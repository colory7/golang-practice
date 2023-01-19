package error_demo

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
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

func TestError3(t *testing.T) {
	err := errors.New("")
	fmt.Println(err == nil)

	err2 := make([]interface{}, 0)
	fmt.Println(err2 == nil)

	var err3 interface{}
	fmt.Println(err3 == nil)
	err3 = errors.New("")
	fmt.Println(err3 == nil)

	var err4 error
	fmt.Println(err4 == nil)

	var err5 *error
	fmt.Println(err5 == nil)

	var t1 time.Time
	fmt.Println(t1)

	var t2 *time.Time
	fmt.Println(t2)
	fmt.Println(*t2)
}
