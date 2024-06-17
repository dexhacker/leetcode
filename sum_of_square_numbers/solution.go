package sum_of_square_numbers

func judgeSquareSum(c int) bool {
	if c == 0 {
		return true
	}

	mp := make(map[int]bool)
	mp[0] = true
	for i := 1; i < c/i; i++ {
		if i*i > c {
			break
		}
		mp[i*i] = true
	}

	for i := 1; i <= c/i; i++ {
		if _, ok := mp[c-i*i]; ok {
			return true
		}
	}

	return false
}
