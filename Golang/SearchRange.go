package main

func searchRange(nums []int, target int) []int {
	
	bin_search := func(nums []int, target int, isLeftSearch bool) int {
		left := 0
		right := len(nums) - 1
		idx := -1
		for left <= right {
			mid := (right + left) / 2
			if (nums[mid] > target) {
				right = mid - 1
			} else if (nums[mid] < target) {
				left = mid + 1
			} else {
				idx = mid
				if (isLeftSearch) {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
		return idx
	}

	return []int{bin_search(nums, target, true), bin_search(nums, target, false)}
}
