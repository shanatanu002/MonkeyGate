package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals
	IDENT = "IDENT" //add, foobar, x, y, ...
	INT   = "INT"   //123456

	//Operators
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	PLUS      = "+"
	ASSIGN    = "="

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
