package maximum_distance_in_arrays

type Item struct {
	Number int
	Index  int
}

func maxDistance(arrays [][]int) int {
	var min1, min2, max1, max2 *Item
	for i := 0; i < len(arrays); i++ {
		first := arrays[i][0]
		last := arrays[i][len(arrays[i])-1]
		if min1 == nil || min1.Number >= first {
			min2 = min1
			min1 = &Item{Number: first, Index: i}
		} else if min2 == nil || min2.Number >= first {
			min2 = &Item{Number: first, Index: i}
		}
		if max1 == nil || max1.Number <= last {
			max2 = max1
			max1 = &Item{Number: last, Index: i}
		} else if max2 == nil || max2.Number <= first {
			max2 = &Item{Number: last, Index: i}
		}
	}

	if min1.Index != max1.Index {
		return max1.Number - min1.Number
	}

	answer := 0
	if max1.Index != min2.Index {
		answer = max1.Number - min2.Number
	}
	if max2.Index != min1.Index {
		if max2.Number-min1.Number > answer {
			return max2.Number - min1.Number
		}
	}
	return answer
}
