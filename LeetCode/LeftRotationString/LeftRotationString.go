package main

import "fmt"

func main() {
	var s = "abcdefg"
	ls := leftRotationString(s,2)
	ls2 := leftRotationString2(s,2)
	fmt.Println(ls)
	fmt.Println(ls2)
}


func leftRotationString(s string, n int)string{
	return s[n:]+s[:n]
}


func leftRotationString2(s string, n int)string{
	var s1 string
	var s2 string
	for i:=0;i<len(s);i++{
		if i<n {
			s1 += string(s[i])
		}else
		{
			s2 += string(s[i])
		}
	}
	return s2+s1
}