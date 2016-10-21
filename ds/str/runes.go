package str

// Runes type represents array of runes, aka `[]rune`
type Runes []rune

// Reverse method returns a copy of reversed Runes/[]runes
func (s Runes) Reverse() Runes {
	sz := len(s)
	rs := make(Runes, sz)
	if sz > 0 {
		for n := 0; n <= sz/2; n++ {
			rs[n], rs[sz-n-1] = s[sz-n-1], s[n]
		}
	}
	return rs
}

func (s Runes) String() string {
	return string(s)
}
