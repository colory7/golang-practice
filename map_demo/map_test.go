package map_demo

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var m map[string]int
	m = make(map[string]int) //不初始化会报错*/
	// 简写  m := make(map[string]int)
	m["张三"] = 10
	m["李四"] = 20
	fmt.Println(m)
	fmt.Println(m["李四"])
	fmt.Printf("类型：%T\n", m)

	//在声明的时候初始化元素
	m1 := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(m1)
}

func TestMap2(t *testing.T) {
	m := make(map[string]int)
	m["张三"] = 10
	m["李四"] = 20
	value, ok := m["张四"]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("查无此人")
	}
}

func TestMap3(t *testing.T) {
	m := make(map[string]int)
	m["张三"] = 10
	m["李四"] = 20

	//使用for range遍历map
	for key, val := range m {
		fmt.Println(key, val)
	}

	//只想遍历key的时候，可以按下面的写法：
	for key := range m {
		fmt.Println(key)
	}
}

func TestMapDelete(t *testing.T) {
	m := make(map[string]int)
	m["小明"] = 50
	m["张三"] = 10
	m["李四"] = 20
	for key, val := range m {
		fmt.Println(key, val)
	}

	delete(m, "张三")
	fmt.Println("删除后的map:")
	for key, val := range m {
		fmt.Println(key, val)
	}
}

func TestMapSlice(t *testing.T) {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("初始化元素：")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string)
	mapSlice[0]["name"] = "小明"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "TBD云集中心"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func TestMapSlice2(t *testing.T) {
	m := make(map[string][]string)
	fmt.Println(m)
	val := []string{"北京", "上海"}
	m["中国"] = val
	fmt.Println(m)
	m["中国"] = append(m["中国"], "广州", "深圳")
	fmt.Println(m)
}
