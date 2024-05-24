package main

func search(nums []int, target int) int {
	for key, val := range nums {
		if val == target {
			return key
		}
	}
	return -1
}
