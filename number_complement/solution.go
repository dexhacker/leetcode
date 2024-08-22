package number_complement

func findComplement(num int) int {
	i := 1
	for {
		if i >= num {
			return num ^ i
		}
		i = (i << 1) + 1
	}
}
