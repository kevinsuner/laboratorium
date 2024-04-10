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
                Val: &String{
                    Token: Token{Type: STRING, Literal: "foobar"},
                    Val: "foobar",
                },
            },
        },
    }
    
    if ast.declarations[0].Ident() != "title" {
        t.Errorf("ast.Ident() wrong. got=%s", ast.declarations[0].Ident())
    }

    if ast.declarations[0].Type() != STRING {
        t.Errorf("ast.Type() wrong. got=%s", ast.declarations[0].Type())
    }

    if ast.declarations[0].Value() != "foobar" {
        t.Errorf("ast.Value() wrong. got=%s", ast.declarations[0].Value())
    }
}
