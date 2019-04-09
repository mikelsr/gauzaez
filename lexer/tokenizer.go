package lexer

import (
	"fmt"
	re "regexp"

	"bitbucket.org/mikelsr/gauzaez/lexer/automaton"
)

// Tokenizer is the main struct used to iterate the source and assign Tokens
type Tokenizer struct {
	Nodes map[string]*automaton.Node
}

// LoadRules transforms preNodes into Nodes given a set of rules
func (t *Tokenizer) LoadRules(rules Rules) error {
	// first iteration creates nodes and sets Final, Token attributes
	for id, n := range rules.Nodes {
		node := Node{Final: n.Final}
		if n.Final {
			if !rules.Tokens[n.Token] {
				return fmt.Errorf("Token '%s' is not listed in rules", n.Token)
			}
			node.Token = n.Token
		}
		t.Nodes[id] = &node
	}

	// seconds iteration creates and assings paths to connect nodes
	for id, n := range rules.Nodes {
		for exp, target := range n.Paths {
			path := Path{
				Exp:    *re.MustCompile(exp),
				Target: t.Nodes[target],
			}
			t.Nodes[id].Paths = append(t.Nodes[id].Paths, path)
		}
	}
	return nil
}
