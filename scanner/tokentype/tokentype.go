package tokentype

type TokenType int

const (
	// Single Character Token

	LeftParentheses TokenType = iota
	RightParentheses
	LeftCurlyBrace
	RightCurlyBrace

	Comma
	Dot
	Semicolon

	Minus
	Plus
	Star
	Percent
	Slash

	// One or Two Character Token

	Not
	NotEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals

	Identifier
	String
	Integer
	Double

	// Keywords

	Or
	And

	False
	True

	If
	Else
	For
	While

	Break
	Return

	Function

	Print
	Var

	Null

	EndOfFile

	// Prolly won't use this stuff

	Class
	This
	Super
)

//go:generate stringer -type=TokenType

var RuneToTokenType = map[rune]TokenType{
	'(': LeftParentheses,
	')': RightParentheses,
	'{': LeftCurlyBrace,
	'}': RightCurlyBrace,
	',': Comma,
	'.': Dot,
	';': Semicolon,
	'-': Minus,
	'+': Plus,
	'*': Star,
	'%': Percent,
}
