package scanner

import (
	"errors"
	"strconv"

	"Better-Language/scanner/tokentype"
	"Better-Language/utils"
)

type Scanner interface {
	ScanTokens() ([]Token, error)
	scanToken() (*Token, bool)
	scanSlashToken() (tokentype.TokenType, bool, error)
	scanStringToken() (any, error)
	scanNumberToken() (tokentype.TokenType, any, error)
	scanIdentifierToken() (tokentype.TokenType, error)
	isAtEnd(offset int) bool
	advanceCurrent() rune
	createToken(tt tokentype.TokenType, literal any) *Token
	match(expected rune) bool
	peek(offset int) rune
}

type scanner struct {
	source         string
	tokens         []Token
	start, current int
	lineNumber     int
	foundError     bool
}

func NewScanner(source string) Scanner {
	return &scanner{
		source:     source,
		tokens:     []Token{},
		start:      0,
		current:    0,
		lineNumber: 1,
		foundError: false,
	}
}

func (sc *scanner) ScanTokens() ([]Token, error) {
	for !sc.isAtEnd(0) {
		sc.start = sc.current
		t, shouldAddToken := sc.scanToken()

		if shouldAddToken {
			sc.tokens = append(sc.tokens, *t)
		}
	}
	sc.tokens = append(sc.tokens, Token{Type: tokentype.EndOfFile, Lexeme: "", Literal: nil, Line: sc.lineNumber})

	return sc.tokens, nil
}

func (sc *scanner) scanToken() (t *Token, shouldAddToken bool) {
	r := sc.advanceCurrent()
	tt := tokentype.Base
	shouldIncrementLineNumber := false
	var literal any = nil
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
			sc.foundError = true
			utils.CreateAndReportScannerErrorf(sc.lineNumber, "Error scanning slash token: %e", err)
			return &Token{}, false
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
		shouldIncrementLineNumber = true
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
			sc.foundError = true
			utils.CreateAndReportScannerErrorf(sc.lineNumber, "Error scanning string token: %e", err)
			return &Token{}, false
		}
	default:
		if utils.IsDigit(r) {
			var err error
			tt, literal, err = sc.scanNumberToken()
			if err != nil {
				sc.foundError = true
				utils.CreateAndReportScannerErrorf(sc.lineNumber, "Error scanning number token: %e", err)
				return &Token{}, false
			}
			break
		}
		if utils.IsAlpha(r) {
			var err error
			tt, err = sc.scanIdentifierToken()
			if err != nil {
				sc.foundError = true
				utils.CreateAndReportScannerErrorf(sc.lineNumber, "Error scanning identifier token: %e", err)
				return &Token{}, false
			}
			break
		}

		sc.foundError = true
		utils.CreateAndReportScannerErrorf(sc.lineNumber, "Unexpected character: %c", r)
	}

	t = sc.createToken(tt, literal)
	if shouldIncrementLineNumber {
		sc.lineNumber++
	}

	return t, shouldAdd

}

func (sc *scanner) scanSlashToken() (tt tokentype.TokenType, shouldAddToken bool, e error) {
	if sc.match('/') {
		for !sc.isAtEnd(0) && sc.peek(0) != '\n' {
			_ = sc.advanceCurrent()
		}
		return tokentype.Base, false, nil
	}
	return tokentype.Slash, true, nil
}

func (sc *scanner) scanStringToken() (literal any, e error) {
	for sc.peek(0) != '"' && !sc.isAtEnd(0) {
		if sc.peek(0) == '\n' {
			sc.lineNumber++
		}
		_ = sc.advanceCurrent()
	}
	if sc.isAtEnd(0) {
		sc.foundError = true
		return nil, errors.New("unterminated string at EOF")
	}

	// Moves the current pointer to the closing quote (bc previously peeked)
	sc.advanceCurrent()

	// Cuts off the quotes
	strValue := sc.source[sc.start+1 : sc.current-1]
	return strValue, nil
}

func (sc *scanner) scanNumberToken() (tt tokentype.TokenType, literal any, e error) {
	tt = tokentype.Integer
	var lit any = nil
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
			sc.foundError = true
			return tokentype.Base, nil, err
		}

	} else {
		var err error = nil
		lit, err = strconv.ParseInt(sc.source[sc.start:sc.current], 10, 64)
		if err != nil {
			sc.foundError = true
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
