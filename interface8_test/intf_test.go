package interface8_test

import "testing"

func Test1(t *testing.T) {
	//定义一个interface类型的变量
	var inter interface{}
	//赋值
	inter = 1
	//定义一个int64的变量
	var i int64
	//将interface类型的inter转为int64
	i = inter.(int64)
	//打印
	fmt.Println(i)
}

func Test2(t *testing.T) {
	//定义一个interface类型的变量
	var inter interface{}
	//赋值
	inter = "1"
	//定义一个string的变量
	var str string
	//将interface类型的inter转为string
	str = inter.(string)
	//打印
	fmt.Println(str)
}

type Person struct {
	Name string
	Age  string
}

func Test3(t *testing.T) {
	//定义一个interface类型的变量
	var inter interface{}
	//赋值
	inter = Person{"student", "18"}
	//定义一个person类型的p
	var p Person
	//将类型为interface的inter转为person类型
	p = inter.(Person)
	//打印
	fmt.Println(p)
}

func Test4(t *testing.T) {
	//定义一个interface类型的变量
	var str string
	//赋值
	str = "1"
	//定义一个int类型
	var i int
	//使用 数据转换包strconv
	//string 转 int
	i, _ = strconv.Atoi(str)
	fmt.Printf("i=%d\n", i)
	//int 转 字符串
	str02 := strconv.Itoa(i)
	fmt.Printf("str02=%s", str02)
}
