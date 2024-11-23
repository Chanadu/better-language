package tokentype

type TokenType int

const (
	// Base Nil Character
	Base TokenType = iota

	// Single Character Token

	OpeningParentheses
	ClosingParentheses
	OpeningCurlyBrace
	ClosingCurlyBrace

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

	//Class
	//This
	//Super
)

//go:generate stringer -type=TokenType

var KeywordsToTokenType = map[string]TokenType{
	"or":       Or,
	"and":      And,
	"false":    False,
	"true":     True,
	"if":       If,
	"else":     Else,
	"for":      For,
	"while":    While,
	"break":    Break,
	"return":   Return,
	"function": Function,
	"print":    Print,
	"var":      Var,
	"null":     Null,
	//"class":    Class,
	//"this":     This,
}
