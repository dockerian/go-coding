package trie

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	u "github.com/dockerian/go-coding/utils"
)

// -------- RuneTrie, constructor, and method receivers
// See https://en.wikipedia.org/wiki/Trie

// RuneTrie struct
type RuneTrie struct {
	root *RuneTrieNode
	ends rune
}

// NewRuneTrie constructs a RuneTrie with specific end and returns its pointer
func NewRuneTrie(endchar rune) *RuneTrie {
	var p = new(RuneTrie)
	p.root = NewRuneTrieNode(' ')
	p.ends = endchar
	return p
}

// FindMatchedPhases returns a list of phases match the prefix
func (t *RuneTrie) FindMatchedPhases(prefix string) []string {
	var phases = make([]string, 0)
	var stack = new(RuneTrieStack)
	hasMatch, _, node, _ := t.HasMatch(prefix)
	// u.Debug("stack: {%+v}, hasMatch = %v, node = %+v\n", stack, hasMatch, node)

	if hasMatch {
		stack.Push(&RuneTrieNodeItem{node, prefix})

		for stack.Size() > 0 {
			ptop := stack.Pop()
			item := *ptop
			p := item.node
			part := item.data
			if p.IsEnd() {
				phases = append(phases, part)
			}
			for _, child := range p.GetChildrenNodes() {
				data := part
				if !child.IsEnd() {
					data = fmt.Sprintf("%s%c", part, child.content)
				}
				stack.Push(&RuneTrieNodeItem{child, data})
			}
		}
	}

	sort.Sort(RunePhases(phases))
	return phases
}

// HasMatch checks if there is a matched prefix in the trie, additionaly
// returns the last matched node and matched length
func (t *RuneTrie) HasMatch(prefix string) (bool, bool, *RuneTrieNode, int) {
	var matched, matchedAll bool
	var s = strings.TrimSpace(prefix)
	if len(s) <= 0 {
		return matched, matchedAll, nil, 0
	}

	// make runes array with the most possible capacity (per phase length)
	var runes = []rune(prefix)
	var p = t.root
	i, bound := 0, len(runes)-1
	// u.Debug("root: %+v, runes: [%+v], bound: %d\n", p, string(runes), bound)
	for p != nil {
		next := p.GetChildNode(runes[i])
		// u.Debug("current: %+v, next: %+v, next rune: '%c', i==%d\n", p, next, runes[i], i)
		p = next
		if next == nil {
			break
		}
		if i == bound {
			matched = true
			matchedAll = next.Contains(t.ends)
			break
		}
		i++
	}
	u.Debug("prefix: '%s', matched: %v, all: %v, at[%d]: %+v\n", prefix, matched, matchedAll, i+1, p)
	return matched, matchedAll, p, i + 1
}

// Load is a RuneTrie pointer receiver to add a string phase to the trie
func (t *RuneTrie) Load(phase string) {
	var s = strings.TrimSpace(phase)
	if len(s) > 0 {
		var node = t.root
		for _, r := range s {
			node = node.Add(r)
		}
		node.Add(t.ends)
	}
}

// -------- RuneTrieNode, constructor, and method receivers

// RuneTrieNode struct
type RuneTrieNode struct {
	content  rune
	children map[rune]*RuneTrieNode
	keys     []rune
}

// NewRuneTrieNode constructs a RuneTrieNode and return its pointer
func NewRuneTrieNode(k rune) *RuneTrieNode {
	var p = new(RuneTrieNode)
	p.content = k
	p.children = make(map[rune]*RuneTrieNode)
	p.keys = make([]rune, 0)
	return p
}

// Add is a RuneTrieNode pointer receiver to add a new RuneTrieNode to children
func (n *RuneTrieNode) Add(k rune) *RuneTrieNode {
	p, okay := n.children[k]
	if !okay {
		p = NewRuneTrieNode(k)
		n.keys = append(n.keys, k)
		n.children[k] = p
	}
	return p
}

// Contains is a RuneTrieNode pointer receiver to check if a child key exists
func (n *RuneTrieNode) Contains(k rune) bool {
	return n.children[k] != nil
}

// Equals checks if two nodes have the same content
func (n *RuneTrieNode) Equals(p *RuneTrieNode) bool {
	return reflect.DeepEqual(*n, *p)
}

// GetChildNode is a RuneTrieNode pointer receiver to get a node by key
func (n *RuneTrieNode) GetChildNode(k rune) *RuneTrieNode {
	return n.children[k]
}

// GetChildren is a RuneTrieNode pointer method receiver to get node's children
func (n *RuneTrieNode) GetChildren() map[rune]*RuneTrieNode {
	return n.children
}

// GetChildrenKeys is a RuneTrieNode pointer receiver to get children's keys
func (n *RuneTrieNode) GetChildrenKeys() []rune {
	keys := make([]rune, len(n.children))
	i := 0
	for key := range n.children {
		keys[i] = key
		i++
	}
	return keys
}

// GetChildrenNodes is a RuneTrieNode pointer receiver to get children's nodes
func (n *RuneTrieNode) GetChildrenNodes() []*RuneTrieNode {
	nodes := make([]*RuneTrieNode, len(n.children))
	i := 0
	for _, v := range n.children {
		nodes[i] = v
		i++
	}
	return nodes
}

// GetContent is a RuneTrieNode pointer method receiver to get content of the node
func (n *RuneTrieNode) GetContent() rune {
	return n.content
}

// IsContent func returns true if the node content is as same as specified rune;
// otherwise, false
func (n *RuneTrieNode) IsContent(k rune) bool {
	return n.content == k
}

// IsEnd checks if the node is the phase end
func (n *RuneTrieNode) IsEnd() bool {
	return len(n.keys) == 0
}

// String func
func (n *RuneTrieNode) String() string {
	return fmt.Sprintf("'%c': [%+v]", n.content, string(n.GetChildrenKeys()))
}
