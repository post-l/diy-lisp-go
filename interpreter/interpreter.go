package interpreter

import (
	"github.com/post-l/diy-lisp-go/ast"
	"github.com/post-l/diy-lisp-go/environment"
)

func Interpret(exp string, env *environment.Environment) (res string, err error) {
	var expAst *ast.Ast
	res = exp
	err = nil
	return
}
