package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/calvincolton/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, and welcome to the Monkey Programming Language!\n", user.Username)
	fmt.Printf("Fell free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
