package ext

// Red returns the string s wrapped in ANSI color codes for red.
//
// Parameters:
//   - s: the string to color
//
// Returns:
//   - the string s wrapped in ANSI color codes for red
func Red(s string) string {
	return "\033[31m" + s + "\033[0m"
}

// Green returns the string s wrapped in ANSI color codes for green.
//
// Parameters:
//   - s: the string to color
//
// Returns:
//   - the string s wrapped in ANSI color codes for green
func Green(s string) string {
	return "\033[32m" + s + "\033[0m"
}

// Yellow returns the string s wrapped in ANSI color codes for yellow.
//
// Parameters:
//   - s: the string to color
//
// Returns:
//   - the string s wrapped in ANSI color codes for yellow
func Yellow(s string) string {
	return "\033[33m" + s + "\033[0m"
}
