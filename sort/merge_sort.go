package sort

//go:generate gotests -all -w .

func MergeSort(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	tmp := make([]int, len(arr))
	merge_sort(arrCopy, tmp, 0, len(arrCopy)-1)
	return arrCopy
}

func merge_sort(arr, tmp[]int, l, r int) {
	if r <= l {
		return
	}
	m := (r+l)/2
	merge_sort(arr, tmp, l, m)
	merge_sort(arr, tmp, m+1, r)
	merge(arr, tmp, l, m , r)
}

func merge(arr, tmp []int, l, m, r int) {
	t := 0
	lp, rp := l, m+1
	for lp <= m && rp <= r {
		if arr[lp] < arr[rp] {
			tmp[t] = arr[lp]
			t++
			lp++
		}else {
			tmp[t] = arr[rp]
			t++
			rp++
		}
	}
	for lp <= m {
		tmp[t] = arr[lp]
		t++
		lp++
	}
	for rp <= r {
		tmp[t] = arr[rp]
		rp++
		t++
	}
	for i:=0; i < t; i++ {
		arr[l+i] = tmp[i]
	}
}