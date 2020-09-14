package main

import (
	"fmt"
	"github.com/golang/glog"
)

func main() {
	a:= isPalindrome(1221)
	fmt.Println(a)
}


func isPalindrome(x int) bool {
	if x < 0  {
		glog.Error(x)
		return false
	}

	if x % 10 == 0 && x != 0 {
		glog.Error(x)
		return false
	} else {
		revertNum := 0
		for x > revertNum {
			revertNum = revertNum * 10 + x % 10
			x /= 10
			glog.Error(revertNum)
		}
		return x == revertNum || x == revertNum / 10
	}

}