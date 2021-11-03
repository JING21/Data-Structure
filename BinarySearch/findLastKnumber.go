package BinarySearch

//查找最后一个值等于给定值的元素
func bsearch2(a []int, value int) int{

	low := 0
	high := len(a)-1

	for low <= high{
		mid := (high-low)/2 + low

		if a[mid] > value{
			high = mid - 1
		}else if a[mid] < value{
			low = mid + 1
		}else {
			if mid == len(a)-1 || a[mid+1] != value{
				return mid
			}else {
				high = mid -1
			}
		}
	}

	return -1
}
