package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/evanboyle/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming Language!\n", user.Username)
	fmt.Printf("Type a command.\n")
	repl.Start(os.Stdin, os.Stdout)
}
