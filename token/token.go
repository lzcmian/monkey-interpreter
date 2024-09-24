package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	BANG     = "!"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

// LookUpType 给定一个字母开头的字面量，返回其对应的 token 类型（IDENT | 语言中的关键词））
func LookUpType(literal string) TokenType {
	if literal == "fn" {
		return FUNCTION
	}
	if literal == "let" {
		return LET
	}
	if literal == "true" {
		return TRUE
	}
	if literal == "false" {
		return FALSE
	}
	if literal == "if" {
		return IF
	}
	if literal == "else" {
		return ELSE
	}
	if literal == "return" {
		return RETURN
	}
	return IDENT
}
