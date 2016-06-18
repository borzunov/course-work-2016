package main

import (
	"../words"
	"fmt"
)

func main() {
	var repr string
	fmt.Scan(&repr)
	s := words.FromRepr(repr)
	n := len(s)

	if n <= 1000 {
		fmt.Println("toRepr:", words.ToRepr(s))
		fmt.Println("lz:    ", words.ToDecompositionRepr(words.Lz(s)))
		fmt.Println("lyndon:", words.ToDecompositionRepr(words.Lyndon(s)))
		fmt.Println()

		bufs := new(words.QSufSortBuffers)
		bufs.Make(n)
		fmt.Println("qsufsort:  ", words.QSufSort(s, bufs))
		fmt.Println("bitsufsort:", words.BitSufSort(s, bufs))
		fmt.Println()
	}

	fmt.Println("countDLyndonWords:", words.CountDifferentLyndonWords(s))
	bufs := new(words.LzBuffers)
	bufs.Make(n)
	fmt.Println("countLz:          ", words.CountLz(s, bufs))
}
