package BinarySearch

//时间复杂度为O(logn),而且需要是有序数组，二分查找只能用在插入、删除操作不频繁，一次排序多次查找的场景中，数据量太小不适合二分查找，数据量太大也不适合二分查找

func binarySearch(a []int, value int) bool{

	low := 0
	high := len(a)-1

	for low <= high{
		mid := (high-low)/2 + low

		if a[mid] == value{
			return true
		}else if a[mid] < value{
			low = mid+1
		}else {
			high = mid -1
		}
	}

	return false
}


