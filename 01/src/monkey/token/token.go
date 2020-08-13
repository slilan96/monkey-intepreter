package token

type TokenType string // this allows us to define the tokens as constants or..enum

// define the type of the Token
type Token struct {
  Type TokenType
  Literal string
}

const (
  ILLEGAL = "ILLEGAL" // inform the parser that character is not identified
  EOF     = "EOF"

  // Identifiers (or variables) + literals
  IDENT = "IDENT"
  INT   = "INT"

  // OPERATORS
  ASSIGN   = "="
  PLUS     = "+"
  MINUS    = "-"
  BANG     = "!"
  ASTERISK = "*"
  SLASH    = "/"
  LT       = "<"
  GT       = ">"

  // DELIMITERS 
  COMMA     = ","
  SEMICOLON = ";"

  RPAREN    = ")"
  LPAREN    = "("
  RBRACE    = "}"
  LBRACE    = "{"

  // KEYWORDS
  FUNCTION = "FUNCTION"
  LET      = "LET"
)

var keywords = map[string]TokenType {
    "fn" : FUNCTION,
    "let": LET,
}


func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
      return tok;
    }

    return IDENT
}
