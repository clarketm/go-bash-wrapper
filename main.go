package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	script := os.Args[1]
	cmd := exec.Command(script, os.Args[2:]...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()
	ch := make(chan string, 100)
	stdoutScan := bufio.NewScanner(stdout)
	stderrScan := bufio.NewScanner(stderr)
	go func() {
		for stdoutScan.Scan() {
			line := stdoutScan.Text()
			ch <- line
		}
	}()
	go func() {
		for stderrScan.Scan() {
			line := stderrScan.Text()
			ch <- line
		}
	}()
	go func() {
		cmd.Wait()
		close(ch)
	}()
	for line := range ch {
		fmt.Println(line)
	}
}
