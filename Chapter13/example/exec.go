package main

import (
	"fmt"
	"os"
	"os/exec"
)

func exectuate() {
	// 方式 1) os.StartProcess
	env := os.Environ()
	// fmt.Println(env)
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin, os.Stdout, os.Stderr,
		},
	}
	// 1st example: list files
	pid, err := os.StartProcess("C:\\Windows\\System32\\calc.exe", []string{""}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!\n", err)
		return
	}
	fmt.Printf("The process id is %d\n", pid)
	// 方式 2) exec.Run
	cmd := exec.Command("notepad")
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error %v execting command!", err)
		return
	}
	fmt.Printf("The command is %v\n", cmd)

}
