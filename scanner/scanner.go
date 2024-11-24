package scanner

import (
	"strconv"

	"Better-Language/scanner/token"
	"Better-Language/scanner/tokentype"
	"Better-Language/utils"
)

type Scanner interface {
	ScanTokens() ([]token.Token, error)
	scanToken() (*token.Token, bool, error)
	scanSlashToken() (tokentype.TokenType, bool, error)
	scanStringToken() (interface{}, error)
	scanNumberToken() (tokentype.TokenType, interface{}, error)
	scanIdentifierToken() (tokentype.TokenType, error)
	isAtEnd(offset int) bool
	advanceCurrent() rune
	createToken(tt tokentype.TokenType, literal interface{}) *token.Token
	match(expected rune) bool
	peek(offset int) rune
}

type scanner struct {
	source         string
	tokens         []token.Token
	start, current int
	lineNumber     int
	foundError     bool
}

func (sc *scanner) ScanTokens() ([]token.Token, error) {
	for !sc.isAtEnd(0) {
		sc.start = sc.current

		t, shouldAddToken, err := sc.scanToken()
		if err != nil {
			return nil, err
		}

		if shouldAddToken {
			sc.tokens = append(sc.tokens, *t)
		}
	}

	sc.tokens = append(sc.tokens, token.Token{Type: tokentype.EndOfFile, Lexeme: "", Literal: nil, Line: sc.lineNumber})

	return sc.tokens, nil
}

func (sc *scanner) scanToken() (t *token.Token, shouldAddToken bool, e error) {
	r := sc.advanceCurrent()
	tt := tokentype.Base
	var literal interface{} = nil
	shouldAdd := true
	switch r {
	case '(':
		tt = tokentype.OpeningParentheses
	case ')':
		tt = tokentype.ClosingParentheses
	case '{':
		tt = tokentype.OpeningCurlyBrace
	case '}':
		tt = tokentype.ClosingCurlyBrace
	case ',':
		tt = tokentype.Comma
	case '.':
		tt = tokentype.Dot
	case ';':
		tt = tokentype.Semicolon
	case '-':
		tt = tokentype.Minus
	case '+':
		tt = tokentype.Plus
	case '*':
		tt = tokentype.Star
	case '%':
		tt = tokentype.Percent
	case '/':
		var err error = nil
		tt, shouldAdd, err = sc.scanSlashToken()
		if err != nil {
			return &token.Token{}, false, err
		}
	case '!':
		if sc.match('=') {
			tt = tokentype.NotEqual
		} else {
			tt = tokentype.Not
		}
	case '=':
		if sc.match('=') {
			tt = tokentype.EqualEqual
		} else {
			tt = tokentype.Equal
		}
	case '>':
		if sc.match('=') {
			tt = tokentype.GreaterEqual
		} else if sc.match('>') {
			tt = tokentype.BitwiseShiftRight
		} else {
			tt = tokentype.Greater
		}
	case '<':
		if sc.match('=') {
			tt = tokentype.LessEqual
		} else if sc.match('<') {
			tt = tokentype.BitwiseShiftLeft
		} else {
			tt = tokentype.Less
		}
	case '|':
		if sc.match('|') {
			tt = tokentype.Or
		} else {
			tt = tokentype.BitwiseOr
		}
	case '&':
		if sc.match('&') {
			tt = tokentype.And
		} else {
			tt = tokentype.BitwiseAnd
		}
	case '^':
		tt = tokentype.BitwiseXor
	case '~':
		tt = tokentype.BitwiseNot
	case ' ', '\r', '\t':
		shouldAdd = false
	case '\n':
		sc.lineNumber++
		// shouldAdd = false
		isStatementEnder := sc.scanNewLineToken()
		if isStatementEnder {
			tt = tokentype.Semicolon
		} else {
			shouldAdd = false
		}
	case '"':
		var err error = nil
		tt = tokentype.String
		literal, err = sc.scanStringToken()

		if err != nil {
			return &token.Token{}, false, err
		}
	// case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
	//	tt = tokentype.Integer
	default:
		if utils.IsDigit(r) {
			var err error
			tt, literal, err = sc.scanNumberToken()
			if err != nil {
				return &token.Token{}, false, err
			}
			break
		}
		if utils.IsAlpha(r) {
			var err error
			tt, err = sc.scanIdentifierToken()
			if err != nil {
				return &token.Token{}, false, err
			}
			break
		}

		sc.foundError = true
		err := utils.CreateAndReportScannerErrorf(sc.lineNumber, "Unexpected character: %c", r)
		if err != nil {
			return &token.Token{}, false, err
		}
	}

	return sc.createToken(tt, literal), shouldAdd, nil
}

func (sc *scanner) scanSlashToken() (tt tokentype.TokenType, shouldAddToken bool, e error) {
	if sc.match('/') {
		for !sc.isAtEnd(0) && sc.peek(0) != '\n' {
			_ = sc.advanceCurrent()
		}
		return tokentype.Base, false, nil
	} else {
		return tokentype.Slash, true, nil
	}
}

func (sc *scanner) scanStringToken() (literal interface{}, e error) {
	for sc.peek(0) != '"' && !sc.isAtEnd(0) {
		if sc.peek(0) == '\n' {
			sc.lineNumber++
		}
		_ = sc.advanceCurrent()
	}
	if sc.isAtEnd(0) {
		err := utils.CreateAndReportScannerError(sc.lineNumber, "Unterminated string at EOF.")
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	// Moves the current pointer to the closing quote (bc previously peeked)
	sc.advanceCurrent()

	// Cuts off the quotes
	strValue := sc.source[sc.start+1 : sc.current-1]
	return strValue, nil
}

func (sc *scanner) scanNumberToken() (tt tokentype.TokenType, literal interface{}, e error) {
	tt = tokentype.Integer
	var lit interface{} = nil
	// Integer Part, Number between 0 and 9
	for sc.peek(0) >= '0' && sc.peek(0) <= '9' {
		_ = sc.advanceCurrent()
	}

	// Checks if the number is a double
	if sc.peek(0) == '.' && sc.peek(1) >= '0' && sc.peek(1) <= '9' {
		_ = sc.advanceCurrent()
		tt = tokentype.Double

		// Double Part, Number between 0 and 9
		for sc.peek(0) >= '0' && sc.peek(0) <= '9' {
			_ = sc.advanceCurrent()
		}

		var err error = nil
		lit, err = strconv.ParseFloat(sc.source[sc.start:sc.current], 64)
		if err != nil {
			return tokentype.Base, nil, err
		}

	} else {
		var err error = nil
		lit, err = strconv.ParseInt(sc.source[sc.start:sc.current], 10, 64)
		if err != nil {
			return tokentype.Base, nil, err
		}
	}

	return tt, lit, nil
}

func (sc *scanner) scanIdentifierToken() (tt tokentype.TokenType, e error) {
	for utils.IsAlpha(sc.peek(0)) || utils.IsDigit(sc.peek(0)) {
		_ = sc.advanceCurrent()
	}

	text := sc.source[sc.start:sc.current]

	var ok bool
	if tt, ok = tokentype.KeywordsToTokenType[text]; ok {
		tt = tokentype.KeywordsToTokenType[text]
	} else {
		tt = tokentype.Identifier
	}

	return tt, nil
}

func (sc *scanner) scanNewLineToken() (isStatementEnder bool) {
	prevToken := sc.tokens[len(sc.tokens)-1]
	if _, ok := tokentype.NewLineSemicolonTokens[prevToken.Type]; ok {
		return true
	} else {
		return false
	}
}
