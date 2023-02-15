package generic_demo

import (
	"fmt"
	"testing"
)

// 声明范型，范型是接口类型
type NumStr interface {
	Num | Str
}

// Num的子实现类是int的子类型或int32的子类型或uint64的子类型
type Num interface {
	~int | ~int32 | ~uint64
}
type Str interface {
	string
}

// 范型函数
func add[T NumStr](a, b T) T {
	return a + b
}

// !int 所有子类
// ｜ 并集
func TestConstraintInf(t *testing.T) {
	fmt.Println(add(3, 4))
	fmt.Println(add("dudu", "yiyi"))
}
