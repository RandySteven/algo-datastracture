package main

func twoSums(arr []int, target int) []int {

	var completeness int
	var numbMap map[int]int
	for i := 0; i < len(arr); i++ {
		completeness = target - arr[i]
		if val, ok := numbMap[completeness]; ok {
			return []int{val, i}
		} else {
			numbMap[arr[i]] = i
		}
	}
	return []int{-1, -1}
}
