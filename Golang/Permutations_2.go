package main

var uniquePermutations [][]int

func swap(nums []int, lhs int, rhs int) {
	tmp := nums[lhs]
	nums[lhs] = nums[rhs]
	nums[rhs] = tmp
}


func permuteUniques (nums []int, n int) {
	if n == len(nums) - 1 {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		uniquePermutations = append(uniquePermutations, tmp)		
		return
	}
	for i := n; i < len(nums); i++{
		if nums[i] == nums[n] && i != n{
			continue
		}
		swap(nums, i, n)
		permuteUniques(nums, n + 1)
		swap(nums, n, i)
	} 
}

func permuteUnique(nums []int) [][]int {
	permuteUniques(nums, 0)
	return uniquePermutations
}
