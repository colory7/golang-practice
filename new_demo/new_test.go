package new_demo

import (
	"fmt"
	"testing"
)

// new 会自动用 zeroed value 来初始化 类型，
// 也就是字串会是""，number 会是 0，channel, func, map, slice 等等则会是 nil
func TestNewMap(t *testing.T) {
	people := new(map[string]string)
	p := *people
	p["name"] = "Kalan" // panic: assignment to entry in nil map
}

func TestNewSlice(t *testing.T) {
	var a *[]int
	fmt.Printf("a: %p %#v \n", &a, a) //a: 0xc042004028 (*[]int)(nil)
	av := new([]int)
	fmt.Printf("av: %p %#v \n", &av, av) //av: 0xc000074018 &[]int(nil)
	(*av)[0] = 8
	fmt.Printf("av: %p %#v \n", &av, av) //panic: runtime error: index out of range

}

func TestNewSlice2(t *testing.T) {
	// new初始化的slice为指针类型
	list := new([]int)
	// append函数不能接收指针类型
	//list = append(list, 1)
	*list = append(*list, 1)
	fmt.Println(*list)
}

func TestNewChannel(t *testing.T) {
	cv := new(chan string)
	fmt.Printf("cv: %p %#v \n", &cv, cv) //cv: 0xc000074018 (*chan string)(0xc000074020)
	//会报 invalid operation: cv <- "good" (send to non-chan type *chan string)
	// cv <- "good"
}
