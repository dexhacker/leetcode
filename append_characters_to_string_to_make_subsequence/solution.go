package append_characters_to_string_to_make_subsequence

func appendCharacters(s string, t string) int {
	j := 0
	for i := range s {
		if s[i] == t[j] {
			j++
		}

		if j >= len(t) {
			return 0
		}
	}

	return len(t) - j
}
