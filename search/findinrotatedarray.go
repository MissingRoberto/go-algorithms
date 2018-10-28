package search

type RotatedArray []int

func (a RotatedArray) IsSorted(i, j int) bool {
	if i == j {
		return true
	}
	return a[j] > a[i]
}

func (a RotatedArray) MayContain(v, i, j int) bool {
	if i == j {
		return v == a[i]
	}
	return a[j] >= v && v >= a[i]
}

func (a RotatedArray) IsValid(i, j int) bool {
	return i >= 0 && j < len(a) && i <= j
}

func findInRotatedArray(v int, arr RotatedArray, start, end int) int {
	mid := (start + end) / 2

	if arr[mid] == v {
		return mid
	}

	if arr.IsValid(start, mid-1) && arr.IsSorted(start, mid-1) {
		if arr.MayContain(v, start, mid-1) {
			return findInRotatedArray(v, arr, start, mid-1)
		}
		return findInRotatedArray(v, arr, mid+1, end)
	}

	if arr.IsValid(start, mid+1) && arr.IsSorted(mid+1, end) {
		if arr.MayContain(v, mid+1, end) {
			return findInRotatedArray(v, arr, mid+1, end)
		}
		return findInRotatedArray(v, arr, start, mid-1)
	}
	return -1
}

func FindInRotatedArray(v int, arr []int) int {
	return findInRotatedArray(v, RotatedArray(arr), 0, len(arr)-1)
}
