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
	charlist []string
}

func NewConstants(e int,d int,s [][]int,c []string) *Constants{
	NC := &Constants{e:e,d:d,s:s,charlist:c}
	return NC
}

func (c Constants) Cost(l int) int {
	return -c.e * (l-1) - c.d
}

func (c Constants) Geted()(int, int){
  return c.e, c.d
}

func (c Constants) arraysearch(query string)(int){
	for i:=0; i < len(c.charlist); i++{
		if(query == c.charlist[i]){return i}
	}
	return -1
}

func (c Constants) Substitution(a string, b string)(int){
  return c.s[c.arraysearch(a)][c.arraysearch(b)]
}
