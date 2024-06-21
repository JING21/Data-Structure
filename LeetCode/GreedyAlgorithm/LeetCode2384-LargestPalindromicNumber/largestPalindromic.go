package LeetCode2384_LargestPalindromicNumber

func largestPalindromic(num string) string {
	count := ['9' + 1]int{}

	for _, v := range num {
		count[v]++
	}

	//特殊情况，num都是0
	if count['0'] == len(num) {
		return "0"
	}

	s := []byte{}

	for
}
