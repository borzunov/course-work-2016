package main

import (
	"../enumerator"
	"../words"
	"fmt"
	"time"
)

type Buffers struct {
	words.LzBuffers
}

func NewBuffers(n int) enumerator.Context {
	result := new(Buffers)
	result.LzBuffers.Make(n)
	return result
}

func main() {
	for length := 1; ; length++ {
		fmt.Println("length =", length)
		start := time.Now()
		enumerator.EnumerateWords(length, 2, NewBuffers, func(s []byte, bufs enumerator.Context) {
			l := words.CountDifferentLyndonWords(s)
			z := words.CountLz(s, &bufs.(*Buffers).LzBuffers)
			if l > z {
				fmt.Println(words.ToRepr(s), l, z)
				if l-z > 1 {
					fmt.Println("\tWarning!")
				}
			}
		})

		seconds := time.Now().Sub(start).Seconds()
		if seconds > 10 {
			fmt.Printf("\tTime: %.1f s\n", seconds)
		}
	}
}
