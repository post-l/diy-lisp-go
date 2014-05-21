package main

import (
	"fmt"
	"github.com/post-l/diy-lisp-go/repl"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("[Comming Soon] interpret_file(%v)\n", os.Args[1])
		// fmt.Println(interpret_file(os.Args[1]))
	} else {
		repl.Repl()
	}
}
