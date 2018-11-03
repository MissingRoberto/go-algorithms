package algorithms

import (
	"sort"
)

// O(N) space, O(N) memory
func FindPairHash(arr []int, sum int) (int, int) {
	found := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if v, ok := found[arr[i]]; ok {
			return v, i
		}

		found[sum-arr[i]] = i
	}
	return -1, -1
}

// O(NlogN) + O(N)
func FindPairWithSort(arr []int, sum int) (int, int) {
	sort.Ints(arr)

	i, j := 0, len(arr)-1
	for i < j {
		complementary := sum - arr[j]
		if arr[i] == complementary {
			return i, j
		}

		if arr[i] < complementary {
			i++
		} else {
			j--
		}
	}

	return -1, -1
}
