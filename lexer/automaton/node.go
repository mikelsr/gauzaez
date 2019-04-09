package automaton

// Node is a state of the Tokenizer automaton
type Node struct {
	Final bool // final state
	Paths []Path
	Token Token
}

// PreNode is used by the Tokenizer constructor to store unprocessed values
// that will form a complete Node
type PreNode struct {
	Final bool              `json:"final"`
	Paths map[string]string `json:"paths"`
	Token Token             `json:"token,omitempty"`
}
