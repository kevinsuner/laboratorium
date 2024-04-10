package main

import "fmt"

type Values map[string]any

func (v Values) GetString(key string) string {
    val, ok := v[key].(string)
    if !ok {
        return ""
    }

    return val
}

func (v Values) GetInt64(key string) int64 {
    val, ok := v[key].(int64)
    if !ok {
        return 0
    }

    return val
}

func (v Values) GetBoolean(key string) bool {
    val, ok := v[key].(bool)
    if !ok {
        return false
    }

    return val
}

func main() {
    input := `title = "foobar";
port = 8080;
debug = true;`

    lexer := NewLexer(input)
    parser := NewParser(lexer)
    config := parser.ParseConfig()
       
    values := make(Values, 0)
    for _, declaration := range config.declarations {
       if _, ok := values[declaration.Ident()]; !ok {
            values[declaration.Ident()] = declaration.Value()
       }
    }

    fmt.Printf("title: \t%s\n", values.GetString("title"))
    fmt.Printf("port: \t%d\n", values.GetInt64("port"))
    fmt.Printf("debug: \t%t\n", values.GetBoolean("debug"))
}
