package main

import (
	"Monkey/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Output: \n")
	filePath := "./test.monkey"
	repl.Start(filePath, os.Stdout)
}
