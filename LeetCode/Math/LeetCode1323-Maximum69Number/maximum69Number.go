package LeetCode1323_Maximum69Number

import "strconv"

func maximum69Number(num int) int {
	s := []byte(strconv.Itoa(num))
	for i := range s {
		if s[i] == '6' {
			s[i] = '9'
			break
		}
	}
	num, _ = strconv.Atoi(string(s))
	return num
}
