package lexer

import (
	"monkey-interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // 最后一次读取的位置
	readPosition int  // 当前该读取的位置
	ch           byte // 当前正在检查的字符 == input[position]
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
	case '=':
		if l.peekChar() == '=' { // ==
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = newToken(token.ASSIGN, '=')
		}
	case '+':
		tok = newToken(token.PLUS, '+')
	case '-':
		tok = newToken(token.MINUS, '-')
	case '!':
		if l.peekChar() == '=' { // !=
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = newToken(token.BANG, '!')
		}
	case '/':
		tok = newToken(token.SLASH, '/')
	case '*':
		tok = newToken(token.ASTERISK, '*')
	case '<':
		tok = newToken(token.LT, '<')
	case '>':
		tok = newToken(token.GT, '>')
	case ',':
		tok = newToken(token.COMMA, ',')
	case ';':
		tok = newToken(token.SEMICOLON, ';')
	case '(':
		tok = newToken(token.LPAREN, '(')
	case ')':
		tok = newToken(token.RPAREN, ')')
	case '{':
		tok = newToken(token.LBRACE, '{')
	case '}':
		tok = newToken(token.RBRACE, '}')
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	// Identifiers + literals + Keywords
	default:
		if isLetter(l.ch) { // 以字母开头，只有可能是关键词或标识符（变量名）
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpType(tok.Literal)
			return tok // 提前return，因为 readIdentifier() 已经将 readPosition 移到下一个字符了
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok // 提前return，因为 readIdentifier() 已经将 readPosition 移到下一个字符了
		} else if l.ch == 0 { // 读完了
			tok.Literal = ""
			tok.Type = token.EOF
		} else { // 未知情况 => 非法
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}
	l.readChar()
	return tok

}

// isLetter 判断 ch 是否是字母
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

// isDigit 判断 ch 是否是数字
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// newToken 返回一个单字符字面量的 token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

// readChar 将 lexer 的 ch 变为下一个字符，并更新 position、readPosition字段。若读完了，则将 ch 变为0
// 为什么将 ch = 0 能够安全地表示读取完毕? 因为我们是逐字符读取, 用户不可能输入一个byte值 = 0的字符, 因为这需要通过多个字符来转义.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // 已经读完了
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// readIdentifier 读取标识符（变量名、关键词）, 并将 read
func (l *Lexer) readIdentifier() string {
	cur := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[cur:l.position]
}

func (l *Lexer) readNumber() string {
	cur := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[cur:l.position]
}

// skipWhiteSpace 跳过所有空白符,
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readString() string {
	position := l.position + 1 // 字符串内容的第一个字符的位置
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}
