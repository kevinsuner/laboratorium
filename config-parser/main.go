package main

import "fmt"

// TODO's
// Use Value() to return the value of the expression
//  - Don't know if it should return string or any
// Add TokenType() to return the token type of the declaration/expression

func main() {
    input := `title = "foobar";
port = 8080;
debug = true;`

    lexer := NewLexer(input)
    parser := NewParser(lexer)

    config := parser.ParseConfig()
    for _, declaration := range config.declarations {
        fmt.Printf("ident: \t%s\n", declaration.Ident())
        fmt.Printf("type: \t%s\n", declaration.Type())
        fmt.Printf("value: \t%v\n", declaration.Value())
    }
}
