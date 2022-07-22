package LeetCode438_FindAllAnagramsinaString

func findAnagrams(s string, p string) []int {
	var ans []int
	m, n := len(s), len(p)

	if m < n {
		return ans
	}

	var cntsP, cnts [26]int

	for i := 0; i < n; i++ {
		cntsP[p[i]-byte('a')]++
	}

	for i := 0; i < m; i++ {
		cnts[s[i]-byte('a')]++
		if i >= n-1 {
			if cnts == cntsP {
				ans = append(ans, i-n+1)
			}
			cnts[s[i-n+1]-byte('a')]--
		}
	}
	return ans
}
