package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := n - 1

	for i := m; i < len(nums1); i++ {
		nums1[i] = nums2[j]
		j--
	}
	sort.Ints(nums1)
}
