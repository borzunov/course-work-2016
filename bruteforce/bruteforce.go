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

func main() {
	kingpin.Parse()

	startFrom := []byte{0, 0}
	if *resumeFromRepr != "" {
		startFrom = words.FromRepr(*resumeFromRepr)
		log.Printf("Resumed from %s\n", *resumeFromRepr)
	}

	for length := len(startFrom); ; length++ {
		var startDigits []byte
		if length == len(startFrom) {
			startDigits = startFrom
		}

		fmt.Println("length =", length)
		start := time.Now()
		enumerator.EnumerateWords(length, 2, startDigits, NewBuffers, func(s []byte, bufs enumerator.Context) {
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
