package words

func CountDifferentLyndonWords(s []byte) int {
	n := len(s)
	i := 0
	result := 0
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

		//for i <= j {
		//	result = append(result, s[i:i+k-j])
		//	i += k - j
		//}
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
