package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/google/shlex"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)

var builtins = []string{"type", "echo", "exit", "pwd", "cd"}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		//wait for user input
		command, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		if executeCommand(command) {
			break
		}
	}
}

func executeCommand(command string) bool {
	args, _ := shlex.Split(command)

	switch args[0] {
	case "exit":
		return true
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "type":
		executeType(args)
	case "pwd":
		executePwd()
	case "cd":
		executeCd(args)
	default:
		if _, err := exec.LookPath(args[0]); err == nil {
			executeBin(args[0], args[1:])
		} else {
			fmt.Println(command + ": command not found")
		}
	}
	return false
}

func executeType(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: type [COMMAND]")
		return
	}
	cmd := args[1]
	if slices.Contains(builtins, cmd) {
		fmt.Println(cmd + " is a shell builtin")
	} else if path, err := exec.LookPath(cmd); err == nil {
		fmt.Println(cmd + " is " + path)
	} else {
		fmt.Println(cmd + ": not found")
	}
}

func executeBin(command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(cmd.Stderr, "Error running program", err)
	}
}

func executePwd() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting current dir", err)
	}

	fmt.Println(path)
}

func executeCd(args []string) {
	path := args[1]
	path = strings.Replace(path, "~", os.Getenv("HOME"), 1)
	path = filepath.Clean(path)

	if err := os.Chdir(path); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
	}
}
