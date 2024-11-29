package strings

// Contains returns true if the target string is found in the source string.
//
// Parameters:
//   - source: the source string
//   - target: the target string
//
// Returns:
//   - bool: true if the target string is found in the source string, false otherwise
func Contains(source, target string) bool {
	length := len(target)
	if length > len(source) {
		return false
	}

	for i := 0; i <= len(source)-length; i++ {
		if source[i:i+length] == target {
			return true
		}
	}

	return false
}
