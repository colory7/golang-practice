package xxx

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCmd2(t *testing.T) {
	cmdPath, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}
	command := exec.Command(cmdPath, "-l")
	result, err := command.CombinedOutput()
	if err != nil {
		fmt.Println("命令执行失败：" + err.Error())
	}
	fmt.Println(string(result))

}
