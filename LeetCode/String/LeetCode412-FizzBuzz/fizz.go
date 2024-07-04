package LeetCode412_FizzBuzz

import (
	"strconv"
	"strings"
)

func fizzBuzz(n int) (result []string) {
	for i := 1; i <= n; i++ {
		sb := &strings.Builder{}
		if i%3 == 0 {
			sb.WriteString("Fizz")
		}
		if i%5 == 0 {
			sb.WriteString("Buzz")
		}
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		result = append(result, sb.String())
	}

	return result
}
