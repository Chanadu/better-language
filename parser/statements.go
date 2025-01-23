package parser

import (
	"Better-Language/parser/statements"
	"Better-Language/scanner/tokentype"
)

func (p *parser) parseStatement() statements.Statement {
	if p.match(tokentype.Print) {
		return p.parsePrintStatement()
	}
	return p.parseExpressionStatement()
}

func (p *parser) parsePrintStatement() statements.Statement {
	expr := p.parseExpression()
	p.consume(tokentype.Semicolon, "Expect ';' after value.")
	return &statements.Print{
		Expression: expr,
	}
}

func (p *parser) parseExpressionStatement() statements.Statement {
	expr := p.parseExpression()
	p.consume(tokentype.Semicolon, "Expect ';' after expression.")
	return &statements.Expression{
		Expression: expr,
	}
}
