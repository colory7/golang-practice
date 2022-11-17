package file_demo

import (
	"fmt"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	f, err := os.Create("haha.txt")
	if err != nil {
		panic(err)
	}
	f.WriteString("this is test file")
	defer f.Close()
}

func TestRead(t *testing.T) {
	f, err := os.Open("haha.txt")
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	fmt.Println(string(buf[:n]))
	defer f.Close()
}

func TestReadWrite(t *testing.T) {
	f, err := os.Create("haha.txt")
	//  f,err:=os.Open("haha.txt")
	if err != nil {
		panic(err)
	}
	//  buf :=make([]byte,1024)
	//  n,_=f.read(buf)
	//  fmt.Println(string(buf[:n]))
	f.Seek(5, 0)
	f.WriteString("中文乱码 not here!")
	defer f.Close()
}
