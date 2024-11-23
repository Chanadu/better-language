package token

import (
	"fmt"

	"Better-Language/scanner/tokentype"
)

type Token struct {
	Type    tokentype.TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %d %s %v", t.Type, t.Line, t.Lexeme, t.Literal)
}
