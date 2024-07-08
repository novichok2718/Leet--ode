package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {	return 0 }
	if len(nums) == 2 {
		if nums[0] > nums[1] { return 0}
		return 1
	}


	start := 0
	end := len(nums)
	middle := 0
	for start < end {
		middle = (start + end) / 2
		
		if middle == 0 && nums[0] > nums[1] {	return 0 }
		if middle == len(nums) - 1 && nums[middle] > nums[middle - 1] {	return middle }

		if nums[middle] > nums[middle + 1] && nums[middle] > nums[middle - 1] {	return middle }

		if nums[middle] < nums[middle + 1] {
			start = middle + 1
		} else if nums[middle] < nums[middle - 1] {
			end = middle
		}
	}
	return middle
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 4, 3}))
}