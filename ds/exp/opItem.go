package exp

import "fmt"

// OpItem struct
type OpItem struct {
	Left       *OpItem
	Right      *OpItem
	Parent     *OpItem
	Expression string
	Op         string
}

func (o *OpItem) addLeft(s string) *OpItem {
	item := OpItem{Op: s, Expression: s, Parent: o}
	o.Left = &item
	return &item
}

func (o *OpItem) addRight(s string) *OpItem {
	item := OpItem{Op: s, Expression: s, Parent: o}
	o.Right = &item
	return &item
}

// String function
func (o *OpItem) String() string {
	sl, sr, sp, ex := "", "", "", ""
	if o.Left != nil {
		sl = o.Left.Expression
	}
	if o.Right != nil {
		sr = o.Right.Expression
	}
	if o.Parent != nil {
		sp = o.Parent.Expression
	}
	if o.Expression != "" {
		ex = fmt.Sprintf("['%v'] ", o.Expression)
	}
	return fmt.Sprintf("%v %v{%v,%v => %v}", o.Op, ex, sl, sr, sp)
}
