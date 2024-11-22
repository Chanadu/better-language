package scanner

import "Better-Language/scanner/tokentype"

type Scanner struct {
	source         string
	tokens         []Token
	start, current int
	lineNumber     int
	foundError     bool
}

func (sc *Scanner) scanTokens() ([]Token, error) {
	for sc.current < len(sc.source) {
		sc.start = sc.current

		token, err := sc.scanToken()
		if err != nil {
			return nil, err
		}

		sc.tokens = append(sc.tokens, token)
	}

	sc.tokens = append(sc.tokens, Token{Type: tokentype.EndOfFile, Lexeme: "", Literal: nil, Line: sc.lineNumber})

	return sc.tokens, nil
}

func (sc *Scanner) scanToken() (Token, error) {
	r := sc.advanceCurrent()
	tt, ok := tokentype.RuneToTokenType[r]
	if ok {
		sc.addToken(tt, nil)
	}
	return Token{}, nil
}

func (sc *Scanner) advanceCurrent() (r rune) {
	r = rune(sc.source[sc.current])
	sc.current++
	return r
}

func (sc *Scanner) addToken(tt tokentype.TokenType, literal interface{}) {
	sc.tokens = append(sc.tokens,
		Token{
			Type:    tt,
			Lexeme:  sc.source[sc.start:sc.current],
			Literal: literal,
			Line:    sc.lineNumber,
		},
	)
}
