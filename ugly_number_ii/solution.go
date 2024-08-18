package ugly_number_ii

func minimum2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimum3(a, b, c int) int {
	return minimum2(minimum2(a, b), c)
}

func nthUglyNumber(n int) int {
	ugly := make([]int, n)
	ugly[0] = 1
	index2, index3, index5 := 0, 0, 0
	factor2, factor3, factor5 := 2, 3, 5
	for i := 1; i < n; i++ {
		minimum := minimum3(factor2, factor3, factor5)
		ugly[i] = minimum
		if factor2 == minimum {
			index2++
			factor2 = 2 * ugly[index2]
		}
		if factor3 == minimum {
			index3++
			factor3 = 3 * ugly[index3]
		}
		if factor5 == minimum {
			index5++
			factor5 = 5 * ugly[index5]
		}
	}
	return ugly[n-1]
}
