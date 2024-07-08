package twoPointers

func QuickSortStart[K comparable](arr []K, cmp func(K, K) int) {
	quickSort(arr, 0, len(arr)-1, cmp)
}

func quickSort[K comparable](arr []K, l, r int, cmp func(K, K) int) {
	if len(arr) < 2 {
		return
	}
	if l < r {
		part := partition(arr, l, r, cmp)
		//if part == -1 {
		//	return
		//}
		quickSort(arr, l, part, cmp)
		quickSort(arr, part+1, r, cmp)
	}
}

// -1 if left is greater, 0 if equal, 1 if right is greater
func partition[K comparable](arr []K, l, r int, cmp func(K, K) int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return -1
	}

	mid := (l + r) / 2

	for {
		for cmp(arr[l], arr[mid]) == 1 {
			l++
		}
		for cmp(arr[r], arr[mid]) == -1 {
			r--
		}
		if l >= r {
			return r
		}
		arr[l+1], arr[r-1] = arr[r-1], arr[l+1]
		l++
		r--
	}
}
