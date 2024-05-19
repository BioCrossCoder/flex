package list

func ParseCount(length int, counts ...int) int {
	if len(counts) == 0 {
		return 1
	}
	if counts[0] <= 0 {
		return length
	}
	return counts[0]
}

func SearchCount(length int, counts ...int) int {
	if len(counts) == 0 || counts[0] <= 0 {
		return length
	}
	return counts[0]
}
