package main

type TokenType string

type Token struct {
    Type    TokenType
    Literal string
}

const (
    ILLEGAL     = "ILLEGAL"
    EOF         = "EOF"

    // Keywords
    TRUE        = "TRUE"
    FALSE       = "FALSE"

    // Identifiers
    IDENT       = "IDENT"

    // Literals
    INT         = "INT"
    STRING      = "STRING"

    // Operators
    ASSIGN      = "="

    // Delimiters
    SEMICOLON   = ";"
)

var keywords = map[string]TokenType{
    "true":     TRUE,
    "false":    FALSE,
}

type Lexer struct {
    input           string
    position        int     // current position in input (points to current char)
    readPosition    int     // current reading position in input (after current char)
    ch              byte    // current character under examination
}

func NewLexer(input string) *Lexer {
    l := &Lexer{input: input}
    l.readCharacter()
    return l
}

func (l *Lexer) readCharacter() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }

    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) nextToken() Token {
    var tok Token
    l.eatWhitespace()

    switch l.ch {
    case '=':
        tok = newToken(ASSIGN, l.ch)
    case ';':
        tok = newToken(SEMICOLON, l.ch)
    case '"':
        tok.Type = STRING
        tok.Literal = l.readString()
    case 0:
        tok.Type = EOF
        tok.Literal = ""
    default:
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = lookupIdentifier(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Type = INT
            tok.Literal = l.readNumber()
            return tok
        } else {
            tok = newToken(ILLEGAL, l.ch)
        }
    }

    l.readCharacter()
    return tok
}

func newToken(t TokenType, ch byte) Token {
    return Token{Type: t, Literal: string(ch)}
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func lookupIdentifier(identifier string) TokenType {
    if tok, ok := keywords[identifier]; ok {
        return tok
    }

    return IDENT
}

func (l *Lexer) eatWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readCharacter()
    }
}

func (l *Lexer) readString() string {
    position := l.position + 1
    for {
        l.readCharacter()
        if l.ch == '"' || l.ch == 0 {
            break
        }
    }

    return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readCharacter()
    }

    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readCharacter()
    }

    return l.input[position:l.position]
}
