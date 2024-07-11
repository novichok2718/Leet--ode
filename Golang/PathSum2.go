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

func main() {
	//root := &TreeNode{Val: 5}
	//root.Left = &TreeNode{Val: 4}
	//root.Right = &TreeNode{Val: 8}
	//root.Left.Left = &TreeNode{Val: 11}
	//root.Left.Left.Left = &TreeNode{Val: 7}
	//root.Left.Left.Right = &TreeNode{Val: 2}
	//root.Right.Left = &TreeNode{Val: 13}
	//root.Right.Right = &TreeNode{Val: 4}
	//root.Right.Right.Left = &TreeNode{Val: 5}
	//root.Right.Right.Right = &TreeNode{Val: 1}
	pathSum(nil, 2)
}