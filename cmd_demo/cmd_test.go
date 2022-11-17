package xxx

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	cmdPath, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cmdPath)
}
