package lexer

import (
	"testing"

	"github.com/evanboyle/monkey/token"
	"github.com/stretchr/testify/assert"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()

		assert.Equal(t, tt.expectedLiteral, tok.Literal, "unexpected token literal")
		assert.Equal(t, tt.expectedType, tok.Type, "unexpected token type")
	}
}
