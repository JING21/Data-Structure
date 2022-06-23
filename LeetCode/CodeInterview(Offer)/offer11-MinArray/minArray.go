package offer11_MinArray


func minArray(numbers []int) int{
		low := 0
		high := len(numbers)-1
		for low < high{
			mid := (high-low) / 2 + low
			if numbers[mid] < numbers[high]{
					high = mid
			}else if numbers[mid] > numbers[high]{
					 low = mid+1
			}else {
				high--
			}
		}
		return numbers[low]
}