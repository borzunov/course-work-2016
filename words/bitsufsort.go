package words

import "sort"

type bitSuffixInfo struct {
	bits, mask uint64
	suffixes   []int
}

func (p bitSuffixInfo) Len() int      { return len(p.suffixes) }
func (p bitSuffixInfo) Swap(i, j int) { p.suffixes[i], p.suffixes[j] = p.suffixes[j], p.suffixes[i] }

func (p bitSuffixInfo) Less(j1, j2 int) bool {
	offset1 := p.suffixes[j1]
	offset2 := p.suffixes[j2]
	suf1 := (p.bits << uint(offset1)) & p.mask
	suf2 := (p.bits << uint(offset2)) & p.mask
	if suf1 != suf2 {
		return suf1 < suf2
	}
	return offset1 > offset2
}

func BitSufSort(s uint64, n int, bufs *SufSortBuffers) []int {
	mask := (uint64(1) << uint(n)) - 1

	suffixes := bufs.sa
	for i := range suffixes {
		suffixes[i] = i
	}

	info := bitSuffixInfo{s, mask, suffixes}
	sort.Sort(info)
	return info.suffixes
}
