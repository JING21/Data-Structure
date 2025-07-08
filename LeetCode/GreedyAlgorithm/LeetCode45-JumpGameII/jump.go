package LeetCode45_JumpGameII

func jump(nums []int) int {
	//当前覆盖最大下标
	curRight := 0
	//下一步覆盖的最远距离下标
	nextRight := 0
	//记录走的步数
	ans := 0

	//i=n−2 时，如果 i<curRight，说明可以到达 n−1；如果 i=curRight，我们会造桥，这样也可以到达 n−1。
	//所以无论是何种情况，都只需要遍历到 n−2。或者说，n−1 已经是终点了
	for i, num := range nums[:len(nums)-1] {
		//遍历路程，记录下一步覆盖的最远距离下标
		nextRight = max(nextRight, i+num)
		//遇到当前覆盖的最远距离下标
		if i == curRight {
			//更新当前距离的最远下标
			curRight = nextRight
			ans++
		}
	}
	return ans
}
