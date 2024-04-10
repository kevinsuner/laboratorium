package main

type Node interface {
    Ident()      string
    Type()      string
    Value()     any
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

type Statement struct {
    Token   Token
    Name    *Identifier
    Val     Expression
}

func (s *Statement) declarationNode() {}
func (s *Statement) Ident() string { return s.Token.Literal }
func (s *Statement) Type() string { return s.Val.Type() }
func (s *Statement) Value() any { return s.Val.Value() }

type Identifier struct {
    Token   Token
    Val     string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) Ident() string { return "" } // noop
func (i *Identifier) Type() string { return string(i.Token.Type) }
func (i *Identifier) Value() any { return i.Val }

type Integer struct {
    Token Token
    Val int64
}

func (i *Integer) expressionNode() {}
func (i *Integer) Ident() string { return "" } // noop
func (i *Integer) Type() string { return string(i.Token.Type) }
func (i *Integer) Value() any { return i.Val }

type String struct {
    Token Token
    Val string
}

func (s *String) expressionNode() {}
func (s *String) Ident() string { return "" } // noop
func (s *String) Type() string { return string(s.Token.Type) }
func (s *String) Value() any { return s.Val }

type Boolean struct {
    Token Token
    Val bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) Ident() string { return "" } // noop
func (b *Boolean) Type() string { return string(b.Token.Type) }
func (b *Boolean) Value() any { return b.Val }
