package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

func shell(strCmd string) {
	cmds := strings.Split(strCmd, " | ")

	for _, cmd := range cmds {
		c := strings.Split(cmd, " ")[0]

		switch c {
		case "cd":
			dirChangeCmd(cmd)
		case "pwd":
			pwdCmd()
		case "echo":
			echoCmd(cmd)
		case "kill":
			killPsCmd(cmd)
		case "ps":
			psCmd()
		case `\q`:
			exitCmd()
		default:
			fmt.Printf("unknown command [%s]\n", c)
		}
	}
}

func dirChangeCmd(cmd string) {
	err := os.Chdir(strings.Replace(cmd, "cd ", "", 1))
	if err != nil {
		fmt.Println(err)
	}
}

func pwdCmd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir)
}

func echoCmd(cmd string) {
	fmt.Println(strings.Replace(cmd, "echo ", "", 1))
}

func killPsCmd(cmd string) {
	pid, err := strconv.Atoi(strings.Replace(cmd, "kill ", "", 1))
	if err != nil {
		fmt.Println(err)
	}
	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
	}
	err = proc.Kill()
	if err != nil {
		fmt.Println(err)
	}
}

func psCmd() {
	procs, err := ps.Processes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, proc := range procs {
		fmt.Printf("PID: %v    CMD: %v\n", proc.Pid(), proc.Executable())
	}
}

func exitCmd() {
	fmt.Println("exit")
	os.Exit(0)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		shell(sc.Text())
	}
}
