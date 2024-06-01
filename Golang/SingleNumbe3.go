package main

func singleNumber(nums []int) []int {
	ans := []int{}
	count := 0
	for i := 0; i != len(nums); i++ {
		for j := 0; j != len(nums); j++ {
			if nums[i] ^ nums[j] == 0 {
				count++
			}
		}
		if count == 1 {
			ans = append(ans, nums[i])
		}
		count = 0
	}
	return ans
}
