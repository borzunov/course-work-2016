package main

import (
	"../words"
	"fmt"
)

func main() {
	var repr string
	fmt.Scan(&repr)
	s := words.FromRepr(repr)

	if len(s) < 1000 {
		fmt.Println("toRepr:", words.ToRepr(s))
		fmt.Println("lz:    ", words.ToDecompositionRepr(words.Lz(s)))
		fmt.Println("lyndon:", words.ToDecompositionRepr(words.Lyndon(s)))
		fmt.Println()
	}

	fmt.Println("countDLyndonWords:", words.CountDifferentLyndonWords(s))
	bufs := new(words.LzBuffers)
	bufs.Make(len(s))
	fmt.Println("countLz:          ", words.CountLz(s, bufs))
}
