package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)

var builtins = []string{"type", "echo", "exit"}

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
	args := strings.Fields(command)

	switch args[0] {
	case "exit":
		return true
	case "echo":
		executeEcho(command)
	case "type":
		executeType(args)
	default:
		if _, err := exec.LookPath(args[0]); err == nil {
			executeBin(args[0], args[1:])
		} else {
			fmt.Println(command + ": command not found")
		}
	}
	return false
}

func executeEcho(command string) {
	if len(command) <= 5 {
		fmt.Println()
	} else {
		fmt.Println(command[5:])
	}
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
