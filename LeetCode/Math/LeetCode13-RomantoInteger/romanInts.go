package LeetCode13_RomantoInteger

func romanToInt(s string) int {
	getValue := func(s byte) int {
		switch s {
		case 'I':
			return 1
		case 'V':
			return 5
		case 'X':
			return 10
		case 'L':
			return 50
		case 'C':
			return 100
		case 'D':
			return 500
		case 'M':
			return 1000
		default:
			return 0
		}
	}

	sum := 0

	for i := range s {
		num := getValue(s[i])
		if i < len(s)-1 && num < getValue(s[i+1]) {
			sum -= num
		} else {
			sum += num
		}
	}
	return sum
}
