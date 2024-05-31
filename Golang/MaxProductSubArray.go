package main

func getNumberOfNegativeNumbers(start int, nums []int) int {
	size := len(nums)
	count := 0
	for i := start; i < size; i ++ {
		if (nums[i] < 0) {
			count++
		}
	} 
	return count
}

func maxProduct(nums []int) int {
	maxProd := nums[0]
	prod := 1
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] < 0 {
			num := getNumberOfNegativeNumbers(i + 1, nums)
			if prod > 0 && num == 0 {
				prod = 1
				continue
			}
		}
		prod *= nums[i]
		maxProd = max(maxProd, prod)
	}
	return maxProd
}
