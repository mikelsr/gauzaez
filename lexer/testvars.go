package lexer

import (
	"fmt"
	"os"
)

var (
	lexer                 *Lexer
	rules                 *Rules
	testSourceFileContent []byte
	tokenizer             *Tokenizer

	inconsistentRules = fmt.Sprintf("%s/inconsistent_rules.json", testConfPath)
	incompleteRules   = fmt.Sprintf("%s/incomplete_rules.json", testConfPath)
	incorrectJSON     = fmt.Sprintf("%s/incorrect_json.json", testConfPath)
	rulesFile         = fmt.Sprintf("%s/src/bitbucket.org/mikelsr/gauzaez/conf/"+
		"lexer_rules.json", os.Getenv("GOPATH"))
	testConfPath   = fmt.Sprintf("%s/src/bitbucket.org/mikelsr/gauzaez/test/lexer", os.Getenv("GOPATH"))
	testSourceFile = "/tmp/gauzaez_lexer_test.txt"
)
