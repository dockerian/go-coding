// Package tree :: tree.go
package tree

import (
	"fmt"
	"strings"

	"github.com/dockerian/go-coding/ds/queue"
)

// BT interface includes basic functions for a binary tree item
type BT interface {
	Left() BT
	Right() BT
	Value() interface{}
	SetLeft(BT)
	SetRight(BT)
	SetValue(interface{})
	Find(interface{}) BT
	Has(interface{}) bool
}

// BTNode represents a binary tree node
type BTNode struct {
	left  BT
	right BT
	value interface{}
}

// NewBTNode constructs an BTNode instance
func NewBTNode(v interface{}, l, r BT) *BTNode {
	var o = &BTNode{value: v, left: l, right: r}
	return o
}

// Left implements BT
func (o *BTNode) Left() BT {
	return o.left
}

// SetLeft implements BT
func (o *BTNode) SetLeft(node BT) {
	o.left = node
}

// Right implements BT
func (o *BTNode) Right() BT {
	return o.right
}

// SetRight implements BT
func (o *BTNode) SetRight(node BT) {
	o.right = node
}

// SetValue implements BT
func (o *BTNode) SetValue(v interface{}) {
	o.value = v
}

// Value implements BT
func (o *BTNode) Value() interface{} {
	return o.value
}

// Find returns the node matches to the value v
func (o *BTNode) Find(v interface{}) BT {
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
func (o *BTNode) Has(v interface{}) bool {
	if o == nil {
		return false
	}
	hasInLeft := o.Left() != nil && o.Left().Has(v)
	hasInRight := o.Right() != nil && o.Right().Has(v)
	return o.value == v || hasInLeft || hasInRight
}

// String returns a string representation of StrBTNode object
func (o *BTNode) String() string {
	if o == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v -> (%v, %v)", o.Value(), o.Left(), o.Right())
}

/*******************************************************************************
 * Integer node and tree
 *******************************************************************************
 */

// IntBT interface for any binary tree item contains int value
type IntBT interface {
	BT
	Sum() int
}

// IntBTNode represents an integer (int) binary tree node
type IntBTNode struct {
	BTNode
}

// NewIntBTNode constructs an IntBTNode instance
func NewIntBTNode(v int, l, r IntBT) *IntBTNode {
	var o IntBT = &IntBTNode{}
	o.SetLeft(l)
	o.SetRight(r)
	o.SetValue(v)
	return o.(*IntBTNode)
}

// isMirror checks if two BT are in mirror recursively.
func isMirror(left, right BT) bool {
	if left == nil && right == nil {
		return true
	}
	if left.Value() == right.Value() {
		return isMirror(left.Left(), right.Right()) &&
			isMirror(left.Right(), right.Left())
	}
	return false
}

// IsMirror checks if a binary tree is mirror of itself in recursive mode.
func (o *BTNode) IsMirror() bool {
	return o == nil || isMirror(o.Left(), o.Right())
}

// IsSymmetric checks if a binary tree is symmetric in interative mode.
func (o *BTNode) IsSymmetric() bool {
	if o == nil {
		return false
	}
	var q = queue.NewQueue()
	var right = o.Right()
	var left = o.Left()

	q.Enqueue(left)
	q.Enqueue(right)

	for len(q) > 0 {
		right = q.Dequeue().(*BTNode)
		left = q.Dequeue().(*BTNode)

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Value() != right.Value() {
			return false
		}
		q.Enqueue(left.Left())
		q.Enqueue(right.Right())
		q.Enqueue(left.Right())
		q.Enqueue(right.Left())
	}

	return true
}

// Sum adds up values in all tree nodes
func (it *IntBTNode) Sum() int {
	if it == nil {
		return 0
	}
	v := it.value.(int)
	var sumInLeft, sumInRight int
	if it.Left() != nil {
		sumInLeft = it.Left().(*IntBTNode).Sum()
	}
	if it.Right() != nil {
		sumInRight = it.Right().(*IntBTNode).Sum()
	}
	return v + sumInLeft + sumInRight
}

/*******************************************************************************
 * String node and tree
 *******************************************************************************
 */

// StrBT interface for any binary tree item contains string value
type StrBT interface {
	BT
	Join(string) string
}

// StrBTNode represents an integer (int) binary tree node
type StrBTNode struct {
	BTNode
}

// NewStrBTNode constructs an StrBTNode instance
func NewStrBTNode(v string, l, r StrBT) *StrBTNode {
	var o StrBT = &StrBTNode{}
	o.SetLeft(l)
	o.SetRight(r)
	o.SetValue(v)
	return o.(*StrBTNode)
}

// Join concatenates values in all tree nodes
func (it *StrBTNode) Join(sep string) string {
	if it == nil {
		return ""
	}
	v := it.value.(string)
	var strInLeft, strInRight string
	if it.Left() != nil {
		if l, ok := it.Left().(*StrBTNode); ok {
			strInLeft = l.Join(sep)
		}
	}
	if it.Right() != nil {
		if r, ok := it.Right().(*StrBTNode); ok {
			strInRight = r.Join(sep)
		}
	}
	result := strings.Trim(strings.Join([]string{v, strInLeft, strInRight}, sep), sep)
	return strings.Replace(result, sep+sep, sep, -1)
}
