package cmd_demo2

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestCmd(t *testing.T) {
	command := exec.Command("ls", "-lh", "/etc/passwd")
	err := command.Run()
	if err != nil {
		log.Fatalf("cmd.run() failed with %s", err)
	}
}

func TestCmd2(t *testing.T) {
	//command := exec.Command("ls", "-lh", "/usr/local/logs/")
	command := exec.Command("/bin/sh", "-c", " go run /Users/mac/save/code/golang/cobra_demo/main.go")
	out, err := command.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.run() failed with %s", err)
	}
	fmt.Println(string(out))
}

func TestCmd3(t *testing.T) {
	//cmd := exec.Command("ls", "-lh", "/usr/local/logs/*.log")
	cmd := exec.Command("/bin/sh", "-c", " go run /Users/mac/save/code/golang/cobra_demo/main.go")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func TestCmd4(t *testing.T) {
	c1 := exec.Command("grep", "root", "/etc/passwd")
	c2 := exec.Command("wc", "-l")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
}

func TestCmd5(t *testing.T) {
	os.Setenv("NAME", "cloud")
	cmd := exec.Command("echo", os.ExpandEnv("$NAME"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s", out)
}

type Result struct {
	output string
	err    error
}

func TestCmd6(t *testing.T) {
	results := make(chan Result, 100)
	ctc, cancelFunc := context.WithCancel(context.TODO())

	go func() {
		cmd := exec.CommandContext(ctc, "/bin/sh", "-c", "echo hello;sleep 2;echo world")
		output, err := cmd.CombinedOutput()
		result := Result{
			output: string(output),
			err:    err,
		}
		results <- result
	}()
	time.Sleep(time.Duration(1) * time.Second)
	cancelFunc()
	result := <-results
	fmt.Println(result.output, result.err.Error())
}
