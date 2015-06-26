package alignment

import "fmt"

const mu = 1
const s = 2

type NW struct {
	x string
	y string
	c [][]int
	b [][]string
}

func NewNW(y string,x string) *NW {
	xlen := len(x)+1
	ylen := len(y)+1
	c := make([][]int,xlen,xlen)
	b := make([][]string,xlen,xlen)
	for y:=0; y<xlen ; y++{
		c[y] = make([]int,ylen,ylen)
		b[y] = make([]string,ylen)
	}
	NW := &NW{x:x,y:y,c:c,b:b}
	return NW
}

func (l NW) Length() {
	var m = len(l.x)+1
	var n = len(l.y)+1
	for i:=1; i<m ;i++{
		l.c[i][0] = -i * s
	}
	for j:=1; j<n ;j++{
		l.c[0][j] = -j * s
	}

	for i:=1; i<m ;i++{
		for j:=1; j<n ;j++{
			if l.x[i-1]==l.y[j-1] {
				l.c[i][j] = l.c[i-1][j-1]+1
				l.b[i][j] = "!"
			} else {
				l.c[i][j] = l.c[i-1][j-1]-mu
				l.b[i][j] = " "
			}
			if l.c[i-1][j]-s>=l.c[i][j] {
				l.c[i][j] = l.c[i-1][j]-s
				l.b[i][j] = "|"
			}
			if l.c[i][j-1]-s>=l.c[i][j] {
				l.c[i][j] = l.c[i][j-1]-s
				l.b[i][j] = "-"
			}
		}
	}
}

func (l NW) Score() int {
	x,y := l.Strlen()
	return l.c[x][y]
}
func (l NW) Strlen() (int,int){
	return len(l.x),len(l.y)
}

func (l NW) Print(i int,j int) (string,string,string) {
	var p,q,r string
	if i==0 || j==0 {
		return "","",""
	}
	if l.b[i][j] == " " || l.b[i][j]=="!" {
		p,q,r = l.Print(i-1,j-1)
		//fmt.Printf("%c",x[i-1])
		p += fmt.Sprintf("%c",l.x[i-1])
		q += l.b[i][j]
		r += fmt.Sprintf("%c",l.y[j-1])
	} else if l.b[i][j] == "|" {
	    p,q,r =l.Print(i-1,j)
		p += fmt.Sprintf("%c",l.x[i-1])
		q += " "
		r += "-"

	} else {
		p,q,r =l.Print(i,j-1)
		p += "-"
		q += " "
		r += fmt.Sprintf("%c",l.y[j-1])
	}
	//fmt.Println(rt)
	return p,q,r
}
/*
func main() {
	var lcs = NewNW("gctagg","aattgaagg") //stringのGoにおける実装上、半角英数でなければならない。
	lcs.Length()
	fmt.Println(lcs.b)
	fmt.Println(lcs.c)
	var p,q,r =lcs.Print(lcs.b,lcs.x,lcs.y,len(lcs.x),len(lcs.y))
	fmt.Println(p)
fmt.Println(q)
fmt.Println(r)
}
*/

