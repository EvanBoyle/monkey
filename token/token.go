package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGEAL = "ILLEGAL"
	EOF      = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foo, x, y
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
