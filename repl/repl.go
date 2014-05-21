package repl

import (
	"bufio"
	"fmt"
	"github.com/post-l/diy-lisp-go/environment"
	"github.com/post-l/diy-lisp-go/interpreter"
	"io"
	"os"
	"strings"
)

func Repl() {
	fmt.Println("Welcome to the DIY-lisp REPL")
	fmt.Println("Use ^D to exit")
	env := environment.New()
	// interpreter.InterpretFile("stdlib.diy", env)
	for true {
		exp, err := readExpression()
		if err == nil {
			fmt.Printf("[Expression] %v\n", exp)
			res, err := interpreter.Interpret(exp, env)
			if err == nil {
				fmt.Println(res)
			} else {
				fmt.Printf("[Error] %v", err)
			}
		} else if err == io.EOF {
			os.Exit(0)
		} else {
			fmt.Printf("[Error] %v\n", err)
			os.Exit(1)
		}
	}
}

func readExpression() (exp string, err error) {
	line := ""
	openParens := 0
	for exp == "" && openParens <= 0 {
		line, err = readLine("→  ")
		if err != nil {
			return
		}
		line = strings.TrimSpace(line)
		if scIndex := strings.IndexByte(line, ';'); scIndex != -1 {
			line = line[:scIndex]
		}
		exp += line + "\n"
		openParens += strings.Count(line, "(") - strings.Count(line, ")")
	}
	return
}

func readLine(prompt string) (line string, err error) {
	var buf []byte
	stdin := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	for isPrefix := true; isPrefix && err == nil; {
		buf, isPrefix, err = stdin.ReadLine()
		line += string(buf)
	}
	return
}
