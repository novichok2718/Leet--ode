package main

import "slices"
 
func pathSum(root *TreeNode, targetSum int) [][]int {
	sliceOfSum := [][]int{}
	if root == nil {
		return sliceOfSum 
	}
	numbers := []int{}
	var sum func(root *TreeNode)
	sum = func(root *TreeNode) {
		numbers = append(numbers, root.Val)
		if root.Left == nil && root.Right == nil {
			ans := 0
			for _, val := range numbers {
				ans += val
			}
			if ans == targetSum {
				sliceOfSum = append(sliceOfSum, slices.Clone(numbers))
			}
			numbers = numbers[:len(numbers) - 1]
			return
		}
		if root.Left != nil { sum(root.Left)}
		if root.Right != nil {	sum(root.Right)}
		numbers = numbers[: len(numbers) - 1]
	}
	sum(root)
	return sliceOfSum
}
