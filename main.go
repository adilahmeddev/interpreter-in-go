package main

import (
	"fmt"
	"interpreter-in-go/repl"
	"os"
	osUser "os/user"
)

func main() {
	user, err := osUser.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("hello %s! This is the monkey programming language!\n", user.Username)
	fmt.Printf("feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
