package main

import (
	"../words"
	"fmt"
)

func main() {
	var repr string
	fmt.Scan(&repr)
	s := words.FromRepr(repr)
	bits, n := words.BitsFromRepr(repr)

	if n <= 1000 {
		fmt.Println("toRepr:    ", words.ToRepr(s))
		fmt.Println("bitsToRepr:", words.BitsToRepr(bits, n))
		fmt.Println()

		fmt.Println("lz:    ", words.ToDecompositionRepr(words.Lz(s)))
		fmt.Println("lyndon:", words.ToDecompositionRepr(words.Lyndon(s)))
		fmt.Println()

		bufs := new(words.SufSortBuffers)
		bufs.Make(n)
		fmt.Println("qsufsort:  ", words.QSufSort(s, bufs))
		fmt.Println("bitsufsort:", words.BitSufSort(bits, n, bufs))
		fmt.Println()
	}

	fmt.Println("countDLyndonWords:", words.CountDifferentLyndonWords(bits, n))
	bufs := new(words.LzBuffers)
	bufs.Make(n)
	fmt.Println("countLz:          ", words.CountLz(bits, n, bufs))
}
