package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	script := os.Args[1]
	args := strings.Join(os.Args[2:], " ")
	cmd := exec.Command(script, args)
	// cmd.SysProcAttr = &syscall.SysProcAttr{}
	// cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 0, Gid: 0}
	out, _ := cmd.StdoutPipe()
	scan := bufio.NewScanner(out)
	cmd.Start()
	go func() {
		for scan.Scan() {
			line := scan.Text()
			fmt.Println(line)
		}
	}()
	cmd.Wait()
}
