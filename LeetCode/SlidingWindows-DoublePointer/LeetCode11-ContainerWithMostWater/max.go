package LeetCode11_ContainerWithMostWater

func maxArea(height []int) int {
	//可容纳水的最大面积，是由S(i,j)其中较短的短板决定的
	//S(i,j) = min(h[i]*h[j])* (j-i)
	i, j, result := 0, len(height)-1, 0

	for i < j {
		if height[i] < height[j] {
			result = max(result, height[i]*(j-i))
			i++
		} else {
			result = max(result, height[j]*(j-i))
			j--
		}
	}
	return result
}
