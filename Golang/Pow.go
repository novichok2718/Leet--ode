package main

func abs(x int) int {
    if (x < 0) {
        return -x
    }
    return x
}

func myPow(x float64, n int) float64 {
	_n := abs(n)
	product := 1.0
	for i := 0; i < 32; i++ {
		if 1 << i & _n != 0 {
			product = product * x
		}
		x = x * x
	}
	if n < 0 {
		return 1 / product 
	}
	return product
}