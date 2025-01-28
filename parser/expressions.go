package parser

import (
	"fmt"

	"Better-Language/parser/expressions"
	"Better-Language/scanner/tokentype"
)

func (p *parser) parseExpression() expressions.Expression {
	return p.parseTernary()
}

// Ternary -> Equality ( "?" Expression ":" Expression )?
func (p *parser) parseTernary() expressions.Expression {
	condition := p.parseEquality()
	if p.match(tokentype.QuestionMark) {
		trueBranch := p.parseExpression()
		if _, ok := p.consume(tokentype.Colon, "expected ':' after ternary"); ok {
			falseBranch := p.parseExpression()
			return &expressions.Ternary{
				LineNumber:  p.previous().Line,
				Condition:   condition,
				TrueBranch:  trueBranch,
				FalseBranch: falseBranch,
			}
		}
	}
	return condition
}

type parseFunc func() expressions.Expression

// LeftAssociativeBinary -> fn ( (tokens) fn )*
func (p *parser) parseLeftAssociativeBinary(fn parseFunc, tokens []tokentype.TokenType) expressions.Expression {
	left := fn()
	if p.match(tokens...) {
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
		tokentype.Percent,
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

	if p.match(tokentype.Identifier) {
		return &expressions.Variable{
			Name: p.previous(),
		}
	}

	if p.match(tokentype.OpeningParentheses) {
		expression := p.parseExpression()
		_, _ = p.consume(tokentype.ClosingParentheses, "Expect ')' after expression.")
		return &expressions.Grouping{
			InternalExpression: expression,
		}
	}

	p.err = fmt.Errorf("expect expression, found '%s'", p.peek().Lexeme)
	return nil
}
