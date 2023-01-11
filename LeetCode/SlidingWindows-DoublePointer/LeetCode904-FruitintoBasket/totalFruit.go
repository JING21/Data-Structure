package LeetCode904_FruitintoBasket

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

func totalFruit2(fruits []int) int {
	cnt := map[int]int{}
	j := 0
	for _, x := range fruits {
		cnt[x]++
		if len(cnt) > 2 {
			y := fruits[j]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			j++
		}
	}
	return len(fruits) - j
}
