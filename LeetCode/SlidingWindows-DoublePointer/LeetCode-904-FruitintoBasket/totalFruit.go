package LeetCode_904_FruitintoBasket

func totalFruit(fruits []int) int {
	left, maxLen, n := 0, 0, len(fruits)

	if len(fruits) == 0 {
		return 0
	}

	fruitMap := map[int]int{}

	for i := 0; i < n; i++ {
		fruitMap[fruits[i]]++
		for left <= i && len(fruitMap) > 2 {
			fruitMap[fruits[i]]--
			if fruitMap[fruits[left]] == 0 {
				delete(fruitMap, fruitMap[fruits[left]])
			}
			left++
		}
		if maxLen < i-left+1 {
			maxLen = i - left + 1
		}
	}
	return maxLen
}
