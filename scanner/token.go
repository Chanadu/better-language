package scanner

import (
	"Better-Language/scanner/tokentype"
)

type Token struct {
	Type    tokentype.TokenType
	Lexeme  string
	Literal any
	Line    int
}
