package lexer

import (
	"fmt"
	re "regexp"
)

/* -- Structs -- */

// Node is a state of the Tokenizer automaton
type Node struct {
	Final bool // final state
	Paths []Path
	Token Token
}

// preNode is used by the Tokenizer constructor to store unprocessed values
// that will form a complete Node
type preNode struct {
	Final bool              `json:"final"`
	Paths map[string]string `json:"paths"`
	Token Token             `json:"token,omitempty"`
}

// Tokenizer is the main struct used to iterate the source and assign Tokens
type Tokenizer struct {
	Nodes map[string]*Node
}

/* -- Constructors -- */

// NewNode is the default constructor for Node knowing if it's a final state
func NewNode(final bool) *Node {
	n := new(Node)
	n.Final = final
	n.Paths = make([]Path, 0)
	return n
}

// NewTokenizer is the default constructor for Tokernizer
func NewTokenizer() *Tokenizer {
	t := new(Tokenizer)
	t.Nodes = make(map[string]*Node)
	return t
}

/* -- Methods -- */

// LoadRules transforms preNodes into Nodes given a set of rules
func (t *Tokenizer) LoadRules(rules Rules) error {
	// first iteration creates nodes and sets Final, Token attributes
	for id, n := range rules.Nodes {
		node := NewNode(n.Final)
		if n.Final {
			if !rules.Tokens[n.Token] {
				return fmt.Errorf("Token '%s' is not listed in rules", n.Token)
			}
			node.Token = n.Token
		}
		t.Nodes[id] = node
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
