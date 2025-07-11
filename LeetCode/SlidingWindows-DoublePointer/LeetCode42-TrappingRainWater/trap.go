package LeetCode42_TrappingRainWater

// 以宽度来计算，按列计算，每列宽度为1，只需要计算该列i左侧最高柱子和该列右侧最高柱子中，两者的较小值减去自己列的高度
// （ min(MAXLEFTHeight, MAXRightHeight) - hegiht[i] ）* 宽度 = 该列能接水的体积
func trap(height []int) int {
	sum := 0
	n := len(height)

	//表示位置i左侧柱子中最高的高度
	lefth := make([]int, n)
	//表示位置i右侧柱子中最高的高度
	righth := make([]int, n)

	lefth[0] = height[0]
	righth[n-1] = height[n-1]

	//分别算出左右柱子的最高柱
	for i := 1; i < n; i++ {
		lefth[i] = max(lefth[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		righth[i] = max(righth[i+1], height[i])
	}

	for i := 1; i < n-1; i++ {
		h := newMin(lefth[i], righth[i]) - height[i]
		if h > 0 {
			sum += h
		}
	}
	return sum
}

func newMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//单调栈
//当遍历柱子时，当前元素的高度大于栈顶元素高度，就会出现凹槽，比如说遍历i对应高度是3，栈内是6，2，1
//那栈顶元素1弹出，对应的下标就是中间位置，记为mid，高度就是height[mid],即为1
//那右边的元素，即遍历的元素i，即为凹槽的右边，记为right，高度为height[right].即为3
//那左边的元素，即为目前栈顶的元素，即为凹槽的左边，记为left，高度为height[left].即为2
//那雨水高度即为，min(左侧凹槽高度，右侧凹槽高度) - 凹槽底部高度
//雨水的宽度即为，右边下标-左边下标-1，体积即为两者的乘积

func trap(height []int) int {
	stack := make([]int, 0)
	result := 0

	for i := 0; i < len(height); i++ {
		//满足栈不为空，并且当前元素高度大于栈顶元素高度
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			//获得凹槽的高度
			mid := height[stack[len(stack)-1]]
			//吐出栈顶元素
			stack = stack[:len(stack)-1]

			//如果此时栈顶不为空，则栈顶元素为左侧柱子坐标
			if len(stack) > 0 {
				//求得雨水高度
				h := newMin(height[i], height[stack[(len(stack)-1)]]) - mid
				//求得雨水宽度
				w := i - stack[len(stack)-1] - 1
				result += h * w
			}
		}
		//如果栈顶为空，或者当前柱子高度小于对应栈顶元素时，入栈
		stack = append(stack, i)
	}
	return result
}
