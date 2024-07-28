package main

func maxProfitII(prices []int) int {
	size := len(prices)
	total := 0
	i := 0
	j := 0
	currMax := 0
	for ; i != size; i++ {
		currMax = prices[i]
		for j = i + 1;j != size; j++ {
			if prices[j] < currMax {
				total += (prices[j-1] - prices[i])
				i = j - 1
				break
			}
			currMax = prices[j]
		}
		if j == size {
			total += (prices[j-1] - prices[i])
			break
		}
	}
	return total
}
