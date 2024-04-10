package main

import "testing"

func Test_Value(t *testing.T) {
    ast := &Config{
        declarations: []Declaration{
            &Statement{
                Token: Token{Type: IDENT, Literal: "title"},
                Name: &Identifier{
                    Token: Token{Type: IDENT, Literal: "title"},
                    Val: "title",
                },
                Val: &Identifier{
                    Token: Token{Type: IDENT, Literal: "foobar"},
                    Val: "foobar",
                },
            },
        },
    }

    if ast.Value() != "title = foobar" {
        t.Errorf("ast.Value() wrong. got=%q", ast.Value())
    }
}
