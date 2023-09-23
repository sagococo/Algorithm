package sort

//go:generate gotests -all -w .

func BubbleSort(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	bubble_sort(arrCopy)
	return arrCopy
}

func bubble_sort(arr []int) {
	r := len(arr)
	for r > 0 {
		flag := true
		for i := 0; i+1 < r; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = false
			}
		}
		if flag {
			return
		}
		r--
	}
}
