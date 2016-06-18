package words

type LzBuffers struct {
	prev, next, stack []int
	QSufSortBuffers
}

func (bufs *LzBuffers) Make(n int) {
	bufs.prev = make([]int, n+1)
	bufs.next = make([]int, n+1)
	bufs.stack = make([]int, 1, n+2)

	bufs.QSufSortBuffers.Make(n)
}

func calcNearest(s []byte, bufs *LzBuffers) ([]int, []int) {
	sufArr := append(BitSufSort(s, &bufs.QSufSortBuffers), -1)

	// prev and next have a fake zeroth element to handle -1 (endings) in the suffix array
	prev := bufs.prev
	next := bufs.next
	stack := bufs.stack
	stack[0] = -1
	for _, elem := range sufArr {
		for {
			top := stack[len(stack)-1]
			if top <= elem {
				prev[elem+1] = top
				stack = append(stack, elem)
				break
			}
			next[top+1] = elem
			stack = stack[:len(stack)-1]
		}
	}
	return prev[1:], next[1:]
}

func CountLz(s []byte, bufs *LzBuffers) int {
	prev, next := calcNearest(s, bufs)

	n := len(s)
	result := 0
	for i := 0; i < n; {
		jPrev, jNext := prev[i], next[i]
		j := i
		for {
			if jPrev != -1 && (jPrev >= n || j >= n || s[jPrev] != s[j]) {
				jPrev = -1
			}
			if jNext != -1 && (jNext >= n || j >= n || s[jNext] != s[j]) {
				jNext = -1
			}
			if jPrev == -1 && jNext == -1 {
				break
			}

			j++
			if jPrev != -1 {
				jPrev++
			}
			if jNext != -1 {
				jNext++
			}
		}
		if i == j { // A new letter
			j++
		}

		//result = append(result, s[i:j])
		result++

		i = j
	}
	return result
}

func Lz(s []byte) [][]byte {
	bufs := new(LzBuffers)
	bufs.Make(len(s))

	prev, next := calcNearest(s, bufs)

	n := len(s)
	result := make([][]byte, 0)
	for i := 0; i < n; {
		jPrev, jNext := prev[i], next[i]
		j := i
		for {
			if jPrev != -1 && (jPrev >= n || j >= n || s[jPrev] != s[j]) {
				jPrev = -1
			}
			if jNext != -1 && (jNext >= n || j >= n || s[jNext] != s[j]) {
				jNext = -1
			}
			if jPrev == -1 && jNext == -1 {
				break
			}

			j++
			if jPrev != -1 {
				jPrev++
			}
			if jNext != -1 {
				jNext++
			}
		}
		if i == j { // A new letter
			j++
		}

		result = append(result, s[i:j])

		i = j
	}
	return result
}
