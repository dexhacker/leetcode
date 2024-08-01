package number_of_senior_citizens

func countSeniors(details []string) int {
	answer := 0
	for _, detailsItem := range details {
		decade := detailsItem[11]
		year := detailsItem[12]
		if decade > '6' {
			answer++
		}
		if decade == '6' && year > '0' {
			answer++
		}
	}
	return answer
}
