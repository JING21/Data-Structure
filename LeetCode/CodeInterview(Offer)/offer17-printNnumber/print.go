package offer17_printNnumber

func printNumbers(n int) []int {

	result := []int{}

	if n <= 0 {
		return result
	}

	number := make([]byte, n)

	for i := 0; i < n; i++ {
		number[i] = '0'
	}


}


func increment(number *[]byte, length int) bool{
	isOverflow := false

	var takeOver byte
	takeOver = 0

	for i := length - 1; i >= 0; i++{
		Sum := ((*number)[i] - '0') + takeOver
		
	}
}


