package str

import (
	"sort"
	"strings"
)

var (
	// ByCaseInsensitivity is a By function for string sorting
	ByCaseInsensitivity = By(LessByCaseInsensitivity)
	// ByLength is a By function for string sorting
	ByLength = By(LessByLength)
)

// ---------- Solution for case-insensitive sorting ----------

// ByCaseInsensitive sort case-insensitive strings
func ByCaseInsensitive(s []string) {
	sort.Sort(byCaseInsensitive(s))
}

// byCaseInsensitive implements sort.Interface for slices of case-insensitive strings
type byCaseInsensitive []string

// Len implements Len() in sort.Interface for byCaseInsensitive type
func (a byCaseInsensitive) Len() int {
	return len(a)
}

// Less implements Less() in sort.Interface for byCaseInsensitive type
func (a byCaseInsensitive) Less(x, y int) bool {
	return LessByCaseInsensitivity(&a[x], &a[y])
}

// Swap implements Swap() in sort.Interface for byCaseInsensitive
func (a byCaseInsensitive) Swap(x, y int) {
	a[x], a[y] = a[y], a[x]
}

// ---------- Solution for abnormal string sorting ----------

// By is a "Less" function type defines how to sort strings
type By func(p1, p2 *string) bool

type stringSorter struct {
	data []string
	by   func(p1, p2 *string) bool
}

// Sort is a sorting method for By
func (by By) Sort(data []string) {
	sortable := &stringSorter{
		data: data,
		by:   by,
	}
	sort.Sort(sortable)
}

// Len implements Len() in sort.Interface for stringSorter
func (s *stringSorter) Len() int {
	return len(s.data)
}

// Swap implements Swap() in sort.Interface for stringSorter
func (s *stringSorter) Less(x, y int) bool {
	return s.by(&s.data[x], &s.data[y])
}

// Swap implements Swap() in sort.Interface for stringSorter
func (s *stringSorter) Swap(x, y int) {
	s.data[x], s.data[y] = s.data[y], s.data[x]
}

// ---------- Generic string comparison functions ----------

// ComparebyCaseInsensitive compares strings by case-insensitivity
func ComparebyCaseInsensitive(a, b *string) int {
	return strings.Compare(strings.ToLower(*a), strings.ToLower(*b))
}

// CompareByLength compares strings by length
func CompareByLength(a, b *string) int {
	if len(*a) == len(*b) {
		return 0
	} else if len(*a) > len(*b) {
		return 1
	}
	return -1
}

// ---------- Generic string Less() functions ----------

// LessByCaseInsensitivity compares strings by case-insensitivity
func LessByCaseInsensitivity(a, b *string) bool {
	return ComparebyCaseInsensitive(a, b) == -1
}

// LessByLength compares strings by length
func LessByLength(a, b *string) bool {
	return CompareByLength(a, b) == -1
}
