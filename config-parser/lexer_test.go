package main

import "testing"

func Test_nextToken(t *testing.T) {
    input := `title = "foobar";
port = 8080;
debug = true;`

    tests := []struct{
        expectedType    TokenType
        expectedLiteral string
    }{
        {IDENT, "title"},
        {ASSIGN, "="},
        {STRING, "foobar"},
        {SEMICOLON, ";"},
        {IDENT, "port"},
        {ASSIGN, "="},
        {INT, "8080"},
        {SEMICOLON, ";"},
        {IDENT, "debug"},
        {ASSIGN, "="},
        {TRUE, "true"},
        {SEMICOLON, ";"},
        {EOF, ""},
    }

    l := NewLexer(input)
    for i, tc := range tests {
        tok := l.nextToken()
        if tok.Type != tc.expectedType {
            t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
                i, tc.expectedType, tok.Type)
        }

        if tok.Literal != tc.expectedLiteral {
            t.Fatalf("tests[%d] - token literal wrong. expected=%q, got=%q",
                i, tc.expectedLiteral, tok.Literal)
        }
    }
}
