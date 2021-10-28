package BinarySearch
//递归实现
func BinarySearch(a []int, val int) bool{

	return bSearch(a,0, len(a)-1, val)
}


func bSearch(a []int, low, high, value int) bool{
	if low > high{
		return false
	}

	mid := (high-low)/2 + low

	if a[mid] == value{
		return true
	}else if a[mid] < value{
		return bSearch(a, mid+1, high, value)
	}else {
		return bSearch(a, low, mid-1, value)
	}
}