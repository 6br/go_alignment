package alignment

import "strings"

type DPMatrix interface {
	Length()
	Print(int, int) (string, string, string)
	Strlen() (int, int)
	Score() int
}

type Constants struct {
	e        int
	d        int
	s        [][]int
	charlist []uint8
}

func NewConstants(d int, e int, s [][]int, c string) *Constants {
	var ch []uint8
	c = strings.ToLower(c)
	for i := 0; i < len(c); i++ {
		ch = append(ch, c[i])
	}
	NC := &Constants{e: e, d: d, s: s, charlist: ch}
	return NC
}

func (c Constants) Cost(l int) int {
	return -c.e*(l-1) - c.d
}

func (c Constants) Geted() (int, int) {
	return c.e, c.d
}

func (c Constants) arraysearch(query uint8) int {
	for i := 0; i < len(c.charlist); i++ {
		if query == c.charlist[i] {
			return i
		}
	}
	return -1
}

func (c Constants) Substitution(a uint8, b uint8) int {
	return c.s[c.arraysearch(a)][c.arraysearch(b)]
}

func (c Constants) Getscore(a int, b int) int {
	return c.s[a-1][b-1]
}

func Max(a int, b int, c int) (int, int) {
	if c > a && c > b {
		return c, 2
	} else if b > a && b > c {
		return b, 1
	} else {
		return a, 0
	}
}
