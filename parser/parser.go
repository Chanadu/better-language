package parser

import (
	"errors"

	"Better-Language/globals"
	"Better-Language/parser/expressions"
	"Better-Language/scanner"
	"Better-Language/scanner/tokentype"
)

type Parser interface {
	Parse() (expressions.Expression, error)
	parseExpression() expressions.Expression
	parseLeftAssociativeBinary(parseFunc, []tokentype.TokenType) expressions.Expression
	parseEquality() expressions.Expression
	parseComparison() expressions.Expression
	parseTerm() expressions.Expression
	parseFactor() expressions.Expression
	parseUnary() expressions.Expression
	parsePrimary() expressions.Expression
	peek() scanner.Token
	previous() scanner.Token
	advance() scanner.Token
	check(tokentype.TokenType) bool
	match(...tokentype.TokenType) bool
	consume(tokentype.TokenType, string) tokentype.TokenType
	isAtEnd() bool
}

type parser struct {
	tokens  []scanner.Token
	current int
}

func NewParser(tokenSlice []scanner.Token) Parser {
	return &parser{
		tokens:  tokenSlice,
		current: 0,
	}
}

func (p *parser) parseExpression() expressions.Expression {
	return p.parseEquality()
}

type parseFunc func() expressions.Expression

// LeftAssociativeBinary -> fn ( (tokens) fn )* ;
func (p *parser) parseLeftAssociativeBinary(fn parseFunc, tokens []tokentype.TokenType) expressions.Expression {
	left := fn()
	for p.match(tokens...) {
		operator := p.previous()
		right := fn()

		left = &expressions.Binary{
			Left:     left,
			Operator: operator,
			Right:    right,
		}
	}
	return left
}

// Equality -> Comparison ( ( "!=" | "==" ) Comparison )* ;
func (p *parser) parseEquality() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseComparison, []tokentype.TokenType{
		tokentype.NotEqual,
		tokentype.EqualEqual,
	})
}

// Comparison -> Term ( ( ">" | ">=" | "<" | "<=" ) Term )* ;
func (p *parser) parseComparison() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseTerm, []tokentype.TokenType{
		tokentype.Greater,
		tokentype.GreaterEqual,
		tokentype.Less,
		tokentype.LessEqual,
	})
}

// Term -> Factor ( ( "-" | "+" ) Factor )* ;
func (p *parser) parseTerm() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseFactor, []tokentype.TokenType{
		tokentype.Minus,
		tokentype.Plus,
	})
}

// Factor -> Unary ( ( "*" | "/" ) Unary )* ;
func (p *parser) parseFactor() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseUnary, []tokentype.TokenType{
		tokentype.Star,
		tokentype.Slash,
	})
}

// Unary -> ( "-" | "!" ) Unary | Primary ;
func (p *parser) parseUnary() expressions.Expression {
	if p.match(tokentype.Minus, tokentype.Not) {
		operator := p.previous()
		right := p.parseUnary()

		return &expressions.Unary{
			Operator: operator,
			Right:    right,
		}
	}
	return p.parsePrimary()
}

// Primary -> Integer | Double | String | True | False | "(" Expression ")" | Null;
func (p *parser) parsePrimary() expressions.Expression {
	if p.match(tokentype.True) {
		return &expressions.Literal{
			Value: true,
		}
	}
	if p.match(tokentype.False) {
		return &expressions.Literal{
			Value: false,
		}
	}
	if p.match(tokentype.Integer, tokentype.Double, tokentype.String) {
		return &expressions.Literal{
			Value: p.previous().Literal,
		}
	}
	if p.match(tokentype.Null) {
		return &expressions.Literal{
			Value: nil,
		}
	}

	if p.match(tokentype.OpeningParentheses) {
		expression := p.parseExpression()
		p.consume(tokentype.ClosingParentheses, "Expect ')' after expression.")
		return &expressions.Grouping{
			InternalExpression: expression,
		}
	}

	return nil
}

func (p *parser) Parse() (expressions.Expression, error) {
	if p.tokens == nil {
		globals.HasErrors = true
		return nil, errors.New("no tokens to parse, need to add tokens to parser")
	}
	return p.parseExpression(), nil
}
