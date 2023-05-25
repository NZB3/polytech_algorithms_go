package search

func Search(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int, target int) int {
	var max, min int
	max = len(arr) - 1
	min = 0

	for min <= max {
		var mid_idx int = (max + min) / 2
		mid_val := arr[mid_idx]
		if mid_val == target {
			return mid_idx
		} else if mid_val < target {
			min = mid_idx + 1
		} else {
			max = mid_idx - 1
		}
	}

	return -1
}
