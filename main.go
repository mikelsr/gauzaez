package main

import (
	"flag"
	"log"
	"os"

	"bitbucket.org/mikelsr/gauzaez/lexer"
)

func main() {

	sourcePath := flag.String("source", "", "Source file to process")
	rulesPath := flag.String("rules", defaultRules,
		"[optional] JSON representing automaton that\n\tdefines the behaviour of the lexer")
	flag.Parse()

	if *sourcePath == "" || *rulesPath == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// rules
	rules, err := lexer.MakeRules(*rulesPath)
	if err != nil {
		panic(err)
	}

	// create lexer and apply rules to lexer
	lex, err := lexer.MakeLexer(*sourcePath, rules)
	if err != nil {
		panic(err)
	}

	// tokenize source
	err = lex.Tokenize()
	if err != nil {
		panic(err)
	}
	log.Printf("\n%s\n", lex.TokenTable)
}
