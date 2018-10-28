package search

func couldContain(v int, arr []int, i, j int) bool {
	if arr[i] < arr[j] {
		return v >= arr[i] && v <= arr[j]
	}
	return v <= arr[i] && v >= arr[j]
}

func FindInRotatedArray(v int, arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == v {
			return mid
		}

		if couldContain(v, arr, left, mid-1) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func findInRotateArrayRecursive(v int, arr []int, left, right int) int {
	if len(arr) == 0 || left > right {
		return -1
	}

	mid := (right + left) / 2
	if arr[mid] == v {
		return mid
	}
	if couldContain(v, arr, left, mid-1) {
		return findInRotateArrayRecursive(v, arr, left, mid-1)
	}
	return findInRotateArrayRecursive(v, arr, mid+1, right)
}

func FindInRotatedArrayRecursive(v int, arr []int) int {
	return findInRotateArrayRecursive(v, arr, 0, len(arr)-1)
}
