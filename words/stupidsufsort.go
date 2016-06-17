package words

import "sort"

type suffixInfo struct {
	str      []byte
	suffixes []int
}

func (p suffixInfo) Len() int      { return len(p.suffixes) }
func (p suffixInfo) Swap(i, j int) { p.suffixes[i], p.suffixes[j] = p.suffixes[j], p.suffixes[i] }

func (p suffixInfo) Less(j1, j2 int) bool {
	s1, s2 := p.suffixes[j1], p.suffixes[j2]

	n := len(p.str)
	for i := 0; s1+i < n && s2+i < n; i++ {
		a := p.str[s1+i]
		b := p.str[s2+i]
		if a != b {
			return a < b
		}
	}
	return s1 > s2
}

func stupidSufSort(s []byte) []int {
	info := suffixInfo{append(s, 0), make([]int, len(s))}
	for i := range info.suffixes {
		info.suffixes[i] = i
	}

	sort.Sort(info)
	return info.suffixes
}
