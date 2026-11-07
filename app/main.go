package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println(command + ": command not found")
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
	} else {
		fmt.Println(cmd + ": not found")
	}
}
