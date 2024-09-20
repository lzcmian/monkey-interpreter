package lexer

import "monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpace() // 跳过所有空白符
	switch l.ch {
	// =+,;(){}
	case '=':
		tok = newToken(token.ASSIGN, "=")
	case '+':
		tok = newToken(token.PLUS, "+")
	case ',':
		tok = newToken()
	case ';':
		tok = newToken()
	case '(':
		tok = newToken()
	case ')':
		tok = newToken()
	case '{':
		tok = newToken()
	case '}':
		tok = newToken()

	// Identifiers + literals + Keywords
	default:
		if isLetter(l.ch){  // 以字母开头，只有可能是关键词或标识符（变量名）
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpType(tok.Literal)
			// ret ?
		} else if isDigit(l.ch){
			tok = newToken(token.INT, l.readNumber())
			// ret ?
		}else{  // 未知情况 => 非法
			tok = newToken(token.ILLEGAL, l.ch)

		}


	}

}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// readChar 将 lexer 的 ch 变为下一个字符，并更新 position、readPosition字段。若读完了，则将 ch 变为0
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // 已经读完了
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}
