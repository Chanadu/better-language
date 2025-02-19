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
	if !ok {
		return nil, false
	}

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
	if p.match(tokentype.OpeningCurlyBrace) {
		stmts, ok := p.parseBlock()
		if !ok {
			return nil, false
		}
		return statements.Block{
			Statements: stmts,
		}, true
	}
	if p.match(tokentype.If) {
		return p.parseIfStatement()
	}
	if p.match(tokentype.While) {
		return p.parseWhileStatement()
	}
	if p.match(tokentype.For) {
		return p.parseForStatement()
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

func (p *parser) parseBlock() (stmts []statements.Statement, ok bool) {
	stmts = make([]statements.Statement, 0)

	for !p.check(tokentype.ClosingCurlyBrace) && !p.isAtEnd() {
		stmt, ok := p.parseDeclaration()
		if !ok {
			return nil, false
		}
		stmts = append(stmts, stmt)
	}

	_, ok = p.consume(tokentype.ClosingCurlyBrace, "Expect '}' after block.")
	if !ok {
		return nil, false
	}
	p.match(tokentype.Semicolon)
	// _, ok = p.consume(tokentype.Semicolon, "Expect ';' after block('}').")
	// if !ok{
	// 	return nil, false
	// }

	return stmts, true
}

func (p *parser) parseIfStatement() (s statements.Statement, ok bool) {
	_, ok = p.consume(tokentype.OpeningParentheses, "Expect '(' after 'if'.")

	if !ok {
		return nil, false
	}

	expr := p.parseExpression()
	p.consume(tokentype.ClosingParentheses, "Expect ')' after if condition.")
	thenBranch, ok := p.parseStatement()
	if !ok {
		return nil, false
	}
	var elseBranch statements.Statement = nil
	if p.match(tokentype.Else) {
		elseBranch, ok = p.parseStatement()
		if !ok {
			return nil, false
		}
	}
	return &statements.If{
		Condition: expr,
		Then:      thenBranch,
		Else:      elseBranch,
	}, true
}

func (p *parser) parseWhileStatement() (s statements.Statement, ok bool) {
	_, ok = p.consume(tokentype.OpeningParentheses, "Expect '(' after 'while'.")
	if !ok {
		return nil, false
	}

	expr := p.parseExpression()
	_, ok = p.consume(tokentype.ClosingParentheses, "Expect ')' after while condition.")
	if !ok {
		return nil, false
	}

	stmt, ok := p.parseStatement()
	if !ok {
		return nil, false
	}

	return &statements.While{
		Condition: expr,
		Body:      stmt,
	}, true

}

func (p *parser) parseForStatement() (s statements.Statement, ok bool) {
	p.consume(tokentype.OpeningParentheses, "Expect '(' after 'for'.")

	var initializer statements.Statement = nil
	if !p.match(tokentype.Semicolon) {
		if p.match(tokentype.Var) {
			initializer, ok = p.parseVarDeclaration()
			if !ok {
				return nil, false
			}
		} else {
			initializer, ok = p.parseExpressionStatement()
			if !ok {
				return nil, false
			}
		}
	}

	var condition expressions.Expression = nil
	if !p.check(tokentype.Semicolon) {
		condition = p.parseExpression()
	}

	p.consume(tokentype.Semicolon, "Expect ';' after loop condition.")

	var increment expressions.Expression = nil
	if !p.check(tokentype.ClosingParentheses) {
		increment = p.parseExpression()
	}
	p.consume(tokentype.ClosingParentheses, "Expect ')' after for clauses.")

	body, ok := p.parseStatement()
	if !ok {
		return nil, false
	}

	if increment != nil {
		body = &statements.Block{
			Statements: []statements.Statement{
				body,
				&statements.Expression{Expression: increment},
			},
		}
	}
	if condition == nil {
		condition = &expressions.Literal{Value: true}
	}
	body = &statements.While{
		Condition: condition,
		Body:      body,
	}

	if initializer != nil {
		body = &statements.Block{
			Statements: []statements.Statement{
				initializer,
				body,
			},
		}
	}

	return body, true
}
