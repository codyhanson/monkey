package parser

import (
	"github.com/codyhanson/monkey/ast"
	"github.com/codyhanson/monkey/lexer"
	"github.com/codyhanson/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens so cur and peek are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.peekToken = p.curToken
	p.curToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
