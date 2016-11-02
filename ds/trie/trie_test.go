// +build all ds tree trie test

package trie

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	u "github.com/dockerian/go-coding/utils"
	"github.com/stretchr/testify/assert"
)

const (
	eng1000wDict = "en-1000.dict"
	eng4000wDict = "en-4000.dict"
	trieTestDict = "trie_test.dict"
	trieTestFile = "trie_test.json"
)

// StringTrieTest struct
type StringTrieTest struct {
	data []StringTrieTestCase
	trie *Trie
}

// StringTrieTestCase struct
type StringTrieTestCase struct {
	Prefix string   `json:"prefix,omitempty"`
	Phases []string `json:"result,omitempty"`
}

// LoadTrieDictionary loads dictionary to trie
func LoadTrieDictionary(t *testing.T, trie *Trie, dictFile string) error {
	testName := u.GetTestName(t)
	dict, err1 := os.Open(dictFile)
	if err1 != nil {
		return fmt.Errorf("Cannot open dict file '%v': %v\n", dictFile, err1)
	}
	defer dict.Close()

	t.Logf("Test [%s]: Loading dictionary: %s\n", testName, dictFile)
	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		phase := scanner.Text()
		trie.Load(phase)
	}

	if err2 := scanner.Err(); err2 != nil {
		return fmt.Errorf("Error in reading dict '%v': %v\n", dictFile, err2)
	}

	return nil
}

// LoadTrieTest func retruns a StringTrieTest for TestTrie
func LoadTrieTest(t *testing.T, dictFiles []string, testFile string) *StringTrieTest {
	trie := NewTrie()
	data := []StringTrieTestCase{}

	file, err1 := os.Open(testFile)
	if err1 != nil {
		t.Errorf("Cannot open test file '%v': %v\n", testFile, err1)
		t.Fail()
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err2 := decoder.Decode(&data); err2 != nil {
		t.Errorf("Cannot decode test file '%v': %v\n", testFile, err2)
		t.Fail()
	}

	for _, dictFile := range dictFiles {
		err3 := LoadTrieDictionary(t, trie, dictFile)
		if err3 != nil {
			t.Errorf(err3.Error())
			t.Fail()
		}
	}

	return &StringTrieTest{data, trie}
}

// TestTrie func tests Trie
func TestTrie(t *testing.T) {
	dicts := []string{eng1000wDict, trieTestDict}
	runeTest := LoadTrieTest(t, dicts, trieTestFile)

	for index, test := range runeTest.data {
		var set = runeTest.trie.FindMatchedPhases(test.Prefix)
		var msg = fmt.Sprintf("expecting '%v' => %+v", test.Prefix, test.Phases)
		t.Logf("Test %03d: %v\n", index+1, msg)
		assert.Equal(t, test.Phases, set, msg)
	}
}
