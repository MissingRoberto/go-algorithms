package sortalgorithms

func merge(left, right []int) []int {
	mergedArray := make([]int, len(left)+len(right))
	var i, j, k int
	for k < len(mergedArray) {
		var v int
		if i >= len(left) {
			v = right[j]
			j++
		} else if j >= len(right) {
			v = left[i]
			i++
		} else if left[i] <= right[j] {
			v = left[i]
			i++
		} else {
			v = right[j]
			j++
		}
		mergedArray[k] = v
		k++
	}
	return mergedArray
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}
