package lexer

import "fmt"

// Token is used to identify token types
type Token string

const (
	Access         Token = "access"
	Assignator     Token = "assignator"
	BinaryOperator Token = "binary-operator"
	Block          Token = "block"
	Brace          Token = "brace"
	Comparator     Token = "comparator"
	EOS            Token = "end-of-statement"
	Identifier     Token = "identifier"
	Index          Token = "index"
	Hexadecimal    Token = "hexadecimal"
	Negation       Token = "negation"
	Number         Token = "number"
	Operator       Token = "operator"
	Separator      Token = "separator"
	String         Token = "string"
	Whitespace     Token = "whitespace"
	Unknown        Token = "unknown"
)

// Tokens is an array containing all Tokens
var Tokens = []Token{
	Access, Assignator, BinaryOperator, Block, Brace, Comparator,
	EOS, Identifier, Index, Hexadecimal, Negation, Number,
	Operator, Separator, String, Whitespace,
}

// GetToken returns a token given a string
// TODO: change iterative search for more efficient search
func GetToken(token string) (Token, error) {
	for _, t := range Tokens {
		if string(t) == token {
			return t, nil
		}
	}
	return Unknown, fmt.Errorf("incorrect token type: %s", token)
}
