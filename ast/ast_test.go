package ast

import (
	"testing"

	"github.com/evanboyle/monkey/token"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	progStr := `let myVar = anotherVar;`

	assert.Equal(t, progStr, program.String(), "error in program serialization")
}
