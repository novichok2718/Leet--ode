package main

func maxSubArray(nums []int) int {
	size := len(nums)
	sum := 0
	maxSum := 0
	for i := 0; i < size; i++ {
		sum += nums[i]
		if maxSum < sum {
			maxSum = sum
		} 
		sum = max(0, sum)
	}
	return maxSum
}
