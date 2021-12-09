package day08

// Any of the provided functions return true.
func Any(value string, fns ...func(string) bool) bool {
	for _, fn := range fns {
		if fn(value) {
			return true
		}
	}

	return false
}

// RuneDiff returns with the difference between two strings.
func RuneDiff(value1, value2 string) []rune {
	collection := map[rune]int{}
	diff := []rune{}

	for _, character := range value1 {
		collection[character]++
	}

	for _, character := range value2 {
		collection[character]++
	}

	for character, count := range collection {
		if count == 1 {
			diff = append(diff, character)
		}
	}

	return diff
}
