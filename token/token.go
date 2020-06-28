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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if v, ok := keywords[ident]; ok {
		return v
	}
	return IDENT
}
