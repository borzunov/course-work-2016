package words

func CountDifferentLyndonWords(s uint64, n int) int {
	i := 0
	result := 0
	begin := uint64(1) << uint(n-1)
	var j, k int
	for i < n {
		j = i
		k = i + 1
		for k < n {
			s_j := s & (begin >> uint(j)) != 0
			s_k := s & (begin >> uint(k)) != 0
			if s_j && !s_k {
				break
			}
			if !s_j && s_k {
				j = i
			} else {
				j++
			}
			k++
		}

		wLen := k - j
		i += ((j-i)/wLen + 1) * wLen
		result++
	}
	return result
}

func Lyndon(s []byte) [][]byte {
	n := len(s)
	i := 0
	result := make([][]byte, 0)
	var j, k int
	for i < n {
		j = i
		k = i + 1
		for k < n && s[j] <= s[k] {
			if s[j] < s[k] {
				j = i
			} else {
				j++
			}
			k++
		}

		for i <= j {
			result = append(result, s[i:i+k-j])
			i += k - j
		}
	}
	return result
}
