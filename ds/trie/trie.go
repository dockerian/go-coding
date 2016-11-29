package trie

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	u "github.com/dockerian/go-coding/utils"
)

// -------- Node, constructor, and method receivers

// Node struct
type Node struct {
	key      string
	keys     []string
	children map[string]*Node
	data     interface{}
	ends     bool
}

// NewNode constructs a Node and return its pointer
func NewNode(k string) *Node {
	var p = new(Node)
	p.key = k
	p.children = make(map[string]*Node)
	p.keys = make([]string, 0)
	return p
}

// Add is a Node pointer receiver to add a new Node to children
func (n *Node) Add(k string) *Node {
	p, okay := n.children[k]
	if !okay {
		p = NewNode(k)
		n.keys = append(n.keys, k)
		n.children[k] = p
	}
	return p
}

// Contains is a Node pointer receiver to check if a child key exists
func (n *Node) Contains(k string) bool {
	return n.children[k] != nil
}

// Equals checks if two nodes have the same content
func (n *Node) Equals(p *Node) bool {
	return reflect.DeepEqual(*n, *p)
}

// GetChildNode is a Node pointer receiver to get a node by key
func (n *Node) GetChildNode(k string) *Node {
	return n.children[k]
}

// GetChildren is a Node pointer method receiver to get node's children
func (n *Node) GetChildren() map[string]*Node {
	return n.children
}

// GetChildrenKeys is a Node pointer receiver to get children's keys
func (n *Node) GetChildrenKeys() []string {
	keys := make([]string, len(n.children))
	i := 0
	for key := range n.children {
		keys[i] = key
		i++
	}
	return keys
}

// GetChildrenNodes is a Node pointer receiver to get children's nodes
func (n *Node) GetChildrenNodes() []*Node {
	nodes := make([]*Node, len(n.children))
	i := 0
	for _, v := range n.children {
		nodes[i] = v
		i++
	}
	return nodes
}

// GetContent is a Node pointer method receiver to get content of the node
func (n *Node) GetContent() string {
	return n.key
}

// IsContent func returns true if the node content is as same as specified string;
// otherwise, false
func (n *Node) IsContent(k string) bool {
	return n.key == k
}

// IsEnd checks if the node is the phase end
func (n *Node) IsEnd() bool {
	return n.ends
}

// String func
func (n *Node) String() string {
	return fmt.Sprintf("'%s': [%+v]", n.key, n.GetChildrenKeys())
}

// -------- Trie, constructor, and method receivers

// Trie struct
type Trie struct {
	root *Node
}

// NewTrie constructs a Trie with specific end and returns its pointer
func NewTrie() *Trie {
	var p = new(Trie)
	p.root = NewNode("")
	return p
}

// FindMatchedPhases returns a list of phases match the prefix
func (t *Trie) FindMatchedPhases(prefix string) []string {
	var phases = make([]string, 0)
	var stack = new(NodeStack)
	hasMatch, _, node, _ := t.HasMatch(prefix)
	// u.Debug("stack: {%+v}, hasMatch = %v, node = %+v\n", stack, hasMatch, node)

	if hasMatch {
		stack.Push(&NodeItem{node, prefix})

		for stack.Size() > 0 {
			ptop := stack.Pop()
			item := *ptop
			p := item.node
			part := item.data
			if p.IsEnd() {
				phases = append(phases, part)
			}
			for _, child := range p.GetChildrenNodes() {
				stack.Push(&NodeItem{child, part + child.key})
			}
		}
	}

	sort.Sort(RunePhases(phases))
	return phases
}

// HasMatch checks if there is a matched prefix in the trie, additionaly
// returns the last matched node and matched length
func (t *Trie) HasMatch(prefix string) (bool, bool, *Node, int) {
	var matched, matchedAll bool
	var s = strings.TrimSpace(prefix)
	if len(s) <= 0 {
		return matched, matchedAll, nil, 0
	}

	// make strings array with the most possible capacity (per phase length)
	var runes = []rune(prefix)
	var p = t.root
	i, bound := 0, len(runes)-1
	// u.Debug("root: %+v, runes: [%+v], bound: %d\n", p, string(runes), bound)
	for p != nil {
		next := p.GetChildNode(string(runes[i]))
		// u.Debug("current: %+v, next: %+v, next string: '%c', i==%d\n", p, next, strings[i], i)
		p = next
		if next == nil {
			break
		}
		if i == bound {
			matched = true
			matchedAll = next.ends
			break
		}
		i++
	}
	u.Debug("prefix: '%s', matched: %v, all: %v, at[%d]: %+v\n", prefix, matched, matchedAll, i+1, p)
	return matched, matchedAll, p, i + 1
}

// Load is a Trie pointer receiver to add a string phase to the trie
func (t *Trie) Load(phase string) {
	var s = strings.TrimSpace(phase)
	if len(s) > 0 {
		var node = t.root
		for _, r := range s {
			node = node.Add(string(r))
		}
		node.ends = true
	}
}
