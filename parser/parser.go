package parser

import (
	"fmt"

	"github.com/mvrsss/json-parser/ast"
	"github.com/mvrsss/json-parser/lexer"
	"github.com/mvrsss/json-parser/token"
)

type Parser struct {
	Lexer *lexer.Lexer
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{Lexer: l}
}

func parseArray(p *Parser) ast.Json {
	array := []ast.Json{}
	tok := p.Lexer.PeakToken()

	if tok.Type == token.RBRACKET {
		return &ast.Array{array}
	} else {
		array = append(array, p.Parse())
		tok = p.Lexer.NewToken()
		if tok.Type == token.RBRACKET {
			return &ast.Array{array}
		}
	}

	for {
		array = append(array, p.Parse())
		tok = p.Lexer.NewToken()
		if tok.Type == token.RBRACKET {
			break
		}

		if tok.Type != token.COMMA {
			panic(fmt.Sprintf("was expecting ',' got %s in array parse", string(tok.Literal)))
		}
	}
	return &ast.Array{array}
}

func parseObject(p *Parser) ast.Json {
	object := map[string]ast.Json{}
	tok := p.Lexer.NewToken()

	if tok.Type == token.RBRACE { // nothing inside
		return &ast.Object{object}
	} else {
		key := string(tok.Literal)
		p.Lexer.NewToken() // ':'
		object[key] = p.Parse()
		tok = p.Lexer.NewToken()
		if tok.Type == token.RBRACE {
			return &ast.Object{object}
		}
	}

	for {
		key := string(p.Lexer.NewToken().Literal)
		tok = p.Lexer.NewToken() // ':'
		if tok.Type != token.COLON {
			panic(fmt.Sprintf("was expecting ':' got %s", string(tok.Literal)))
		}
		object[key] = p.Parse()
		tok = p.Lexer.NewToken() // ','

		if tok.Type == token.RBRACE {
			break
		}

		if tok.Type != token.COMMA {
			panic(fmt.Sprintf("was expecting ',' got %s", string(tok.Literal)))
		}
	}
	return &ast.Object{object}
}

func (p *Parser) Parse() ast.Json {
	tok := p.Lexer.NewToken()
	switch tok.Type {
	case token.STRING:
		return &ast.String{string(tok.Literal)}
	case token.INTEGER:
		return &ast.Integer{string(tok.Literal)}
	case token.LBRACE:
		return parseObject(p)
	case token.LBRACKET:
		return parseArray(p)
	}
	return nil
}
