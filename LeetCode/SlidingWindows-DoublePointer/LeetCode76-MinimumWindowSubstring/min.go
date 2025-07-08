package LeetCode76_MinimumWindowSubstring

import (
	"math"
)

func minWindow(s string, t string) string {
	//一个字符串t，一个字符串s
	//返回s中包含t的最小子串
	hs, ht := make(map[byte]int), make(map[byte]int)

	//ht表示字符串t中，每个字母出现的次数
	for i := range t {
		ht[t[i]]++
	}
	sLen := len(s)
	length := math.MaxInt32

	ansLeft, ansRight := -1, -1

	//对比如果滑动窗口中的字母数量少于所需字母，则不满足要求
	check := func() bool {
		for k, v := range ht {
			if hs[k] < v {
				return false
			}
		}
		return true
	}

	//开始进行滑动
	for left, right := 0, 0; right < sLen; right++ {
		//将最右边的字母验证，t中是否需要
		if right < sLen && ht[s[right]] > 0 {
			hs[s[right]]++
		}
		//满足hs[]数量大于t需要数量
		for check() && left <= right {
			//比较最小长度
			if right-left+1 < length {
				length = right - left + 1
				ansLeft, ansRight = left, left+length
			}
			if _, ok := ht[s[left]]; ok {
				hs[s[left]] -= 1
			}
			left++
		}
	}
	//表示不存在
	if ansLeft == -1 {
		return ""
	}
	return s[ansLeft:ansRight]
}
