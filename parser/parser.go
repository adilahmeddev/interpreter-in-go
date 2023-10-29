package parser

import (
	"fmt"
	"interpreter-in-go/ast"
	"interpreter-in-go/lexer"
	"interpreter-in-go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l,
		errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := ast.LetStatement{
		Token: p.curToken,
	}
	if p.peekToken.Type != token.IDENT {
		p.peekError(token.IDENT)
		return nil
	}
	p.nextToken()
	stmt.Name = &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	if p.peekToken.Type != token.ASSIGN {
		p.peekError(token.ASSIGN)
		return nil
	}
	p.nextToken()
	for p.peekToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	return &stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := ast.ReturnStatement{
		Token: p.curToken,
	}
	if p.curToken.Type != token.RETURN {
		p.peekError(token.RETURN)
		return nil
	}

	p.nextToken()
	for p.peekToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	return &stmt
}
func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

//func (p *Parser) parseExpression() ast.Expression {

//}
