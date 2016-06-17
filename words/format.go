package words

import "strings"

func FromRepr(repr string) []byte {
	result := make([]byte, len(repr))
	for i, ch := range repr {
		if ch != '0' && ch != '1' {
			panic("Invalid character")
		}
		result[i] = byte(ch - '0')
	}
	return result
}

func ToRepr(s []byte) string {
	result := make([]rune, len(s))
	for i, ch := range s {
		result[i] = rune(ch + '0')
	}
	return string(result)
}

func ToDecompositionRepr(words [][]byte) string {
	result := make([]string, len(words))
	for i, word := range words {
		result[i] = ToRepr(word)
	}
	return strings.Join(result, " ")
}
