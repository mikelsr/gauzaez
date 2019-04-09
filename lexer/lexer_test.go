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
	testSourceFileContent = []byte("x = 1 +   2.3\ny = 4 ** 5\n\ntext=\"hi\"")
	if err := ioutil.WriteFile(testSourceFile, testSourceFileContent, 0644); err != nil {
		log.Fatal(err)
	}

	// load rules
	rules, _ = MakeRules(rulesFile)

	// configure lexer
	lexer, _ = MakeLexer(testSourceFile, *rules)
	exitCode := m.Run()
	os.Remove(testSourceFile)
	os.Exit(exitCode)
}

func TestLexer_loadFile(t *testing.T) {
	t.Log("Loading file contents...")
	lexer.loadFile()
	if len(lexer.chars) != len(testSourceFileContent) {
		t.Errorf("Test and file content do not match.\n`%s`, found %d, expected %d\n.",
			testSourceFile, len(lexer.chars), len(testSourceFileContent))
	}
}
