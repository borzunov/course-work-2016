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
	words.LzBuffers
}

func NewBuffers(n int) enumerator.Context {
	result := new(Buffers)
	result.LzBuffers.Make(n)
	return result
}

var resumeFromRepr = kingpin.Arg("resumeFrom", "Word to resume enumeration from").String()

func check(s uint64, n int, l, z int) {
	if l > z {
		fmt.Println(words.BitsToRepr(s, n), l, z)
		if l-z > 2 {
			fmt.Printf("\tWarning! (%d)\n", l-z)
		}
	}
}

func main() {
	kingpin.Parse()

	startWord, startWordLen := words.BitsFromRepr("00")
	if *resumeFromRepr != "" {
		startWord, startWordLen = words.BitsFromRepr(*resumeFromRepr)
		log.Printf("Resumed from %s\n", *resumeFromRepr)
	}

	for length := startWordLen; ; length++ {
		fmt.Println("length =", length)
		startTime := time.Now()

		curStartWord := uint64(0)
		if length == startWordLen {
			curStartWord = startWord >> 1
		}

		mask := (uint64(1) << uint(length)) - 1
		enumerator.EnumerateBinaryWords(length-1, curStartWord,
			func() enumerator.Context {
				return NewBuffers(length)
			},
			func(prefix uint64, context enumerator.Context) {
				bufs := context.(*Buffers)
				s := prefix << 1

				z := words.CountLz(s, length, &bufs.LzBuffers)

				check(s, length, words.CountDifferentLyndonWords(s, length), z)

				s ^= mask
				check(s, length, words.CountDifferentLyndonWords(s, length), z)
			})

		seconds := time.Now().Sub(startTime).Seconds()
		if seconds > 10 {
			fmt.Printf("\tTime: %.1f s\n", seconds)
		}
	}
}
