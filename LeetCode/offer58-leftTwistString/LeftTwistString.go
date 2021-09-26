package offer58_leftTwistString



func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverse(b,0, n-1)
	reverse(b,n, len(b)-1)
	reverse(b,0, len(b)-1)
	return s[n:]+s[:n]
}


func reverse(b []byte, left, right int){
	for left < right{
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}