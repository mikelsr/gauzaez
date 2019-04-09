package automaton

import re "regexp"

// Path connects two nodes over a regular expression
type Path struct {
	Exp    re.Regexp // regular expression used as transition condition
	Target *Node     // target node
}
