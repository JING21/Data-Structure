package LeetCode560_SubarraySumEqualsK

// 定义前缀和为pre[i],表示以i结尾的[0..i]里所有数的和
// pre[i] = pre[i-1] + nums[i]
// 那么[j..i]这个子数组的和为k就可以转换成pre[i] - pre[j-1] = k
// 就可以得到 pre[j-1] = pre[i] - k
// map[pre[i]-k]记录pre[j-1]的值，有多少个前缀和为pre[i]-k的pre[j]即可
// 某子数组的和为K，表示pre

func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	preSum := 0

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if _, ok := hash[preSum-k]; ok {
			count += hash[preSum-k]
		}
		hash[preSum-k]++
	}
	return count
}
