// +build all ds tree trie test

package trie

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	runeTrieDict = "runeTrie_test.dic"
	runeTrieTest = "runeTrie_test.json"
	phaseEndChar = '$'
)

// RuneTrieTest struct
type RuneTrieTest struct {
	data []RuneTrieTestCase
	trie *RuneTrie
}

// RuneTrieTestCase struct
type RuneTrieTestCase struct {
	Prefix   string   `json:"a,omitempty"`
	HasMatch bool     `json:"b,omitempty"`
	Phases   []string `json:"c,omitempty"`
}

// getRuneTrieTest func retruns a RuneTrieTest for TestRuneTrie
func getRuneTrieTest(t *testing.T, dictFile, testFile string) *RuneTrieTest {
	trie := NewRuneTrie(phaseEndChar)
	data := []RuneTrieTestCase{}

	file, err1 := os.Open(testFile)
	if err1 != nil {
		t.Errorf("Cannot open test file '%v': %v\n", testFile, err1)
		t.Fail()
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err2 := decoder.Decode(&data); err2 != nil {
		t.Errorf("Cannot decode test file '%v': %v\n", testFile, err2)
	}

	dict, err3 := os.Open(dictFile)
	if err3 != nil {
		t.Errorf("Cannot open dict file '%v': %v\n", dictFile, err3)
	}
	defer dict.Close()

	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		phase := scanner.Text()
		trie.Load(phase)
	}

	if err4 := scanner.Err(); err4 != nil {
		t.Errorf("Error in reading dict '%v': %v\n", dictFile, err4)
	}

	return &RuneTrieTest{data, trie}
}

// TestRuneTrie func tests RuneTrie
func TestRuneTrie(t *testing.T) {
	runeTest := getRuneTrieTest(t, runeTrieDict, runeTrieTest)

	for index, test := range runeTest.data {
		var set = runeTest.trie.FindMatchedPhases(test.Prefix)
		var msg = fmt.Sprintf("expecting '%v' => %+v", test.Prefix, test.Phases)
		t.Logf("Test %03d: %v\n", index+1, msg)
		assert.Equal(t, test.HasMatch, len(set) > 0, msg)
		assert.Equal(t, test.Phases, set, msg)
	}
}
