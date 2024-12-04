package ext

// Ternary is a simple ternary operator for Go.
// It returns the first value if the condition is true, otherwise the second value.
//
// Parameters:
//   - condition: the condition to check
//   - first: the value to return if the condition is true
//   - second: the value to return if the condition is false
//
// Returns:
//   - the first value if the condition is true, otherwise the second value
func Ternary[T any](condition bool, first T, second T) T {
	if condition {
		return first
	}
	return second
}
