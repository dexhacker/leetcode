package lemonade_change

func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			fives++
		} else if bill == 10 {
			fives--
			if fives < 0 {
				return false
			}
			tens++
		} else {
			fives--
			if tens == 0 {
				fives -= 2
			} else {
				tens--
			}

			if fives < 0 || tens < 0 {
				return false
			}
		}
	}
	return true
}
