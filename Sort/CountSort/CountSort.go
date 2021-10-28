package CountSort

//计数排序只能用在数据范围不大的场景中，如果数据范围 k 比要排序的数据 n 大很多，就不适合用计数排序了。
//而且，计数排序只能给非负整数排序，如果要排序的数据是其他类型的，要将其在不改变相对大小的情况下，转化为非负整数。



func CountingSort(a []int, n int){
	if n <= 1{
		return
	}

	max := a[0]

	for _, v := range a {
		if max < v{
			max = v
		}
	}

	c := make([]int, max+1)

	for k, _ := range c {
		c[k] = 0
	}

	for i, _ := range a {
		c[a[i]]++
	}

	for i := 1; i <= max; i++{
		c[i] = c[i-1]+c[i]
	}

	r := make([]int, n)

	for i := n-1; i>= 0; i--{
		index := c[a[i]] -1
		r[index] = a[i]
		c[a[i]]--
	}

	for i := 0; i < n; i++{
		a[i] = r[i];
	}
}