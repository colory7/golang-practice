package xxx

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
)

func Test(t *testing.T) {

	ps := exec.Command("ps", "aux")
	grep := exec.Command("grep", "-i", "chrome")
	// 创建一个管道
	r, w := io.Pipe()
	defer r.Close()
	defer w.Close()
	// ps向管道的一端写
	ps.Stdout = w
	// grep从管道的一端读
	grep.Stdin = r
	var buffer bytes.Buffer
	// grep的输出为buffer
	grep.Stdout = &buffer
	_ = ps.Start()
	_ = grep.Start()
	ps.Wait()
	w.Close()
	grep.Wait()
	io.Copy(os.Stdout, &buffer)
	fmt.Println(buffer.String())

}
