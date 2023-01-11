package LeetCode1047_RemoveAllAdjacentDuplicatesInString

func removeDuplicates(s string) string {
	slow, fast := 0, 0

	stack := []byte(s)

	for fast < len(s) {
		stack[slow] = stack[fast]

		if slow > 0 && stack[slow] == stack[slow-1] {
			slow--
		} else {
			slow++
		}

		fast++
	}

	return string(stack)
}
