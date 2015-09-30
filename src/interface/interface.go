package alignment

import "strings"

type DPMatrix interface {
	Length()
	Print(int,int) (string,string,string)
	Strlen() (int,int)
	Score() int
}

type Constants struct {
  e int
	d int
	s [][]int
	charlist []uint8
}

func NewConstants(e int,d int,s [][]int,c string) *Constants{
	var ch []uint8
  c = strings.ToLower(c)
	for i:=0; i<len(c); i++ {
		ch = append(ch, c[i])
	}
	NC := &Constants{e:e,d:d,s:s,charlist:ch}
	return NC
}

func (c Constants) Cost(l int) int {
	return -c.e * (l-1) - c.d
}

func (c Constants) Geted()(int, int){
  return c.e, c.d
}

func (c Constants) arraysearch(query uint8)(int){
	for i:=0; i < len(c.charlist); i++{
		if(query == c.charlist[i]){return i}
	}
	return -1
}

func (c Constants) Substitution(a uint8, b uint8)(int){
  return c.s[c.arraysearch(a)][c.arraysearch(b)]
}
