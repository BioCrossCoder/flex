// Package list provides functions for working with lists.
package list

// ParseCount function returns the count of elements based on the provided length and counts
func ParseCount(length int, counts ...int) int {
	if len(counts) == 0 {
		return 1
	}
	if counts[0] <= 0 {
		return length
	}
	return counts[0]
}

// SearchCount function returns the count of elements for searching based on the provided length and counts
func SearchCount(length int, counts ...int) int {
	if len(counts) == 0 || counts[0] <= 0 {
		return length
	}
	return counts[0]
}
