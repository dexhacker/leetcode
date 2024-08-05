package kth_distinct_string_in_an_array

func kthDistinct(arr []string, k int) string {
	mp := make(map[string]int)
	for _, item := range arr {
		mp[item]++
	}
	list := make([]string, 0)
	for _, item := range arr {
		if mp[item] == 1 {
			list = append(list, item)
		}
	}

	if len(list) < k {
		return ""
	}
	return list[k-1]
}
