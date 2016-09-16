package utils

// GetTernary returns a if condition is true; otherwise returns b
func GetTernary(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
