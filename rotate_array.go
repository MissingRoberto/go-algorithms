package algorithms

func reverse(arr *[]int, i, j int) {
	for i < j {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		i++
		j--
	}
}

func Rotate(n int, arr *[]int) {
	l := len(*arr)
	n %= l
	reverse(arr, 0, l-1-n)
	reverse(arr, l-n, l-1)
	reverse(arr, 0, l-1)
}
