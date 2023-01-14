package token

import "fmt"

type Type string

type Literal []rune

type Position struct {
	Offset int
	Line   int
	Column int
}

type Token struct {
	Type
	Literal
	Position
}

const (
	INVALID  = "INVALID"
	EOF      = "EOF"
	COMMA    = "COMMA"
	COLON    = "COLON"
	LBRACE   = "LBRACE"
	RBRACE   = "RBRACE"
	LBRACKET = "LBRACKET"
	RBRACKET = "RBRACKET"
	STRING   = "STRING"
	INTEGER  = "INTEGER"
)

func NewToken(typ Type, literal string) Token {
	return Token{Type: typ, Literal: []rune(literal)}
}

func (t *Token) String() string {
	return fmt.Sprintf("{token.%s, '%s'},", t.Type, string(t.Literal))
}
