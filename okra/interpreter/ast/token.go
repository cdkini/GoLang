package ast

// A Token is a substring of source text given context or meaning by the scanner
type Token struct {
	Type    TokenType
	Lexeme  string // Exact substring lexed by scanner
	Literal interface{}
	Line    int
	Col     int
}

// Only used for testing purposes within interpret_test package
func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int, col int) *Token {
	return &Token{tokenType, lexeme, literal, line, col}
}

type TokenType int

const (
	And TokenType = iota
	Bang
	BangEqual
	Class
	Comma
	Construct
	Dot
	Else
	EOF
	Equal
	EqualEqual
	False
	For
	Func
	Greater
	GreaterEqual
	Identifier
	If
	Invalid
	LeftBrace
	LeftBracket
	LeftParen
	Less
	LessEqual
	Minus
	Null
	Numeric
	Or
	Plus
	Print
	Return
	RightBrace
	RightBracket
	RightParen
	Semicolon
	Slash
	Star
	String
	Struct
	Super
	This
	True
	Variable
)

var KeywordDict = map[string]TokenType{
	"class":  Class,
	"else":   Else,
	"false":  False,
	"for":    For,
	"func":   Func,
	"if":     If,
	"null":   Null,
	"print":  Print,
	"return": Return,
	"super":  Super,
	"this":   This,
	"true":   True,
	"var":    Variable,
}