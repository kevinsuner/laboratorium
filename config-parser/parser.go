package main

import (
	"fmt"
	"strconv"
)

type Parser struct {
    lexer *Lexer
    errs  []error

    currentToken    Token
    peekToken       Token

    prefixParseFns  map[TokenType]prefixParseFn
}

type prefixParseFn func() Expression

func NewParser(lexer *Lexer) *Parser {
    p := &Parser{lexer: lexer, errs: make([]error, 0)}
    
    p.prefixParseFns = make(map[TokenType]prefixParseFn)
    p.registerPrefix(INT, p.parseInteger)
    p.registerPrefix(STRING, p.parseString)
    p.registerPrefix(TRUE, p.parseBoolean)
    p.registerPrefix(FALSE, p.parseBoolean)
    
    p.nextToken()
    p.nextToken()
    return p
}

func (p *Parser) registerPrefix(t TokenType, fn prefixParseFn) {
    p.prefixParseFns[t] = fn
}
 
func (p *Parser) errors() []error {
    return p.errs
}

func (p *Parser) nextToken() {
    p.currentToken = p.peekToken
    p.peekToken = p.lexer.nextToken()
}

func (p *Parser) ParseConfig() *Config {
    config := &Config{}
    config.declarations = []Declaration{}

    for p.currentToken.Type != EOF {
        declaration := p.parseDeclaration()
        if declaration != nil {
            config.declarations = append(config.declarations, declaration)
        }
        p.nextToken()
    }

    return config
}

func (p *Parser) parseDeclaration() Declaration {
    switch p.currentToken.Type {
    case IDENT:
        return p.parseStatement()
    default:
        return nil
    }
}

func (p *Parser) parseStatement() *Statement {
    statement := &Statement{Token: p.currentToken}
    statement.Name = &Identifier{Token: p.currentToken, Val: p.currentToken.Literal}

    if !p.expectPeek(ASSIGN) {
        return nil
    }

    p.nextToken()
    statement.Val = p.parseExpression()

    for !p.currentTokenIs(SEMICOLON) {
        p.nextToken()
    }

    return statement
}

func (p *Parser) parseExpression() Expression {
    prefix := p.prefixParseFns[p.currentToken.Type]
    if prefix == nil {
        return nil
    }

    return prefix()
}

func (p *Parser) parseInteger() Expression {
    integer := &Integer{Token: p.currentToken}
    value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
    if err != nil {
        p.errs = append(p.errs, err)
        return nil
    }

    integer.Val = value
    return integer
}

func (p *Parser) parseString() Expression {
    return &String{Token: p.currentToken, Val: p.currentToken.Literal}
}

func (p *Parser) parseBoolean() Expression {
    return &Boolean{Token: p.currentToken, Val: p.currentTokenIs(TRUE)}
}

func (p *Parser) expectPeek(t TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    } else {
        p.peekError(t)
        return false
    }
}

func (p *Parser) currentTokenIs(t TokenType) bool {
    return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
    return p.peekToken.Type == t
}

func (p *Parser) peekError(t TokenType) {
    p.errs = append(p.errs, fmt.Errorf(
        "expected next token to be %s, got %s instead",
        t, p.peekToken.Type,
    ))
}
