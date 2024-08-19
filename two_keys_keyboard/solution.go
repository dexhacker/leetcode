package two_keys_keyboard

func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			number := n / i
			if number == 1 {
				return i
			}
			return i + minSteps(number)
		}
	}
	return n
}
