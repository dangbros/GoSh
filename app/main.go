package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	fmt.Print("$ ")

	//wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	if command != "" {
		fmt.Printf("%s: command not found", command[:len(command)-1])
	}

}
