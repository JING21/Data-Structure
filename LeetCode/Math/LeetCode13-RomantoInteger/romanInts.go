package LeetCode13_RomantoInteger

func romanToInt(s string) int {

	var getValue func(s string) int
	{
		switch s {
		case "I":
			return 1
		case "V":
			return 5
		case "X":
			return 10
		case "L":
			return 50
		case "C":
			return 100
		case "D":
			return 500
		case "M":
			return 1000
		}
	}

	sum := 0
	preNum := getValue(string(s[0]))

	for i := 1; i < len(s); i++ {
		num := getValue(string(s[i]))
		if preNum < num {
			sum -= preNum
		} else {
			sum += preNum
		}
		preNum = sum
	}

	sum += preNum

	return sum
}
