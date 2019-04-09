package lexer

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"bitbucket.org/mikelsr/gauzaez/lexer/automaton"
)

// Lexer is used to build and store the token table
type Lexer struct {
	ahead       int
	buffer      bytes.Buffer // buffer stores characters of the current token
	currentChar byte
	position    position
	source      []byte // source contains the source file content
	tokenizer   Tokenizer
}

type position struct {
	column int
	index  int
	line   int
}

// MakeLexer is the default constructor for lexer
func MakeLexer(rules Rules) (*Lexer, error) {
	l := new(Lexer)
	l.tokenizer = Tokenizer{Nodes: make(map[string]*automaton.Node)}
	err := l.tokenizer.LoadRules(rules)
	if err != nil {
		return nil, err
	}
	l.Reset()
	return l, nil
}

func (l *Lexer) consume(n int) {
	l.position.index += n
	l.position.column += l.ahead
	l.ahead = 0
	l.loadChar()
}

func (l *Lexer) loadChar() {
	l.currentChar = l.source[l.position.index+l.ahead]
}

func (l *Lexer) peek(n int) {
	l.ahead++
	l.loadChar()
}

// processToken manages lexer position in case of EOS or ↵
func (l *Lexer) processToken(t automaton.Token, v string) {
	if v == "\n" {
		l.position.column = 0
		l.position.line++
	}
}

// Reset readys the lexer for tokenizing
func (l *Lexer) Reset() {
	l.ahead = 0
	l.buffer = *bytes.NewBuffer(nil)
	l.position.column = 1
	l.position.index = 0
	l.position.line = 1
}

// Tokenize tokenizes all the input in the io.Reader
func (l *Lexer) Tokenize(in io.Reader) (*TokenTable, error) {
	source, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}

	tokenTable := &TokenTable{}
	l.source = source

	l.loadChar()
	// q0 is the initial token
	l.tokenize(l.tokenizer.Nodes["q0"], tokenTable)

	// tokenize returns end if analyzed element was last element
	for l.position.index < len(l.source) {
		end, err := l.tokenize(l.tokenizer.Nodes["q0"], tokenTable)
		if err != nil {
			return nil, err
		}
		if end {
			break
		}
	}
	return tokenTable, nil
}

// tokenize a single token
func (l *Lexer) tokenize(node *automaton.Node, tt *TokenTable) (bool, error) {
	// if l.currentChar matches a pattern, follow corresponding path
	for _, path := range node.Paths {
		if path.Exp.MatchString(string(l.currentChar)) {
			log.Printf("Char %s sent to: %v",
				escape(string(l.currentChar)), path)
			l.buffer.WriteByte(l.currentChar)

			// if it's the last character
			if l.position.index+l.ahead+1 >= len(l.source) {
				if path.Target.Final {
					l.writeToken(path.Target.Token, tt)
					return true, nil
				}
			}
			// if it's not the last character
			l.peek(1)
			return l.tokenize(path.Target, tt)
		}
	}
	// if current state is a final state, return token & consume
	if node.Final {
		l.writeToken(node.Token, tt)
		l.consume(l.ahead)
		return false, nil
	}
	// if state is not final and there are no valid paths, unrecognized token
	return true, fmt.Errorf("unrecognized token '%v' (line %d, column %d)",
		l.buffer.String(), l.position.line, l.position.column)
}

// writeToken writes the token t to the tokentable tt using the positions at
// the lexer
func (l *Lexer) writeToken(t automaton.Token, tt *TokenTable) {
	value := l.buffer.String()

	l.buffer.Reset()
	// if character shouldn't be ignored, write it on token table

	tt.writeToken(t, value, uint(l.position.line),
		uint(l.position.column), uint(l.position.column+l.ahead))

	l.processToken(t, value)
}
