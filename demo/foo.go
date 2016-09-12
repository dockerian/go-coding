package demo

import (
	"fmt"
)

// Foo struct is a template of using interface data
type Foo struct {
	Anything interface{}
}

func (f *Foo) GetAnything() interface{} {
	return f.Anything
}

func (f *Foo) string() string {
	return fmt.Sprintf("%v", f.Anything)
}
