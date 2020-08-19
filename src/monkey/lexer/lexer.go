package lexer

import "monkey/token"

type Lexer struct {
    input        string
    position     int  // current position in input (points to current char)
    readPosition int  // current reading position in input (after current char)
    ch           byte // current char under examination
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar();
    return l
}

// TODO change this to support unicode and UTF-8 characters
func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        // 0 is the ACII code for 'NUL' which signifies haven't read anything or 'NUL' character
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }

    l.position = l.readPosition
    l.readPosition += 1
}

func (l* Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        return 0
    } else {
        return l.input[l.readPosition]
    }
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()

    switch l.ch {
        case '=':
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                // need to construct '=='
                literal := string(ch) + string(l.ch)
                // can't reuse newToken here since the literal is string and not ch
                // still a golang n00b but, maybe we can overload inputs? or is it bad style/not supported
                tok = token.Token{Type: token.EQ, Literal:  literal}
            } else {
                tok = newToken(token.ASSIGN, l.ch)
            }
        case '+':
            tok = newToken(token.PLUS, l.ch)
        case '-':
            tok = newToken(token.MINUS, l.ch)
        case '!':
            // chance to refactor?
            if l.peekChar() == '=' {
                ch := l.ch
                l.readChar()
                literal := string(ch) + string(l.ch)
                tok = token.Token{Type: token.NOT_EQ, Literal: literal}
            } else {
                tok = newToken(token.BANG, l.ch)
            }
        case '/':
            tok = newToken(token.SLASH, l.ch)
        case '*':
            tok = newToken(token.ASTERISK, l.ch)
        case '<':
            tok = newToken(token.LT, l.ch)
        case '>':
            tok = newToken(token.GT, l.ch)
        case ';':
            tok = newToken(token.SEMICOLON, l.ch)
        case '(':
            tok = newToken(token.LPAREN, l.ch)
        case ')':
            tok = newToken(token.RPAREN, l.ch)
        case ',':
            tok = newToken(token.COMMA, l.ch)
        case '{':
            tok = newToken(token.LBRACE, l.ch)
        case '}':
            tok = newToken(token.RBRACE, l.ch)
        case 0:
            tok.Literal = ""
            tok.Type = token.EOF
        default:
            if isLetter(l.ch) {
                tok.Literal = l.readIdentifier()
                tok.Type = token.LookupIdent(tok.Literal)
                // short circuit here since we have read enough chars
                return tok
            } else if isDigit(l.ch) {
                tok.Type = token.INT
                tok.Literal = l.readNumber()
                return tok
            }  else {
                tok = newToken(token.ILLEGAL, l.ch)
            }
    }
  l.readChar()
  return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}


// readNumber and readIndentifier can be refactored into a single function
func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

// extracts identifier from input string
func (l *Lexer) readIdentifier() string {
    position := l.position

    for isLetter(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

// different languages have different rules for handling this altogether
func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

func isLetter(ch byte) bool {
    // TODO I wonder if a regex would be simpler?
    return 'a' <= ch &&  ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// TODO this only supports integers, maybe expand to other formats?
func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}
