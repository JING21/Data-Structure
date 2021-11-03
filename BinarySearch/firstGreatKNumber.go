package BinarySearch

//查找第一个大于等于给定值的元素
func bsearch3(a []int, value int) int{

	low := 0
	high := len(a)-1

	for low <= high{
		mid := (high-low)/2 + low

		if a[mid] >= value{
			if mid == 0 || a[mid-1] < value{
				return mid
			}else {
				high = mid - 1
			}
		}else {
			low = mid +1
			}
		}
	return -1
	}
