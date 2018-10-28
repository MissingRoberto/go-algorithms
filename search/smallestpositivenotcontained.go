package search

func FindSmallestPositiveNotContained(arr []int) int {
	found := make(map[int]bool)
	for _, v := range arr {
		found[v] = true
	}

	minimum := 1
	for found[minimum] {
		minimum++
	}
	return minimum
}
