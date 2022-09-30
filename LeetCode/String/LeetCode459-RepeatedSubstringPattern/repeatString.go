package LeetCode459_RepeatedSubstringPattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	newString := s + s
	deleteFirst := newString[1:]
	deleteEnd := deleteFirst[:len(deleteFirst)-1]
	return strings.Contains(deleteEnd, s)
}

//KMP
