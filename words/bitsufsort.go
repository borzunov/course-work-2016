package words

import "sort"

type bitSuffixInfo struct {
	bits, mask uint64
	suffixes   []int
}

func (p bitSuffixInfo) Len() int      { return len(p.suffixes) }
func (p bitSuffixInfo) Swap(i, j int) { p.suffixes[i], p.suffixes[j] = p.suffixes[j], p.suffixes[i] }

func (p bitSuffixInfo) Less(j1, j2 int) bool {
	offset1, offset2 := p.suffixes[j1], p.suffixes[j2]
	return (p.bits<<uint(offset1))&p.mask < (p.bits<<uint(offset2))&p.mask
}

func BitSufSort(s []byte, bufs *QSufSortBuffers) []int {
	bits := uint64(0)
	for _, ch := range s {
		bits = (bits << 1) | uint64(ch)
	}
	n := len(s)
	mask := (uint64(1) << uint(n)) - 1

	suffixes := bufs.sa
	for i := range suffixes {
		suffixes[i] = i
	}

	info := bitSuffixInfo{bits, mask, suffixes}
	sort.Sort(info)
	return info.suffixes
}
