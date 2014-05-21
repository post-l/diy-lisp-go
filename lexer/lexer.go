package lexer

import (
	"strings"
)

const (
	LPAREN = iota
	RPAREN
	FALSE
	TRUE
	INTEGER
	STRING
	PLUS
	MINUS
	MULT
	DIV
	MOD
	MORE
	LESS
	EQ
	IF
	DEFINE
	LAMBDA
	CONS
	HEAD
	TAIL
)

struct Token {

	str string
}

struct Lexer {
	tokens Token[]
	pos int
}

func New(source string) *Lexer {
	source = strings.Replace(source, "(", " ( ", -1)
	source = strings.Replace(source, ")", " ) ", -1)
	s := strings.Split(source, " ")
	tokens := make(Token[], len(s))
	for i := range s {

	}
	return &Lexer{}
}
