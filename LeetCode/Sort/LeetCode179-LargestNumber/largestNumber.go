package LeetCode179_LargestNumber

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	tmp := make([]string, len(nums))
	for k, v := range nums {
		tmp[k] = strconv.Itoa(v)
	}

	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i][0] == tmp[j][0] {
			return tmp[i]+tmp[j] > tmp[j]+tmp[i]
		}
		return tmp[i] > tmp[j]
	})

	result := strings.Join(tmp, "")
	if result[0] == '0' {
		return "0"
	}

	return result
}
