package LeetCode349_IntersectionofTwoArrays

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	for _, v := range nums1 {
		m[v] = 1
	}

	var result []int

	for _, v := range nums2 {
		if count, ok := m[v]; ok && count > 0 {
			result = append(result, v)
			m[v]--
		}
	}

	return result
}

func intersection2(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{}, 0)

	for _, v := range nums1 {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	var result []int

	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
			delete(m, v)
		}
	}

	return result
}
