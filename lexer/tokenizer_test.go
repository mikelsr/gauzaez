package lexer

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var tokenizer *Tokenizer
var CONFPATH = fmt.Sprintf("%s/src/github.com/mikelsr/gauzaez/conf/test", os.Getenv("GOPATH"))
var failrulesFile = fmt.Sprintf("%s/lexer_malformed_rules.json", CONFPATH)
var incompleterulesFile = fmt.Sprintf("%s/lexer_incomplete_rules.json", CONFPATH)
var badJSON = fmt.Sprintf("%s/badformatted.json", CONFPATH)

func TestLoadRules(t *testing.T) {
	tokenizer = NewTokenizer()

	// this will return an error
	invalidRules, _ := MakeRules(failrulesFile)
	tokenizer.LoadRules(invalidRules)
	// this won't
	tokenizer.LoadRules(rules)
}

func TestMakeRules(t *testing.T) {
	// Both should fail
	_, _ = MakeRules("I-really-hope-this-is-not-a-file")
	_, _ = MakeRules(badJSON)
}

// TestTokenize test correct lexer and
// builds a lexer with an incomplete automaton
func TestTokenize(t *testing.T) {

	inclompleteRules, _ := MakeRules(incompleterulesFile)
	// incomplete lexer
	failLexer, _ := MakeLexer(testSourceFile, inclompleteRules)
	failLexer.Tokenize()
	// correct lexer
	lexer.Tokenize()
	log.Printf("\n%s\n", lexer.TokenTable.String())
}
