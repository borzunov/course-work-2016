package main

import (
	"../enumerator"
	"../words"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"time"
)

type Buffers struct {
	s []byte
	words.LzBuffers
}

func NewBuffers(n int) enumerator.Context {
	result := new(Buffers)
	result.s = make([]byte, n)
	result.LzBuffers.Make(n)
	return result
}

func invertWord(word []byte) {
	for i := range word {
		word[i] ^= 1
	}
}

var resumeFromRepr = kingpin.Arg("resumeFrom", "Word to resume enumeration from").String()

func check(s []byte, l, z int) {
	if l > z {
		fmt.Println(words.ToRepr(s), l, z)
		if l-z > 1 {
			fmt.Printf("\tWarning! (%d)\n", l-z)
		}
	}
}

func main() {
	kingpin.Parse()

	startFrom := []byte{0, 0}
	if *resumeFromRepr != "" {
		startFrom = words.FromRepr(*resumeFromRepr)
		log.Printf("Resumed from %s\n", *resumeFromRepr)
	}

	for length := len(startFrom); ; length++ {
		fmt.Println("length =", length)
		start := time.Now()

		var startDigits []byte
		if length == len(startFrom) {
			startDigits = startFrom[:length-1]
		}

		enumerator.EnumerateWords(length-1, 2, startDigits,
			func() enumerator.Context {
				return NewBuffers(length)
			},
			func(prefix []byte, context enumerator.Context) {
				bufs := context.(*Buffers)
				s := bufs.s
				copy(s, prefix)
				s[length-1] = 0

				z := words.CountLz(s, &bufs.LzBuffers)

				check(s, words.CountDifferentLyndonWords(s), z)

				invertWord(s)
				check(s, words.CountDifferentLyndonWords(s), z)
			})

		seconds := time.Now().Sub(start).Seconds()
		if seconds > 10 {
			fmt.Printf("\tTime: %.1f s\n", seconds)
		}
	}
}
