package lexer

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var tokenizer *Tokenizer
var CONFPATH = fmt.Sprintf("%s/src/bitbucket.org/mikelsr/gauzaez/test/lexer", os.Getenv("GOPATH"))
var inconsistentRules = fmt.Sprintf("%s/inconsistent_rules.json", CONFPATH)
var incompleteRules = fmt.Sprintf("%s/incomplete_rules.json", CONFPATH)
var incorrectJSON = fmt.Sprintf("%s/incorrect_json.json", CONFPATH)

func TestLoadRules(t *testing.T) {
	tokenizer = NewTokenizer()

	// this will return an error
	invalidRules, _ := MakeRules(inconsistentRules)
	tokenizer.LoadRules(*invalidRules)
	// this won't
	tokenizer.LoadRules(*rules)
}

func TestMakeRules(t *testing.T) {
	// Both should fail
	_, _ = MakeRules("I-really-hope-this-is-not-a-file")
	_, _ = MakeRules(incorrectJSON)
}

// TestTokenize test correct lexer and
// builds a lexer with an incomplete automaton
func TestTokenize(t *testing.T) {

	inclompleteRules, _ := MakeRules(incompleteRules)
	// incomplete lexer
	failLexer, _ := MakeLexer(testSourceFile, *inclompleteRules)
	failLexer.Tokenize()
	// correct lexer
	lexer.Tokenize()
	log.Printf("\n%s\n", lexer.TokenTable.String())
}
