package lexer

import (
	"fmt"
	"os"
)

var (
	testSourceFile = "/tmp/gauzaez_lexer_test.txt"
	rulesFile      = fmt.Sprintf("%s/src/bitbucket.org/mikelsr/gauzaez/conf/"+
		"lexer_rules.json", os.Getenv("GOPATH"))
	testSourceFileContent []byte
	lexer                 *Lexer
	rules                 *Rules
)
