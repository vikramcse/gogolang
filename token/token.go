package token

// TokenType specifies the type of token which is string
type TokenType string

// Token is data structure for holding a token type and its value
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL Special Type
	ILLEGAL = "ILLEGAL"
	// EOF End Of File Special Type
	EOF = "EOF"

	// IDENT Identifiers
	IDENT = "IDENT" // add, delete, x, y, ...

	// INT Literal
	INT = "INT" // 123456789

	// ASSIGN operators
	ASSIGN = "="

	// PLUS operators
	PLUS = "+"

	// MINUS Operator
	MINUS = "-"

	// BANG Operator
	BANG = "!"

	// ASTERISK Operator
	ASTERISK = "*"

	// SLASH Operator
	SLASH = "/"

	// LT Operator
	LT = "<"

	// GT Operator
	GT = ">"

	// COMMA Delimiters
	COMMA = ","

	// SEMICOLON Delimiters
	SEMICOLON = ";"

	// LPAREN Delimiters
	LPAREN = "("
	// RPAREN Delimiters
	RPAREN = "]"
	// LBRACE Delimiters
	LBRACE = "{"
	// RBRACE Delimiters
	RBRACE = "}"

	// FUNCTION Keywords
	FUNCTION = "FUNCTION"

	// LET Keywords
	LET = "LET"
	// TRUE Keywords
	TRUE = "TRUE"
	// FALSE Keywords
	FALSE = "FALSE"
	// IF Keywords
	IF = "IF"
	// ELSE Keywords
	ELSE = "ELSE"
	// RETURN Keywords
	RETURN = "RETURN"

	// EQ Operator
	EQ = "=="

	// NOTEQ Operator
	NOTEQ = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent will return identifier for given keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
