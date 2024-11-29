package parser

import (
	"Better-Language/scanner"
	"Better-Language/scanner/tokentype"
	"Better-Language/utils"
)

func (p *parser) peek() scanner.Token {
	return p.tokens[p.current]
}

func (p *parser) previous() scanner.Token {
	return p.tokens[p.current-1]
}

func (p *parser) isAtEnd() bool {
	return p.peek().Type == tokentype.EndOfFile
}

func (p *parser) advance() scanner.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *parser) check(tokenType tokentype.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *parser) match(tokenTypes ...tokentype.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			_ = p.advance()
			return true
		}
	}
	return false
}

func (p *parser) consume(tokenType tokentype.TokenType, errorMessage string) tokentype.TokenType {
	if p.check(tokenType) {
		return p.advance().Type
	}
	utils.CreateAndReportParsingError(p.peek(), errorMessage)
	return tokentype.Base
}

func (p *parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if _, ok := tokentype.ParseSynchronizationTokens[p.peek().Type]; ok {
			return
		}
		p.advance()
	}
}
