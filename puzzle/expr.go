package puzzle

// CheckExpression checks if an expression has proper openings and closings of
// e.g. `()`, `[]`, and `{}`.
// TODO: supporting any string pairs, e.g. "<head>": "</head>"
func CheckExpression(expr string) bool {
	history := make([]rune, 0, 16)
	closing := map[rune]rune{}
	opening := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for k, v := range opening {
		closing[v] = k
	}

	for _, chr := range expr {
		if v, ok := opening[chr]; ok {
			history = append(history, v)
		} else if _, ok := closing[chr]; ok {
			bound := len(history) - 1
			if bound >= 0 {
				if chr == history[bound] {
					history = history[0:bound]
					continue
				}
				return false
			}
			return false
		}
	}
	return len(history) == 0
}
