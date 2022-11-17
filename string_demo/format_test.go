package string_demo

import (
	"fmt"
	"os"
	"testing"
)

type point struct {
	x, y int
}

func TestFormat(t *testing.T) {
	// Go提供了几种打印格式，用来格式化一般的Go值，例如
	// 下面的%v打印了一个point结构体的对象的值
	p := point{1, 2}
	fmt.Printf("%v", p)
	fmt.Println()

	// 如果所格式化的值是一个结构体对象，那么`%+v`的格式化输出
	// 将包括结构体的成员名称和值
	fmt.Printf("%+v", p)
	fmt.Println()

	// `%#v`格式化输出将输出一个值的Go语法表示方式。
	fmt.Printf("%#v", p)
	fmt.Println()

	// 使用`%T`来输出一个值的数据类型
	fmt.Printf("%T", p)
	fmt.Println()

	// 格式化布尔型变量
	fmt.Printf("%t", true)
	fmt.Println()

	// 有很多的方式可以格式化整型，使用`%d`是一种
	// 标准的以10进制来输出整型的方式
	fmt.Printf("%d", 123)
	fmt.Println()

	// 这种方式输出整型的二进制表示方式
	fmt.Printf("%b", 14)
	fmt.Println()

	// 这里打印出该整型数值所对应的字符
	fmt.Printf("%c", 33)
	fmt.Println()

	// 使用`%x`输出一个值的16进制表示方式
	fmt.Printf("%x", 456)
	fmt.Println()

	// 浮点型数值也有几种格式化方法。最基本的一种是`%f`
	fmt.Printf("%f", 78.9)
	fmt.Println()

	// `%e`和`%E`使用科学计数法来输出整型
	fmt.Printf("%e", 123400000.0)
	fmt.Println()
	fmt.Printf("%E", 123400000.0)
	fmt.Println()

	// 使用`%s`输出基本的字符串
	fmt.Printf("%s", "string")
	fmt.Println()

	// 输出像Go源码中那样带双引号的字符串，需使用`%q`
	fmt.Printf("%q", "string")
	fmt.Println()

	// `%x`以16进制输出字符串，每个字符串的字节用两个字符输出
	fmt.Printf("%x", "hex this")
	fmt.Println()

	// 使用`%p`输出一个指针的值
	fmt.Printf("%p", &p)
	fmt.Println()

	// 当输出数字的时候，经常需要去控制输出的宽度和精度。
	// 可以使用一个位于%后面的数字来控制输出的宽度，默认
	// 情况下输出是右对齐的，左边加上空格
	fmt.Printf("|%6d|%6d|", 12, 345)
	fmt.Println()

	// 你也可以指定浮点数的输出宽度，同时你还可以指定浮点数
	// 的输出精度
	fmt.Printf("|%6.2f|%6.2f|", 1.2, 3.45)
	fmt.Println()

	// To left-justify, use the `-` flag.
	fmt.Printf("|%-6.2f|%-6.2f|", 1.2, 3.45)
	fmt.Println()

	// 你也可以指定输出字符串的宽度来保证它们输出对齐。默认
	// 情况下，输出是右对齐的
	fmt.Printf("|%6s|%6s|", "foo", "b")
	fmt.Println()

	// 为了使用左对齐你可以在宽度之前加上`-`号
	fmt.Printf("|%-6s|%-6s|", "foo", "b")
	fmt.Println()

	// `Printf`函数的输出是输出到命令行`os.Stdout`的，你
	// 可以用`Sprintf`来将格式化后的字符串赋值给一个变量
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// 你也可以使用`Fprintf`来将格式化后的值输出到`io.Writers`
	fmt.Fprintf(os.Stderr, "an %s", "error")
}
