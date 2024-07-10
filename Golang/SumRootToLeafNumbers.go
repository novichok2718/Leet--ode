package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pow(a int, b int) int {
	ans := 1
	for i := 0; i != b; i++ {
		ans *= a
	}
	return ans
}

func sumNumbers(root *TreeNode) int {
	digits := make([]int, 0)
	ans := 0
	var sum func(root *TreeNode)
	sum = func(root *TreeNode) {
		digits = append(digits, root.Val)
		if root.Left == nil && root.Right == nil {
			size := len(digits)
			for i := 0; i != size; i++ {
				ans += digits[i] * pow(10, size - i - 1) 
			}
			digits = digits[:len(digits) - 1]
			return
		}
		if root.Left != nil {	sum(root.Left)}
		if root.Right != nil {	sum(root.Right)}
		digits = digits[:len(digits) - 1]
	}
	sum(root)
	return ans
}
