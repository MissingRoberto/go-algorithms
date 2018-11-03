package sortalgorithms

func swap(arr []int, i, j int) { arr[i], arr[j] = arr[j], arr[i] }

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := arr[len(arr)-1]
	left, right := 0, len(arr)-2
	for left < right {
		if arr[left] < pivot {
			left++
		} else {
			swap(arr, left, right)
			right--
		}
	}
	swap(arr, left, len(arr)-1)
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
