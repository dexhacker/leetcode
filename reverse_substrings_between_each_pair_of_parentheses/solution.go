package reverse_substrings_between_each_pair_of_parentheses

func reverse(l int, r int, s []rune) []rune {
	for l < r {
		tmp := s[l]
		s[l] = s[r]
		s[r] = tmp
		l++
		r--
	}
	return s
}

func reverseParentheses(s string) string {
	runes := []rune(s)
	brackets := make([]int, 0)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' {
			brackets = append(brackets, i)
		}
		if runes[i] == ')' {
			last := len(brackets) - 1
			runes = reverse(brackets[last]+1, i-1, runes)
			runes[brackets[last]] = '_'
			runes[i] = '_'
			brackets = brackets[:len(brackets)-1]
		}
	}

	answer := ""
	for _, item := range runes {
		if item != '_' {
			answer += string(item)
		}
	}
	return answer
}
