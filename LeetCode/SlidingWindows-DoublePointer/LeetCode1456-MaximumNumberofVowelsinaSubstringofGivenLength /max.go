package LeetCode1456_MaximumNumberofVowelsinaSubstringofGivenLength_

//定长滑动窗口分3个步骤
//1.入: 下标为i的元素进入窗口，更新统计，如果窗口的左端点i < k-1,即尚未形成长度为k的窗口，则继续第一步，延长窗口
//2.更新: 进行计算，计算窗口中的值，一般为最大值或者最小值
//3.出：下标为i-k+1的元素离开窗口，更新统计，为下一个循环做准备

func maxVowels(s string, k int) int {
	count := 0
	result := 0

	for i, ss := range s {
		//1.进入窗口
		if ss == 'a' || ss == 'o' || ss == 'e' || ss == 'i' || ss == 'u' {
			count++
		}

		//扩充直到满足窗口长度
		if i < k-1 {
			continue
		}

		//2.更新
		result = max(result, count)

		//3.出窗口
		out := s[i-k+1]
		if out == 'a' || out == 'e' || out == 'i' || out == 'o' || out == 'u' {
			count--
		}
	}
	return result
}
