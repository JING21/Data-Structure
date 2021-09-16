/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
 

Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
Example 4:

Input: s = "([)]"
Output: false
Example 5:

Input: s = "{[]}"
Output: true
 

Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
 */

package Valid_Parentheses_LeetCode20

type Stack []interface {}

func (s *Stack) IsEmpty() bool{
	if s.Length() == 0{
		return true
	}else {
		return false
	}
}

func (s Stack) Length() int{
	return len(s)
}

func (s *Stack) Cap() int {
	return s.Cap()
}


func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)

}

func (s Stack) Top() interface{}{
	if s.Length() == 0{
		return nil
	}
	topIndex := s.Length()
	return s[topIndex-1]
}

func (s *Stack) Pop() interface{}{
	newStack := *s
	if s.Length() == 0{
		return nil
	}
	value := newStack[newStack.Length()-1]
	*s = newStack[:newStack.Length()-1]
	return value
}


func isValid(s string) bool {
	newStack := Stack{}
	Hash := map[byte]byte{'}':'{',  ')':'(', ']':'['}
	if len(s) % 2 ==1 {
		return false
	}
	for i := 0; i<len(s); i++ {
		top := newStack.Top()
		value, _ := top.(byte)
		if s[i] == '{' || s[i] == '[' || s[i] == '('{
			newStack.Push(s[i])
		}else if value == Hash[s[i]] {
			newStack.Pop()
		}else{
			return false
		}
	}

	if newStack.IsEmpty() {
		return true
	}
	return false
}



