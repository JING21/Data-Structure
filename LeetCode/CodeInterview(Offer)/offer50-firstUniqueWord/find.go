//package offer50_firstUniqueWord
package main
import "fmt"

func firstUniqChar(s string) byte {
	stringMap := make(map[uint8]int, len(s))
	for i := 0; i <len(s); i++{
		if _, ok := stringMap[s[i]]; ok{
			stringMap[s[i]] = -1
		}else {
			stringMap[s[i]] = i
		}
	}

	min := len(s)-1
	var rs []int
	count := 0
	for _, v := range stringMap {
		if v == -1{
			count++
		}else {
			min = findMin(min, v)
			rs = append(rs, min)
		}
	}

	if  len(rs) == 0 {
		return ' '
	}
	return s[min]
}

func findMin(a, b int) int{
	if a > b{
		return b
	}
	return a
}

func main() {
	a := firstUniqChar2("leetcode")
	fmt.Println(string(a))
}


func firstUniqChar2(s string) byte {
	var list [26]int
	length := len(s)
	for i:=0;i<length;i++ {
		list[s[i]-'a']++
	}
	for i:=0;i<length;i++{
		if list[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}