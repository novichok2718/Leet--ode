package main

func max(a int, b int) int {
	if (a > b) {
		return a
	}
	return b
}

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

func main() {
	maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
}