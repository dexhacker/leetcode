package integer_to_english_words

var mpNumbers = map[int]string{
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}

var mpTens = map[int]string{
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
}

var mpDecades = map[int]string{
	2: "Twenty",
	3: "Thirty",
	4: "Forty",
	5: "Fifty",
	6: "Sixty",
	7: "Seventy",
	8: "Eighty",
	9: "Ninety",
}

var thousands = []string{"Billion", "Million", "Thousand"}

func toNumber(num int) string {
	answer := ""
	hundreds := num / 100
	tens := (num % 100) / 10
	ones := num % 10
	if hundreds > 0 {
		answer += mpNumbers[hundreds] + " Hundred"
	}
	if tens > 0 {
		if hundreds > 0 {
			answer += " "
		}
		if tens == 1 {
			answer += mpTens[num%100]
			return answer
		} else {
			answer += mpDecades[tens]
		}
	}
	if ones > 0 {
		if hundreds > 0 || tens > 0 {
			answer += " "
		}
		answer += mpNumbers[ones]
	}
	return answer
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	answer := ""
	devider := 1000000000
	for i := 0; i < 4; i++ {
		if i == 0 {
			if num/devider > 0 {
				answer += toNumber(num/devider) + " " + thousands[i]
			}
		} else {
			currentThousand := (num % devider) / (devider / 1000)
			if currentThousand > 0 {
				if answer != "" && answer[len(answer)-1] != ' ' {
					answer += " "
				}
				if devider == 1000 {
					answer += toNumber(currentThousand)
				} else {
					answer += toNumber(currentThousand) + " " + thousands[i]
				}
			}
			devider /= 1000
		}
	}

	return answer
}
