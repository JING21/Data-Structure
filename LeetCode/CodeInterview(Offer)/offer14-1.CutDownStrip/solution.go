package offer14_1_CutDownStrip

import "math"

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	quotient := n / 3
	remainder := n % 3
	if remainder == 0 {
		return int(math.Pow(3, float64(quotient)))
	} else if remainder == 1 {
		return int(math.Pow(3, float64(quotient - 1))) * 4
	}
	return int(math.Pow(3, float64(quotient))) * 2
}

