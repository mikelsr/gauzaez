package lexer

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// TestMain runs preparations for the other tests
func TestMain(m *testing.M) {
	// create an invalid source code file
	testSourceFileContent = []byte("x = 1 +\t2.3\ny = 4 ** 5\n\ntext=\"hi\"")
	if err := ioutil.WriteFile(testSourceFile, testSourceFileContent, 0644); err != nil {
		log.Fatal(err)
	}

	// load rules
	rules, _ = MakeRules(rulesFile)

	// configure lexer
	lexer, _ = MakeLexer(*rules)
	exitCode := m.Run()
	os.Remove(testSourceFile)
	os.Exit(exitCode)
}

func TestLexer_Tokenize(t *testing.T) {

	// incomplete lexer
	source, err := os.Open(testSourceFile)
	inclompleteRules, _ := MakeRules(incompleteRules)
	failLexer, _ := MakeLexer(*inclompleteRules)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := failLexer.Tokenize(source); err == nil {
		t.Fatalf("successfully read invalid file '%s'", testSourceFile)
	}

	// correct lexer
	source, _ = os.Open(testSourceFile)
	table, err := lexer.Tokenize(source)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("\n%s\n", table.String())
}
