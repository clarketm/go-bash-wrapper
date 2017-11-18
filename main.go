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
