package maps

import (
	"fmt"
	"strings"

	"github.com/dockerian/go-coding/ds/mathEx"
)

// LinkedList struct
type LinkedList struct {
	Data interface{}
	Next *LinkedList
}

// LinkedListNumber struct, the Data holds a single decimal digit
type LinkedListNumber struct {
	Data uint8
	Next *LinkedListNumber
}

// Add method adds LinkedListNumber objects.
func (pList *LinkedListNumber) Add(aList *LinkedListNumber) *LinkedListNumber {
	sumLinkedList := &LinkedListNumber{0, nil}
	ps := sumLinkedList
	p1 := pList
	p2 := aList

	for p1 != nil || p2 != nil {
		// make sure at least one is not nil
		if p1 == nil {
			p1 = p2
			p2 = nil
		}
		var data1 = p1.Data
		var data2 uint8

		// get the data, and move to next item
		if p2 != nil {
			data2 = p2.Data
			p2 = p2.Next
		}
		p1 = p1.Next

		sum := data1 + data2 + ps.Data
		mod := sum % 10
		hcb := sum / 10 // carrying bit
		ps.Data = mod

		if p1 != nil || p2 != nil || hcb > 0 {
			ps.Next = &LinkedListNumber{hcb, nil}
			ps = ps.Next
		}
	}

	return sumLinkedList
}

// AddNumber method adds an integer to LinkedListNumber
func (pList *LinkedListNumber) AddNumber(number uint64) *LinkedListNumber {
	aList := GetLinkedListNumber(number)

	return pList.Add(aList)
}

// Compare method compares LinkedListNumber objects
func (pList *LinkedListNumber) Compare(aList *LinkedListNumber) int {
	s1 := pList.String()
	s2 := aList.String()

	if len(s1) > len(s2) {
		return 1
	}
	if len(s1) < len(s2) {
		return -1
	}

	return strings.Compare(s1, s2)
}

// GetLength method gets the length of LinkedListNumber
func (pList *LinkedListNumber) GetLength() int {
	len := 0
	ptr := pList
	for ptr != nil {
		len++
		ptr = ptr.Next
	}
	return len
}

// GetMiddle method returns the middle item of LinkedListNumber

// GetNumber method gets integer representation of LinkedListNumber.
func (pList *LinkedListNumber) GetNumber() (uint64, error) {
	var n uint64
	var x uint64 = 1
	p := pList
	for p != nil {
		u := uint(p.Data)
		m, err := mathEx.MultiplyUint64(x, uint64(u))
		if err != nil {
			return 0, err
		}
		n, err = mathEx.SumUint64(n, m)
		if err != nil {
			return 0, err
		}
		p = p.Next
		x *= 10
	}
	return n, nil
}

// String method
func (pList *LinkedListNumber) String() string {
	ptr := pList
	len := pList.GetLength()
	str := make([]byte, len)
	idx := len - 1

	for ptr != nil {
		str[idx] = '0' + byte(ptr.Data)
		ptr = ptr.Next
		idx--
	}

	return string(str)
}

// GetLinkedListNumber converts a positive number to LinkedListNumber.
func GetLinkedListNumber(number uint64) *LinkedListNumber {
	n := &LinkedListNumber{0, nil}
	p := n
	for number != 0 {
		mod := uint8(number % 10)
		number /= 10
		p.Data = mod

		if number != 0 {
			p.Next = &LinkedListNumber{0, nil}
			p = p.Next
		}
	}

	return n
}

// GetLinkedListNumberFromString converts a number string to LinkedListNumber.
func GetLinkedListNumberFromString(input string) (*LinkedListNumber, error) {
	n := &LinkedListNumber{0, nil}
	p := n

	input = strings.TrimSpace(input)
	input = strings.TrimLeft(input, "0")

	for k := len(input); k >= 0; k-- {
		val := input[k] - '0'
		if 0 <= val && val >= 9 {
			p.Data = val
		} else {
			return nil, fmt.Errorf("invalid number string: '%v'", input)
		}
		if k > 0 {
			p.Next = &LinkedListNumber{0, nil}
			p = p.Next
		}
	}

	return n, nil
}

// GetNumberFromLinkedList converts LinkedListNumber to integer.
func GetNumberFromLinkedList(pList *LinkedListNumber) (uint64, error) {
	if pList != nil {
		return pList.GetNumber()
	}
	return 0, nil
}
