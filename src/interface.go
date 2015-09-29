package alignment

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
}

func NewConstants(e int,d int,s [][]int) *Constants{
	NC := &Constants{e:e,d:d,s:s}
	return NC
}

func (c Constants) Cost(l int) int {
	return -c.e * (l-1) - c.d
}
