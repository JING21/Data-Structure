package LeetCode350_IntersectionofTwoArrays_II

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	m := map[int]int{}

	for _, v := range nums1 {
		m[v]++
	}

	result := []int{}

	for _, v := range nums2 {
		if m[v] > 0 {
			result = append(result, v)
			m[v]--
		}
	}

	return result
}

func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	len1, len2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	result := []int{}
	for index1 < len1 && index2 < len2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			result = append(result, nums1[index1])
			index1++
			index2++
		}
	}
	return result
}
