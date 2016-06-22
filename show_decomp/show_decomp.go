package main

import (
	"../words"
	"bytes"
	"fmt"
)

func countWithoutDuplicates(words [][]byte) int {
	result := 0
	for i, word := range words {
		if i == 0 || !bytes.Equal(words[i-1], word) {
			result++
		}
	}
	return result
}

func main() {
	var repr string
	fmt.Scan(&repr)

	s := words.FromRepr(repr)
	n := len(s)
	lz := words.Lz(s)
	lyndon := words.Lyndon(s)
	if n <= 1000 {
		fmt.Println("ToRepr:", words.ToRepr(s))
		fmt.Println("Lyndon:", words.ToDecompositionRepr(lyndon))
		fmt.Println("Lz:    ", words.ToDecompositionRepr(lz))
		fmt.Println()
	}
	fmt.Println("countWD(lyndon):", countWithoutDuplicates(lyndon))
	fmt.Println("len(lz):        ", len(lz))

	if n <= 63 {
		bits, n := words.BitsFromRepr(repr)
		fmt.Println()

		fmt.Println("CountDLyndonWords:", words.CountDifferentLyndonWords(bits, n))

		bufs := new(words.LzBuffers)
		bufs.Make(n)
		fmt.Println("CountLz:          ", words.CountLz(bits, n, bufs))
	}
}
