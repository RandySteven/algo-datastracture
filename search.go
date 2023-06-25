package main

func linearSearch(arr []int, target int) (bool, int) {
	index := -1
	search := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			index = i
			search = true
		}
	}
	return search, index
}

func binarySearch(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high {
		median := (low + high) / 2

		if arr[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(arr) || arr[low] != target {
		return false
	}

	return true
}

func interpolationSearch(arr []int, target int) int {
	min, max := arr[0], arr[len(arr)-1]
	low, high := 0, len(arr)-1
	for {
		if target < min {
			return low
		}

		if target > max {
			return high + 1
		}

		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size-1) * (float64(target-min) / float64(max-min)))
			guess = low + offset
		}

		if arr[guess] == target {
			for guess > 0 && arr[guess-1] == target {
				guess--
			}
			return guess
		}

		if arr[guess] > target {
			high = guess - 1
			max = arr[high]
		} else {
			low = guess + 1
			min = arr[low]
		}
	}
}
