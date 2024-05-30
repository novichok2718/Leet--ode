package main

func combine(n int, k int) [][]int {
	res := [][]int{}

	var backtrack func(start int, curSet []int)
	backtrack = func(start int, curSet []int) {
		size := len(curSet)
		if size == k {
			curSetCopy := make([]int, size)
			copy(curSetCopy, curSet)
			res = append(res, curSetCopy)
			return
		}
		for i := start; i <= n; i++ {
			curSet = append(curSet, i)
			backtrack(i + 1, curSet)
			curSet = curSet[: len(curSet) - 1]
		}
	}
	backtrack(1, []int{})
	
	return res
}

