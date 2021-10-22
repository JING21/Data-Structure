package offer66_constructArr


func constructArr(a []int) []int {
	length := len(a)

	if length == 0{
		return nil
	}

	b := make([]int, length)
	b[0] = 1
	tmp := 1
	for i := 1; i < length; i++{
		b[i] = b[i-1] * a[i-1]
	}

	for i := length-2; i >= 0; i--{
          tmp *= a[i+1]
          b[i] *= tmp
	}

	return b
}