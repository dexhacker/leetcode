package kth_largest_element_in_a_stream

import "sort"

type KthLargest struct {
	K    int
	Nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{K: k, Nums: nums}
}

func (kLargest *KthLargest) Add(val int) int {
	kLargest.Nums = append(kLargest.Nums, val)
	sort.Ints(kLargest.Nums)
	return kLargest.Nums[len(kLargest.Nums)-kLargest.K]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
