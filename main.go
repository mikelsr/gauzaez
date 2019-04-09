package main

import (
	"flag"
	"fmt"
	"os"

	"bitbucket.org/mikelsr/gauzaez/lexer"
)

func main() {

	sourcePath := flag.String("source", "", "Source file to process")
	rulesPath := flag.String("rules", "",
		"[optional] JSON representing automaton that\n\tdefines the behaviour of the lexer")
	flag.Parse()

	if *sourcePath == "" || *rulesPath == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	// rules
	rules, err := lexer.MakeRules(*rulesPath)
	if err != nil {
		fmt.Printf("error parsing rules: '%s'\n", err)
		os.Exit(1)
	}

	// create lexer and apply rules to lexer
	lex, err := lexer.MakeLexer(*rules)
	if err != nil {
		panic(err)
	}

	// create reader for source file
	source, err := os.Open(*sourcePath)
	if err != nil {
		fmt.Printf("failed to open file '%s'\n", err)
		os.Exit(1)
	}

	// tokenize source
	table, err := lex.Tokenize(source)
	if err != nil {
		fmt.Printf("failed to lex file '%s': '%s'\n", *sourcePath, err)
		os.Exit(1)
	}
	fmt.Printf("\n%s\n", table)
}
