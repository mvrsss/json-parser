package lexer

import (
	_ "fmt"
	"unicode"

	"github.com/mvrsss/json-parser/token"
)

type Lexer struct {
	input []rune
	start int
	end   int
	char  rune
}

func (l *Lexer) readChar() {
	if l.end >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.end]
	}
	l.start = l.end
	l.end += 1
}

func skipWhitespace(l *Lexer) {
	for {
		switch l.char {
		case ' ', '\n', '\t', '\r':
			l.readChar()
		default:
			return
		}
	}
}

func isInteger(l *Lexer) bool {
	if !unicode.IsDigit(l.char) {
		return false
	}

	for unicode.IsDigit(l.char) {
		l.end += 1
		l.char = l.input[l.end]
	}
	return true
}

func isString(l *Lexer) bool {
	if l.char != '"' {
		return false
	}

	for l.end < len(l.input) {
		l.end += 1
		l.char = l.input[l.end]

		if l.input[l.end] == '"' {
			l.end += 1
			l.char = l.input[l.end]
			return true
		}
	}

	return false
}

func NewLexer(input []byte) *Lexer {
	l := &Lexer{input: []rune(string(input))}
	l.readChar()
	return l
}

func (l *Lexer) NewToken() token.Token {
	var tok token.Token
	skipWhitespace(l)

	switch l.char {
	case ':':
		tok = token.NewToken(token.COLON, string(l.char))
	case ',':
		tok = token.NewToken(token.COMMA, string(l.char))
	case '{':
		tok = token.NewToken(token.LBRACE, string(l.char))
	case '}':
		tok = token.NewToken(token.RBRACE, string(l.char))
	case '[':
		tok = token.NewToken(token.LBRACKET, string(l.char))
	case ']':
		tok = token.NewToken(token.RBRACKET, string(l.char))

	default:
		if isString(l) {
			tok = token.NewToken(token.STRING, string(l.input[l.start:l.end]))
		} else if isInteger(l) {
			tok = token.NewToken(token.INTEGER, string(l.input[l.start:l.end]))
		} else if l.char == rune(0) {
			tok = token.NewToken(token.EOF, "")
		} else {
			tok = token.NewToken(token.INVALID, string(l.char))
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) PeakToken() token.Token {
	start := l.start
	end := l.end
	tok := l.NewToken()

	l.start = start
	l.end = end
	l.char = l.input[l.end-1]
	return tok
}
