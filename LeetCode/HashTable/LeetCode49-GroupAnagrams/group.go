package LeetCode49_GroupAnagrams

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	//声明一个对应的哈希表，key为排序完后的字符串，value为是该字符串的对应字符串数组，数组中的元素应该是排序前的字符串
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		//排序对应的字符串，相同的字符组成的字符串，排序后应该相同
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		//排序完后的字符串
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}

	ans := make([][]string, 0, len(mp))

	for _, v := range mp {
		ans = append(ans, v)
	}

	return ans
}

func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{}

	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}

	ans := make([][]string, 0, len(mp))

	for _, v := range mp {

		ans = append(ans, v)
	}

	return ans
}
