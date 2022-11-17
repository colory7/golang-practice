package error_demo

import (
	"fmt"
	"testing"
)

type NameEmtpyError struct {
	name string
}

//NameEmtpyError实现了 Error() 方法的对象都可以
func (e *NameEmtpyError) Error() string {
	return "name 不能为空"
}

func NameCheck(name string) (bool, error) {
	if name == "" {
		return false, &NameEmtpyError{name} // 注意error这里必须是地址&引用
	}
	return true, nil
}

func TestError2(t *testing.T) {
	name := ""
	if check, err := NameCheck(name); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(check)
	}
}
