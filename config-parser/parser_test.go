package main

import "testing"

func Test_declarations(t *testing.T) {
    input := `title = "foobar";
port = 8080;
debug = true;`

    l := NewLexer(input)
    p := NewParser(l)

    config := p.ParseConfig()
    checkParserErrors(t, p)

    if len(config.declarations) != 3 {
        t.Fatalf("config.declarations does not contain 3 declarations, got=%d",
            len(config.declarations))
    }

    tests := []struct{
        expectedIdentifier string
    }{
        {"title"},
        {"port"},
        {"debug"},
    }

    for i, tc := range tests {
        declaration := config.declarations[i]
        if !testDeclaration(t, declaration, tc.expectedIdentifier) {
            return
        }
    }
}

func checkParserErrors(t *testing.T, p *Parser) {
    errs := p.errors()
    if len(errs) == 0 {
        return
    }

    t.Errorf("parser has %d errors", len(errs))
    for _, err := range errs {
        t.Errorf("parser error: %v", err)
    }

    t.FailNow()
}

func testDeclaration(t *testing.T, d Declaration, name string) bool {
    if len(d.Ident()) == 0 {
        t.Error("d.Ident of zero length")
        return false
    }

    statement, ok := d.(*Statement)
    if !ok {
        t.Errorf("d not *Statement. got=%T", d)
        return false
    }

    if statement.Name.Val != name {
        t.Errorf("statement.Name.Val not '%s'. got=%s", name, statement.Name.Val)
        return false
    }

    return true
}
