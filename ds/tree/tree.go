// Package tree :: tree.go
package tree

import (
	"fmt"
	"strings"
)

// BST interface includes basic functions for a binary search tree item
type BST interface {
	Left() BST
	Right() BST
	Value() interface{}
	SetLeft(BST)
	SetRight(BST)
	SetValue(interface{})
	Find(interface{}) BST
	Has(interface{}) bool
}

// BSTNode represents a binary search tree node
type BSTNode struct {
	left  BST
	right BST
	value interface{}
}

// NewBSTNode constructs an BSTNode instance
func NewBSTNode(v interface{}, l, r BST) *BSTNode {
	var o = &BSTNode{value: v, left: l, right: r}
	return o
}

// Left implements BST
func (o *BSTNode) Left() BST {
	return o.left
}

// SetLeft implements BST
func (o *BSTNode) SetLeft(node BST) {
	o.left = node
}

// Right implements BST
func (o *BSTNode) Right() BST {
	return o.right
}

// SetRight implements BST
func (o *BSTNode) SetRight(node BST) {
	o.right = node
}

// SetValue implements BST
func (o *BSTNode) SetValue(v interface{}) {
	o.value = v
}

// Value implements BST
func (o *BSTNode) Value() interface{} {
	return o.value
}

// Find returns the node matches to the value v
func (o *BSTNode) Find(v interface{}) BST {
	if o == nil {
		return nil
	}
	if o.value == v {
		return o
	} else if left := o.Left(); left != nil {
		return left.Find(v)
	} else if right := o.Right(); right != nil {
		return right.Find(v)
	}
	return nil
}

// Has returns true if value v is found in any node; otherwise false
func (o *BSTNode) Has(v interface{}) bool {
	if o == nil {
		return false
	}
	hasInLeft := o.Left() != nil && o.Left().Has(v)
	hasInRight := o.Right() != nil && o.Right().Has(v)
	return o.value == v || hasInLeft || hasInRight
}

// String returns a string representation of StrBSTNode object
func (o *BSTNode) String() string {
	if o == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v -> (%v, %v)", o.Value(), o.Left(), o.Right())
}

/*******************************************************************************
 * Integer node and tree
 *******************************************************************************
 */

// IntBST interface for any binary search tree item contains int value
type IntBST interface {
	BST
	Sum() int
}

// IntBSTNode represents an integer (int) binary search tree node
type IntBSTNode struct {
	BSTNode
}

// NewIntBSTNode constructs an IntBSTNode instance
func NewIntBSTNode(v int, l, r IntBST) *IntBSTNode {
	var o IntBST = &IntBSTNode{}
	o.SetLeft(l)
	o.SetRight(r)
	o.SetValue(v)
	return o.(*IntBSTNode)
}

// Sum adds up values in all tree nodes
func (it *IntBSTNode) Sum() int {
	if it == nil {
		return 0
	}
	v := it.value.(int)
	var sumInLeft, sumInRight int
	if it.Left() != nil {
		sumInLeft = it.Left().(*IntBSTNode).Sum()
	}
	if it.Right() != nil {
		sumInRight = it.Right().(*IntBSTNode).Sum()
	}
	return v + sumInLeft + sumInRight
}

/*******************************************************************************
 * String node and tree
 *******************************************************************************
 */

// StrBST interface for any binary search tree item contains string value
type StrBST interface {
	BST
	Join(string) string
}

// StrBSTNode represents an integer (int) binary search tree node
type StrBSTNode struct {
	BSTNode
}

// NewStrBSTNode constructs an StrBSTNode instance
func NewStrBSTNode(v string, l, r StrBST) *StrBSTNode {
	var o StrBST = &StrBSTNode{}
	o.SetLeft(l)
	o.SetRight(r)
	o.SetValue(v)
	return o.(*StrBSTNode)
}

// Join concatenates values in all tree nodes
func (it *StrBSTNode) Join(sep string) string {
	if it == nil {
		return ""
	}
	v := it.value.(string)
	var strInLeft, strInRight string
	if it.Left() != nil {
		if l, ok := it.Left().(*StrBSTNode); ok {
			strInLeft = l.Join(sep)
		}
	}
	if it.Right() != nil {
		if r, ok := it.Right().(*StrBSTNode); ok {
			strInRight = r.Join(sep)
		}
	}
	result := strings.Trim(strings.Join([]string{v, strInLeft, strInRight}, sep), sep)
	return strings.Replace(result, sep+sep, sep, -1)
}
