package LeetCode844_BackspaceStringCompare

func backspaceCompare(s, t string) bool {
	deleteS, deleteT := 0, 0
	i, j := len(s)-1, len(t)-1

	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				deleteS++
				i--
			} else if deleteS > 0 {
				deleteS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				deleteT++
				j--
			} else if deleteT > 0 {
				deleteT--
				j--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}

	return true
}
