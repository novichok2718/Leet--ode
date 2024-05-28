package main

import (
	"fmt"
)

func subsets(nums [][]int, n int) {
	size := len(nums) - 1
	if n == size {
		for i := 0; i <= size; i++ {
			if nums[1][i] == 0 {
				fmt.Printf("%d ", nums[0][i])		
			}
		}
		fmt.Printf("\n")
		return
	}
	nums[1][n] = 0;
	subsets(nums, n + 1)
	nums[1][n] = 1
	subsets(nums, n + 1)
}