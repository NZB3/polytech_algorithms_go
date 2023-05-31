package sorts

func BubleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func InsertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		j := i
		for j > 0 && arr[j-1] > curr {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = curr
	}
	return arr
}

//func UnNamedSort(arr []int) []int {
//	min_idx := 0
//	for i := 0; i < len(arr); i++ {
//		if arr[i] < arr[min_idx] {
//			arr[i], arr[min_idx] = arr[min_idx], arr[i]
//			min_idx = i
//		}
//	}
//	return arr
//}

func SelectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		var minIdx = i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[int(len(arr)/2)]
	left := make([]int, 0)
	right := make([]int, 0)
	equal := make([]int, 0)
	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			equal = append(equal, num)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, equal...), right...)
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
