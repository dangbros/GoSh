package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	builtins := []string{"type", "echo", "exit"}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		//wait for user input
		command, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		//evaulate
		if command != "\n" {
			command = strings.TrimSpace(command)
			token := strings.Split(command, " ")

			if command == "exit" {
				break
			} else if token[0] == "type" {
				if len(token) > 1 && slices.Contains(builtins, token[1]) {
					fmt.Println(token[1] + " is a shell builtin")
				} else if len(token) > 1 {
					fmt.Println(token[1] + ": not found")
				} else {
					fmt.Println("Usage: type [COMMAND]")
				}
			} else if strings.HasPrefix(command, "echo ") {
				fmt.Println(command[5:])
			} else {
				fmt.Printf("%s: command not found\n", command)
			}
		}

	}
}
