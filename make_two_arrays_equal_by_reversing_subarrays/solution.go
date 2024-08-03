package make_two_arrays_equal_by_reversing_subarrays

func canBeEqual(target []int, arr []int) bool {
	mp := make(map[int]int)
	for i := 0; i < len(target); i++ {
		if _, ok := mp[target[i]]; ok {
			mp[target[i]]++
		} else {
			mp[target[i]] = 1
		}

		if _, ok := mp[arr[i]]; ok {
			mp[arr[i]]--
		} else {
			mp[arr[i]] = -1
		}
	}

	for _, v := range mp {
		if v != 0 {
			return false
		}
	}

	return true
}
