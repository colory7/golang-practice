package slice_demo

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	//value := [数据长度]类型 {}

	arr := [1]string{"1"} // 声明并且赋值

	arr = [1]string{} // 声明未赋值
	arr[0] = "1"

}

func TestSliceCreate(t *testing.T) {
	slice_aa := make([]int, 5)
	fmt.Println(slice_aa)

	slice_bb := make([]int, 3, 5)
	fmt.Println(slice_bb)

	//myNum := make([]int, 5, 3)
	//fmt.Println(myNum)

	myStr := []string{"jack", "mark", "nick"}
	fmt.Println(myStr)

	myNum2 := []int{10, 20, 30, 60}
	fmt.Println(myNum2)

	myStr2 := []string{99: ""}
	fmt.Println(myStr2)
}

func TestSliceCompareToArray(t *testing.T) {
	myArray := [3]int{10, 20, 30}
	fmt.Println(myArray)

	mySlice := []int{10, 20, 30}
	fmt.Println(mySlice)
}

func TestSliceNil(t *testing.T) {
	var myNum []int
	fmt.Println(myNum)

	myNum2 := make([]int, 0)
	fmt.Println(myNum2)

	myNum3 := []int{}
	fmt.Println(myNum3)

}

func TestSliceAsign(t *testing.T) {
	myNum := []int{10, 20, 30}
	myNum[1] = 25

	fmt.Println(myNum)
}
func TestSliceConvert(t *testing.T) {
	myNum := []int{10, 20, 30, 40, 50}
	newNum := myNum[1:3]

	fmt.Println(myNum)
	fmt.Println(newNum)

	fmt.Println("===========================")
	newNum[1] = 35

	fmt.Println(myNum)
	fmt.Println(newNum)
	fmt.Println(len(newNum))
	fmt.Println(cap(newNum))
	//fmt.Println(newNum[3])
}

func TestSliceAppend(t *testing.T) {
	myNum := []int{10, 20, 30, 40, 50}
	// 创建新的切片，其长度为 2 个元素，容量为 4 个元素
	newNum := myNum[1:3]
	// 使用原有的容量来分配一个新元素
	// 将新元素赋值为 60
	fmt.Println(myNum)
	fmt.Println(newNum)

	newNum = append(newNum, 60)

	fmt.Println(myNum)
	fmt.Println(newNum)

}

func TestSliceAppend2(t *testing.T) {
	// 创建两个切片，并分别用两个整数进行初始化
	num1 := []int{1, 2}
	num2 := []int{3, 4}
	// 将两个切片追加在一起，并显示结果
	fmt.Printf("%v\n", append(num1, num2...))
}

func TestSliceLimitCap(t *testing.T) {
	// 创建长度和容量都是 5 的字符串切片
	fruit := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 将第三个元素切片，并限制容量
	// 其长度为 1 个元素，容量为 2 个元素
	myFruit := fruit[2:3:4]

	fmt.Println(fruit)
	fmt.Println(myFruit)
	fmt.Println(len(fruit))
	fmt.Println(cap(fruit))
	fmt.Println(len(myFruit))
	fmt.Println(cap(myFruit))

}

func TestSliceLimitCap2(t *testing.T) {
	fruit := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	myFruit := fruit[2:3:4]
	//myFruit := fruit[2:3:5]
	// 向 myFruit 追加新字符串
	myFruit = append(myFruit, "Kiwi")

	fmt.Println(fruit)
	fmt.Println(myFruit)
	fmt.Println("===================")
	fmt.Println(len(fruit))
	fmt.Println(cap(fruit))
	fmt.Println(len(myFruit))
	fmt.Println(cap(myFruit))

}

func TestSliceErgodic(t *testing.T) {
	myNum := []int{10, 20, 30, 40, 50}
	// 迭代每一个元素，并显示其值
	for index, value := range myNum {
		fmt.Printf("index: %d value: %d\n", index, value)
	}
}

func TestSliceErgodic2(t *testing.T) {
	myNum := []int{10, 20, 30, 40, 50}
	// 修改切片元素的值
	// 使用空白标识符(下划线)来忽略原始值
	for index, _ := range myNum {
		myNum[index] += 1
	}
	for index, value := range myNum {
		fmt.Printf("index: %d value: %d\n", index, value)
	}
}

func TestSliceCopy(t *testing.T) {
	num1 := []int{10, 20, 30}
	num2 := make([]int, 5)
	count := copy(num2, num1)

	fmt.Println(count)
	fmt.Println(num2)
}

func TestSliceFunction(t *testing.T) {
	// 1*10^6=100w
	myNum := make([]int, 1e6)
	// 将 myNum 传递到函数 foo()
	slice := foo(myNum)
	// 函数 foo() 接收一个整型切片，并返回这个切片

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func foo(slice []int) []int {
	return slice
}
