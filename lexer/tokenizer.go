package lexer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	re "regexp"
)

/* -- Structs -- */

// Node is a state of the Tokenizer automaton
type Node struct {
	Final bool // final state
	Paths []*Path
	Token Token
}

// preNode is used by the Tokenizer constructor to store unprocessed values
// that will form a complete Node
type preNode struct {
	Final bool              `json:"final"`
	Paths map[string]string `json:"paths"`
	Token Token             `json:"token,omitempty"`
}

// Path connects two nodes over a regular expression
type Path struct {
	Exp    re.Regexp
	Target *Node
}

// Rules is used by the tokenizer to build the token table
type Rules struct {
	Nodes        map[string]preNode `json:"nodes"`
	TokenStrings []Token            `json:"tokens"`
	Tokens       map[Token]bool
}

// Tokenizer is the main struct used to iterate the source and assign Tokens
type Tokenizer struct {
	Nodes map[string]*Node
}

/* -- Constructors -- */

// MakePath is the default constructor for Path
// t: Target Node
// exp: regular expression used as transition condition
func MakePath(target *Node, exp string) *Path {
	p := new(Path)
	p.Exp = *re.MustCompile(exp)
	p.Target = target
	return p
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

	rules.Tokens = make(map[Token]bool)
	for _, t := range rules.TokenStrings {
		rules.Tokens[t] = true
	}

	return rules, nil
}

// NewNode is the default constructor for Node knowing if it's a final state
func NewNode(final bool) *Node {
	n := new(Node)
	n.Final = final
	n.Paths = *new([]*Path)
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
func (t *Tokenizer) LoadRules(rules *Rules) error {
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
			path := MakePath(t.Nodes[target], exp)
			t.Nodes[id].Paths = append(t.Nodes[id].Paths, path)
		}
	}
	return nil
}
