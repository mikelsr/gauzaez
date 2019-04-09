package lexer

import (
	"testing"

	"bitbucket.org/mikelsr/gauzaez/lexer/automaton"
)

func TestTokenizer_LoadRules(t *testing.T) {
	tokenizer := Tokenizer{Nodes: make(map[string]*automaton.Node)}
	// this will return an error
	invalidRules, _ := MakeRules(inconsistentRules)
	tokenizer.LoadRules(*invalidRules)
	// this won't
	tokenizer.LoadRules(*rules)
}

func TestMakeRules(t *testing.T) {
	// Both should fail
	if _, err := MakeRules("I-really-hope-this-is-not-a-file"); err == nil {
		t.Fatalf("successfully loaded invalid rules")
	}
	if _, err := MakeRules(incorrectJSON); err == nil {
		t.Fatalf("successfully loaded invalid rules")
	}
}
