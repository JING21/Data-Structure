package LeetCode461_HammingDistance

// 具体地，记 s=x⊕y，我们可以不断地检查 s 的最低位，如果最低位为 1，那么令计数器加一，然后我们令 s 整体右移一位，
// 这样 s 的最低位将被舍去，原本的次低位就变成了新的最低位。我们重复这个过程直到 s=0 为止。这样计数器中就累计了 s 的二进制表示中 1 的数量。
func hammingDistance2(x, y int) (ans int) {
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return
}

// Brian Kernighan 算法
// 记 f(x) 表示 x 和 x−1 进行与运算所得的结果（即 f(x)=x&(x−1)），那么 f(x) 恰为 x 删去其二进制表示中最右侧的 1 的结果。
func hammingDistance(x int, y int) int {
	result := 0
	for s := x ^ y; s > 0; s &= s - 1 {
		result++
	}
	return result
}
