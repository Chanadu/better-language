package parser

import (
	"Better-Language/parser/expressions"
	"Better-Language/parser/statements"
	"Better-Language/scanner/tokentype"
)

func (p *parser) parseDeclaration() (s statements.Statement, ok bool) {
	if p.match(tokentype.Var) {
		s, ok = p.parseVarDeclaration()
	} else {
		s, ok = p.parseStatement()
	}
	if !ok {
		p.synchronize()
	}
	return s, ok
}

func (p *parser) parseVarDeclaration() (s statements.Statement, ok bool) {
	varName, ok := p.consume(tokentype.Identifier, "Expect variable name.")

	var initializer expressions.Expression = nil
	if p.match(tokentype.Equal) {
		initializer = p.parseExpression()
	}
	_, ok = p.consume(tokentype.Semicolon, "Expect ';' after variable declaration.")
	return &statements.Var{
		Name:        varName,
		Initializer: initializer,
	}, ok
}

func (p *parser) parseStatement() (s statements.Statement, ok bool) {
	if p.match(tokentype.Print) {
		return p.parsePrintStatement()
	}
	return p.parseExpressionStatement()
}

func (p *parser) parsePrintStatement() (s statements.Statement, ok bool) {
	expr := p.parseExpression()
	_, ok = p.consume(tokentype.Semicolon, "Expect ';' after value.")
	return &statements.Print{
		Expression: expr,
	}, ok
}

func (p *parser) parseExpressionStatement() (s statements.Statement, ok bool) {
	expr := p.parseExpression()
	_, ok = p.consume(tokentype.Semicolon, "Expect ';' after expression.")
	return &statements.Expression{
		Expression: expr,
	}, ok
}
