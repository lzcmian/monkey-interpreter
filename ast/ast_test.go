package ast

import (
	"monkey-interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := Program{
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

	got := program.String()
	want := "let myVar = anotherVar;"
	if got != want {
		t.Errorf("program.String() wrong. got=%q, want=%q", got, want)
	}
}
