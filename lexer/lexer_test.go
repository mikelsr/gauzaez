package lexer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const testSourceFile = "/tmp/gauzaez_lexer_test.txt"

var rulesFile = fmt.Sprintf("%s/src/github.com/mikelsr/gauzaez/conf/"+
	"lexer_rules.json", os.Getenv("GOPATH"))

var testSourceFileContent []byte
var lexer *Lexer
var rules *Rules

func LoadFile(t *testing.T) {
	t.Log("Loading file contents...")

	lexer.loadFile()
	if len(lexer.chars) != len(testSourceFileContent) {
		t.Errorf("Test and file content do not match.\n`%s`, found %d, expected %d\n.",
			testSourceFile, len(lexer.chars), len(testSourceFileContent))
	}
}

// TestMain runs preparations for the other tests
func TestMain(m *testing.M) {

	testSourceFileContent = []byte("x = 1 +   2.3\ny = 4 ** 5\n\ntext=\"hi\"")
	writetestSourceFile()

	// load rules
	rules, _ = MakeRules(rulesFile)

	// configure lexer
	lexer, _ = MakeLexer(testSourceFile, rules)
	exitCode := m.Run()
	os.Remove(testSourceFile)
	os.Exit(exitCode)
}

// Write some bytes to a test source file so it can be loaded to a lexer
func writetestSourceFile() {
	err := ioutil.WriteFile(testSourceFile, testSourceFileContent, 0644)
	if err != nil {
		errStr := fmt.Sprintf("Couldn't write bytes to test file `%s`\n",
			testSourceFile)
		panic(errors.New(errStr))
	}
}
