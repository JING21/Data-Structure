package LeetCode3110_ScoreOfaString

import "math"

func scoreOfString(s string) int {
	result := 0
	for i := 1; i < len(s); i++ {
		result += int(math.Abs(float64(s[i]) - float64(s[i-1])))
	}

	return result
}
