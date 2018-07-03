package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"bitbucket.org/mikelsr/gauzaez/lexer"
)

func main() {

	defaultRules := fmt.Sprintf("%s/src/github.com/mikelsr/gauzaez/conf/lexer_rules.json",
		os.Getenv("GOPATH"))

	source := flag.String("source", "", "Source file to process")
	lexerRules := flag.String("rules", defaultRules,
		"[optional] JSON representing automaton that\n\tdefines the behaviour of the lexer")

	flag.Parse()
	log.Printf("Source: %s\n", *source)
	log.Printf("Rules: %s\n", *lexerRules)

	if *source == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	rules, err := lexer.MakeRules(*lexerRules)
	if err != nil {
		panic(err)
	}
	lex, err := lexer.MakeLexer(*source, rules)
	if err != nil {
		panic(err)
	}
	err = lex.Tokenize()
	if err != nil {
		panic(err)
	}
	log.Printf("\n%s\n", lex.TokenTable)
}
