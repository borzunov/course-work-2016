package words

import (
	"strconv"
	"strings"
)

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

func BitsFromRepr(repr string) (uint64, int) {
	result, err := strconv.ParseUint(repr, 2, 64)
	if err != nil {
		panic(err)
	}
	return result, len(repr)
}

func BitsToRepr(s uint64, n int) string {
	if n == 0 {
		return ""
	}

	repr := strconv.FormatUint(s, 2)
	return strings.Repeat("0", n-len(repr)) + repr
}
