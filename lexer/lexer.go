package lexer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

/* -- Structs -- */

// Lexer is used to build and store the token table
type Lexer struct {
	ahead       int
	buffer      bytes.Buffer
	chars       []byte
	currentChar byte
	position    position
	filename    string
	TokenTable  *TokenTable
	tokenizer   *Tokenizer
}

type position struct {
	column int
	index  int
	line   int
}

/* -- Constructors -- */

// MakeLexer is the default constructor for lexer
func MakeLexer(filename string, rules *Rules) (*Lexer, error) {
	l := new(Lexer)
	l.ahead = 0
	l.buffer = *bytes.NewBuffer(nil)
	l.position.column = 1
	l.position.index = 0
	l.position.line = 1
	l.filename = filename
	l.TokenTable = NewTokenTable()

	l.loadFile()
	l.tokenizer = NewTokenizer()
	err := l.tokenizer.LoadRules(rules)
	if err != nil {
		return nil, err
	}
	return l, nil
}

/* -- Methods -- */

func (l *Lexer) consume(n int) {
	l.position.index += n
	l.position.column += l.ahead
	l.ahead = 0
	l.loadChar()
}

func (l *Lexer) loadChar() {
	l.currentChar = l.chars[l.position.index+l.ahead]
}

func (l *Lexer) loadFile() {
	chars, err := ioutil.ReadFile(l.filename)
	if err != nil {
		panic(err)
	}
	l.chars = chars
	log.Printf("Loaded contents of '%s' (%d bytes in total).\n",
		l.filename, len(l.chars))
}

func (l *Lexer) peek(n int) {
	l.ahead++
	l.loadChar()
}

// processToken manages lexer position in case of EOS or ↵
func (l *Lexer) processToken(t Token, v string) {
	if v == "\n" {
		l.position.column = 0
		l.position.line++
	}
}

// Tokenize iterates l.chars and fills the token table l.TokenTable
func (l *Lexer) Tokenize() error {
	l.loadChar()
	l.tokenize(l.tokenizer.Nodes["q0" /* initial token */])

	// tokenize returns end if analyzed element was last element
	for l.position.index < len(l.chars) {
		end, err := l.tokenize(l.tokenizer.Nodes["q0"])
		if err != nil {
			return err
		}
		if end {
			break
		}
	}
	return nil
}

func (l *Lexer) tokenize(node *Node) (bool, error) {
	// if l.currentChar matches a pattern, follow corresponding path
	for _, path := range node.Paths {
		if path.Exp.MatchString(string(l.currentChar)) {
			log.Printf("Char %s sent to: %v",
				escape(string(l.currentChar)), path)
			l.buffer.WriteByte(l.currentChar)

			// if it's the last character
			if l.position.index+l.ahead+1 >= len(l.chars) {
				if path.Target.Final {
					l.writeToken(path.Target.Token)
					return true, nil
				}
			}
			// if it's not the last character
			l.peek(1)
			return l.tokenize(path.Target)
		}
	}
	// if current state is a final state, return token & consume
	if node.Final {
		l.writeToken(node.Token)
		l.consume(l.ahead)
		return false, nil
	}
	// if state is not final and there are no valid paths, unrecognized token
	return true, fmt.Errorf("unrecognized token '%v' (line %d, column %d)",
		l.buffer.String(), l.position.line, l.position.column)
}

func (l *Lexer) writeToken(t Token) {
	value := l.buffer.String()
	l.buffer.Reset()
	// if character shouldn't be ignored, write it on token table

	l.TokenTable.writeToken(t, value, uint(l.position.line),
		uint(l.position.column), uint(l.position.column+l.ahead))

	l.processToken(t, value)
}
