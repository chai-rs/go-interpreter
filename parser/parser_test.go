package parser_test

import (
	"testing"

	"github.com/chai-rs/go-interpreter/ast"
	"github.com/chai-rs/go-interpreter/lexer"
	"github.com/chai-rs/go-interpreter/parser"
	"github.com/stretchr/testify/assert"
)

func TestParser_LetStatement(t *testing.T) {
	type Testcase struct {
		In       string
		Expected []string
	}

	testcases := []Testcase{
		{
			In: `
			let x = 5;
			let y = 10;
			let foobar = 838383;
			`,
			Expected: []string{"x", "y", "foobar"},
		},
	}

	for _, tc := range testcases {
		l := lexer.New(tc.In)
		p := parser.New(l)
		program := p.ParseProgram()

		assert.Equal(t, len(tc.Expected), len(program.Statements))
		for i, s := range program.Statements {
			assert.Equal(t, tc.Expected[i], s.(*ast.LetStatement).Name.Value)
		}
	}
}

func TestParser_CheckErrors(t *testing.T) {
	type Testcase struct {
		In       string
		Expected []string
	}

	testcases := []Testcase{
		{
			In: `
			let x 5;
			let = 10;
			let 838383;
			`,
			Expected: []string{
				"expected next token to be =, got INT instead",
				"expected next token to be IDENT, got = instead",
				"expected next token to be IDENT, got INT instead",
			},
		},
	}

	for _, tc := range testcases {
		name := tc.In
		l := lexer.New(tc.In)
		parser := parser.New(l)
		parser.ParseProgram()

		t.Run(name, func(t *testing.T) {
			errors := parser.Errors()
			assert.Len(t, errors, len(tc.Expected))
			assert.Equal(t, tc.Expected, errors)
		})
	}

}

func TestParser_ReturnStatement(t *testing.T) {
	type Testcase struct {
		In string
	}

	testcases := []Testcase{
		{
			In: `
			return 5;
			return 10;
			return 993322;
			`,
		},
	}

	for _, tc := range testcases {
		l := lexer.New(tc.In)
		p := parser.New(l)
		program := p.ParseProgram()

		t.Run(tc.In, func(t *testing.T) {
			for _, stmt := range program.Statements {
				returnStmt, ok := stmt.(*ast.ReturnStatement)
				assert.True(t, ok)
				assert.Equal(t, returnStmt.TokenLiteral(), "return")
			}
		})
	}
}
