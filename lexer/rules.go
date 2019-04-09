package lexer

import (
	"encoding/json"
	"io/ioutil"

	"bitbucket.org/mikelsr/gauzaez/lexer/automaton"
)

// Rules is used by the tokenizer to build the token table
type Rules struct {
	Nodes        map[string]automaton.PreNode `json:"nodes"`
	TokenStrings []automaton.Token            `json:"tokens"`
	Tokens       map[automaton.Token]bool
}

// MakeRules loads rules from a JSON file to a Rules struct
func MakeRules(filename string) (*Rules, error) {
	ruleFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	rules := new(Rules)
	err = json.Unmarshal(ruleFile, rules)
	if err != nil {
		return nil, err
	}

	rules.Tokens = make(map[automaton.Token]bool)
	for _, t := range rules.TokenStrings {
		rules.Tokens[t] = true
	}

	return rules, nil
}
