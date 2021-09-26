package offer05_replaceBlank


func replaceSpace (s string) string{
	b := []byte(s)
	length := len(b)
	count := 0

	for _,v := range b{
		if v == ' '{
			count++
		}
	}
	newCount := count*2
	tmp := make([]byte, newCount)
	b = append(b, tmp...)

	i := length-1
	j := len(b)-1

	for i >=0 {
		if b[i] != ' '{
			b[j] = b[i]
			i--
			j--
		}else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j-3
		}

	}
	return string(b)
}