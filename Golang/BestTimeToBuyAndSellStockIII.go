package main

import "fmt"

func profit(prices []int) int {
	size := len(prices)
	max := 0
	for i := 0; i != size; i++ {
		for j := i + 1; j != size; j++ {
			if max < (prices[j] - prices[i]) {
				max = prices[j] - prices[i]
			} 
		}
	}
	return max
}

func maxProfit(prices []int) int {
	leftMax := 0
	rightMax := 0
	size := len(prices)
	for i := 1; i != size; i++ {
		lm := profit(prices[:i])
		rm := profit(prices[i:])
		if (leftMax + rightMax) < (lm + rm) {
			leftMax = lm
			rightMax = rm
		}
	}
	return leftMax + rightMax
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
