package main

import "bytes"

type Node interface {
    TokenLiteral()  string
    Value()         string
}

type Declaration interface {
    Node
    declarationNode()
}

type Expression interface {
    Node
    expressionNode()
}

type Config struct {
    declarations []Declaration
}

func (c *Config) TokenLiteral() string {
    if len(c.declarations) > 0 {
        return c.declarations[0].TokenLiteral()
    } else {
        return ""
    }
}

func (c *Config) Value() string {
    var out bytes.Buffer
    for _, declaration := range c.declarations {
        out.WriteString(declaration.Value())
    }

    return out.String()
}

type Identifier struct {
    Token   Token
    Val     string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) Value() string { return i.Val }

type Statement struct {
    Token   Token
    Name    *Identifier
    Val     Expression
}

func (s *Statement) declarationNode() {}
func (s *Statement) TokenLiteral() string { return s.Token.Literal }
func (s *Statement) Value() string {
    var out bytes.Buffer
    out.WriteString(s.TokenLiteral())
    out.WriteString(" = ")
    
    if s.Val != nil {
        out.WriteString(s.Val.Value())
    }

    return out.String()
}

type Integer struct {
    Token Token
    Val int64
}

func (i *Integer) expressionNode() {}
func (i *Integer) TokenLiteral() string { return i.Token.Literal }
func (i *Integer) Value() string { return i.Token.Literal }

type String struct {
    Token Token
    Val string
}

func (s *String) expressionNode() {}
func (s *String) TokenLiteral() string { return s.Token.Literal }
func (s *String) Value() string { return s.Token.Literal }

type Boolean struct {
    Token Token
    Val bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) Value() string { return b.Token.Literal }
