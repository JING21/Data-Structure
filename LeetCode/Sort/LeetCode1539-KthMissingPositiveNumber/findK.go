package LeetCode1539_KthMissingPositiveNumber

// 确失的第k个整数，至少是K，遍历数组，遇到小于或等于K的，就向后移动，如果k=5,且数组包含1，那么k肯定至少是6
func findKthPositive(arr []int, k int) int {
	for i := range arr {
		if arr[i] <= k {
			k++
		}
	}
	return k
}

// 二分法，正常情况下arr[i]-i-1=0,第一位是arr[0]=1, 1-0-1=0, 缺失正整数的情况下：arr[i] - i - 1 得到arr[i]之前缺正整数的数量
func findKthPositive2(arr []int, k int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if arr[mid]-mid-1 < k {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return k + l
}
