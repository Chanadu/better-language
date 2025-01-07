package parser

import (
	"errors"
	"fmt"

	"Better-Language/globals"
	"Better-Language/parser/expressions"
	"Better-Language/scanner"
	"Better-Language/scanner/tokentype"
)

type Parser interface {
	Parse() (expressions.Expression, error)
}

type parser struct {
	tokens  []scanner.Token
	current int
	err     error
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

// LeftAssociativeBinary -> fn ( (tokens) fn )*
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

// Equality -> Comparison ( ( "!=" | "==" ) Comparison )*
func (p *parser) parseEquality() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseComparison, []tokentype.TokenType{
		tokentype.NotEqual,
		tokentype.EqualEqual,
	})
}

// Comparison -> Bitwise OR ( ( ">" | ">=" | "<" | "<=" ) Bitwise OR)*
func (p *parser) parseComparison() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseBitwiseOR, []tokentype.TokenType{
		tokentype.Greater,
		tokentype.GreaterEqual,
		tokentype.Less,
		tokentype.LessEqual,
	})
}

// Bitwise OR -> Bitwise XOR ( ( "|" ) Bitwise XOR)*
func (p *parser) parseBitwiseOR() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseBitwiseXOR, []tokentype.TokenType{
		tokentype.BitwiseOR,
	})
}

// Bitwise XOR -> Bitwise AND ( ( "^" ) Bitwise AND)*
func (p *parser) parseBitwiseXOR() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseBitwiseAND, []tokentype.TokenType{
		tokentype.BitwiseXOR,
	})
}

// Bitwise AND -> Bitwise Shift ( ( "&" ) Bitwise Shift)*
func (p *parser) parseBitwiseAND() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseBitwiseShift, []tokentype.TokenType{
		tokentype.BitwiseAND,
	})
}

// Bitwise Shift -> Term ( ( "<<" | ">>" ) Term )*
func (p *parser) parseBitwiseShift() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseTerm, []tokentype.TokenType{
		tokentype.BitwiseRightShift,
		tokentype.BitwiseLeftShift,
	})
}

// Term -> Factor ( ( "-" | "+" ) Factor )* ;
func (p *parser) parseTerm() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseFactor, []tokentype.TokenType{
		tokentype.Minus,
		tokentype.Plus,
	})
}

// Factor -> Bitwise ( ( "*" | "/" ) Bitwise )* ;
func (p *parser) parseFactor() expressions.Expression {
	return p.parseLeftAssociativeBinary(p.parseUnary, []tokentype.TokenType{
		tokentype.Star,
		tokentype.Slash,
	})
}

// Unary -> ( "-" | "!" | "~" ) Unary | Primary ;
func (p *parser) parseUnary() expressions.Expression {
	if p.match(tokentype.Minus, tokentype.Not, tokentype.BitwiseNOT) {
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

	p.err = fmt.Errorf("expect expression, found %s", p.peek().Lexeme)
	return nil
}

func (p *parser) Parse() (expressions.Expression, error) {
	if p.tokens == nil {
		globals.HasErrors = true
		return nil, errors.New("no tokens to parse, need to add tokens to parser")
	}
	exp := p.parseExpression()
	if p.err != nil {
		globals.HasErrors = true
	}
	return exp, p.err
}
