package enumerator

import (
	"math"
	"runtime"
	"sync"
)

func iterateMasksLinear(n int, base byte, digits []byte, f func([]byte)) {
	if n == 0 {
		f(digits)
		return
	}

	if digits == nil {
		digits = make([]byte, n)
	}
	if len(digits) != n {
		panic("wrong digit buffer length")
	}

	for {
		f(digits)

		digits[n-1]++
		for i := n - 1; ; i-- {
			if digits[i] == base {
				digits[i] = 0
				if i == 0 {
					return
				}
				digits[i-1]++
			} else {
				break
			}
		}
	}
}

type Context interface{}

func EnumerateWords(n int, base byte, startDigits []byte, newContext func() Context, f func([]byte, Context)) {
	if startDigits != nil && len(startDigits) != n {
		panic("wrong startDigits length")
	}

	threadExp := int(math.Ceil(math.Log(float64(runtime.GOMAXPROCS(0))) / math.Log(float64(base))))
	if n < threadExp+2 {
		bufs := newContext()
		iterateMasksLinear(n, base, startDigits, func(s []byte) { f(s, bufs) })
		return
	}
	remN := n - threadExp

	var wg sync.WaitGroup
	iterateMasksLinear(threadExp, base, nil, func(suffix []byte) {
		digits := make([]byte, n)
		copy(digits, startDigits)
		copy(digits[remN:], suffix)

		wg.Add(1)
		bufs := newContext()
		go func() {
			iterateMasksLinear(remN, base, digits[:remN], func(_ []byte) {
				f(digits, bufs)
			})
			wg.Done()
		}()
	})
	wg.Wait()
}

func iterateBitMasksLinear(n int, start uint64, f func(uint64)) {
	end := uint64(1) << uint(n)
	for s := start; s < end; s++ {
		f(s)
	}
}

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
