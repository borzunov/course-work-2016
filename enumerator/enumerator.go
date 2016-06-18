package enumerator

import (
	"math"
	"runtime"
	"sync"
)

func iterateBitMasksLinear(n int, start uint64, f func(uint64)) {
	end := uint64(1) << uint(n)
	for s := start; s < end; s++ {
		f(s)
	}
}

type Context interface{}

func EnumerateBinaryWords(n int, start uint64, newContext func() Context, f func(uint64, Context)) {
	threadExp := int(math.Ceil(math.Log2(float64(runtime.GOMAXPROCS(0)))))
	if n < threadExp+2 {
		bufs := newContext()
		iterateBitMasksLinear(n, start, func(s uint64) { f(s, bufs) })
		return
	}
	remN := n - threadExp

	var wg sync.WaitGroup
	iterateBitMasksLinear(threadExp, 0, func(suffix uint64) {
		wg.Add(1)
		bufs := newContext()
		go func() {
			iterateBitMasksLinear(remN, start>>uint(threadExp), func(s uint64) {
				f((s<<uint(threadExp))|suffix, bufs)
			})
			wg.Done()
		}()
	})
	wg.Wait()
}
