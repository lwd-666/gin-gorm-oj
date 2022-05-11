package main

import (
	"bytes"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("go","run","code-user/main.go")
	var out,stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}
	io.WriteString(stdinPipe,"23 11\n")
	cmd.Run()
	//根据测试案例运行。拿到输出结果
	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(out.String())

}
