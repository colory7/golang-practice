package method_demo

import (
	"fmt"
	"testing"
)

//结构体
type User struct {
	Name  string
	Email string
}

//方法
func (u User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

func TestMethodDefine(t *testing.T) {
	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()
	// 指针类型调用方法
	u2 := User{"go", "go@go.com"}
	u3 := u2
	u3.Notify()

}
