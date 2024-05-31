package main

import "fmt"

func subsets(nums []int) [][]int {
	_nums := make([][]int, 2)
	for i := 0; i < 2; i++ {
		_nums[i] = make([]int, len(nums))
	}
	_nums[0] = nums

	allSubsets := [][]int{}
	var AllSubsets func(nums [][]int, n int)
	AllSubsets = func(nums [][]int, n int) {
		size := len(nums[0])
		if n == size {
			var sets []int
			for i := 0; i < size; i++ {
				if nums[1][i] == 0 {
					sets = append(sets, nums[0][i])		
				}
			}
			allSubsets = append(allSubsets, sets)
			return
		}
		nums[1][n] = 0;
		AllSubsets(nums, n + 1)
		nums[1][n] = 1
		AllSubsets(nums, n + 1)
	}
	AllSubsets(_nums, 0)
	return allSubsets
}

func main() {
	fmt.Print(subsets([]int{1, 2, 3}))
}