package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/chai-rs/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Go Interpreter, a simple programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
