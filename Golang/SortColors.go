package main

func sortColors(nums []int) {
	//size := len(nums)
	_nums := make([]int, 3)
	for i := 0; i < len(nums); i++ {
		_nums[nums[i]]++
	}

	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < _nums[i]; j++ {
			nums[count] = i
			count++
		} 
	}
}
