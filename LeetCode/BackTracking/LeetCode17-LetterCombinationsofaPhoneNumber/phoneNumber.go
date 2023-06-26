package LeetCode17_LetterCombinationsofaPhoneNumber

var (
	m      []string
	path   []byte
	result []string
)

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, result = make([]byte, 0), make([]string, 0)
	if digits == "" {
		return result
	}

	backTracing(digits, 0)

	return result
}

func backTracing(digits string, start int) {
	if len(path) == len(digits) {
		tmp := string(path)
		result = append(result, tmp)
		return
	}

	digit := int(digits[start] - '0')
	str := m[digit-2]

	for j := 0; j < len(str); j++ {
		path = append(path, str[j])
		backTracing(digits, start+1)
		path = path[:len(path)-1]
	}
}
