package parser

import (
	"testing"

	"github.com/evanboyle/monkey/ast"
	"github.com/evanboyle/monkey/lexer"
	"github.com/stretchr/testify/assert"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 3, len(program.Statements), "Expected 3 statements, got %d", len(program.Statements))

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testStatement(t *testing.T, s ast.Statement, name string) {
	assert.Equal(t, "let", s.TokenLiteral(), "expected token 'let', got: %s", s.TokenLiteral())

	letStmt, ok := s.(*ast.LetStatement)
	assert.Equal(t, true, ok, "extected *ast.LetStatement, got %T", s)

	assert.Equal(t, name, letStmt.Name.Value, "expected let statement identifier to be %s, got %s", name, letStmt.Name.Value)
	assert.Equal(t, name, letStmt.Name.TokenLiteral(), "expected let stmt token ident to be %s, got %s", name, letStmt.Name.TokenLiteral())
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	assert.NotNil(t, program, "ParseProgram() returned nil")
	assert.Equal(t, 3, len(program.Statements), "Expected 3 statements, got %d", len(program.Statements))

	for _, stmt := range program.Statements {
		assert.IsType(t, &ast.ReturnStatement{}, stmt, "expected type *ast.ReturnStatement but got %T", stmt)
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.errors
	if assert.Equal(t, 0, len(errors), "encountered %d parser errors", len(errors)) {
		return
	}

	for _, e := range errors {
		t.Errorf("parser error: %q", e)
	}
	t.FailNow()
}
