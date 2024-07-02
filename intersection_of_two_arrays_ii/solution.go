package intersection_of_two_arrays_ii

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i1 := 0
	i2 := 0
	answer := make([]int, 0)
	for {
		if i1 >= len(nums1) {
			break
		}
		if i2 >= len(nums2) {
			break
		}
		if nums1[i1] == nums2[i2] {
			answer = append(answer, nums1[i1])
			i1++
			i2++
		} else if nums1[i1] > nums2[i2] {
			i2++
		} else {
			i1++
		}
	}
	return answer
}
