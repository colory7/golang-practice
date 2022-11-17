package switch_demo

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	/*
		switch语句结构
		switch表达式 {
		 case 表达式1，表达式2，...:
				语句块1
		 case 表达式3，表达式4，...:
				语句块2
		 这里可以有无限个case语句
		default:
			语句块
		}
		//案例：
		//	请编写一个程序，该程序可以接收一个字符，比如: a,b,c,d,e,f,g	a 表示星期一，b 表示星期二… 根
		//据用户的输入显示相依的信息.要求使用 switch 语句完成
	*/
	var key byte
	fmt.Println("请输入： a,b,c,d,e,f,g 中的任意一个")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("星期一，猴子穿新衣")
	case 'b':
		fmt.Println("星期二，猴子肚子餓")
	case 'c':
		fmt.Println("星期三，猴子去爬山")
	case 'd':
		fmt.Println("星期四，猴子看電視")
	//....
	default:
		fmt.Println("输入错误！！！")

	}
	fmt.Println("--------------------------------")

	/*
		案例
		case/switch 后是一个表达式( 即：常量值、变量、一个有返回值的函数等都可以)
	*/
	var (
		n1 int64 = 10
		n2 int64 = 10
	) //变量值匹配必须是同类型才能匹配
	switch n1 {
	case n2: //invalid case n2 in switch on n1 (mismatched types int64 and int32)
		//不同类型的不可匹配会报错
		fmt.Println("ok1")
	default:
		fmt.Println("out")
	}
	fmt.Println("--------------------------------")

	//案例:case 后面可以带多个表达式，使用逗号间隔。比如 case 表达式 1, 表达式 2
	var (
		n3 int32 = 15
		n4 int32 = 15
	)
	switch n3 {
	case n4, 10, 5:
		fmt.Println("ok2")
	default:
		fmt.Println("out")
	}

	fmt.Println("--------------------------------")

	//案例：case 后面的表达式如果是常量值(字面量)，则要求不能重复
	var (
		n5 int32 = 15
		n6 int32 = 15
		//n7 int32 = 15

	)
	switch n5 {
	case n6, 10, 5:
		fmt.Println("ok2")
	//case n7,10: //duplicate case 10 in switch 因为上一个case有10这个常量了所以不能重复使用
	//	fmt.Println("ok3")
	default:
		fmt.Println("out")
	}

	fmt.Println("--------------------------------")

	//案例：switch 后也可以不带表达式，类似 if --else 分支来使用
	var age int32 = 10
	switch {
	case age == 10:
		fmt.Println("等于")
	case age == 20:
		fmt.Println("不等于")
	default:
		fmt.Println("没有匹配到")
	}

	fmt.Println("--------------------------------")

	var score int32 = 90
	switch {
	case score >= 90:
		fmt.Println("成绩优良")
	case score <= 80:
		fmt.Println("成绩及格")
	case score <= 60:
		fmt.Println("成绩不合格")
	default:
		fmt.Println("成绩不佳继续努力")
	}

	fmt.Println("--------------------------------")
	//案例：switch 后也可以直接声明/定义一个变量，分号结束，不推荐
	switch score := 80; {
	case score >= 90:
		fmt.Println("成绩优良")
	case score <= 80:
		fmt.Println("成绩及格")
	case score <= 60:
		fmt.Println("成绩不合格")
	default:
		fmt.Println("成绩不佳继续努力")
	}

	fmt.Println("--------------------------------")

	/*
		switch 穿透-fallthrough ，如果在 case 语句块后增加 fallthrough ,则会继续执行下一个 case，也叫 switch 穿透
		一般在 switch 语句中不使用 fallthrough 语句
	*/
	switch score := 100; {
	case score >= 90:
		fmt.Println("成绩优良")
		fallthrough
	case score <= 80:
		fmt.Println("成绩及格")
		fallthrough
	case score <= 60:
		fmt.Println("成绩不合格")
	default:
		fmt.Println("成绩不佳继续努力")
	}

}
