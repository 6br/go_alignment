package alignment

import "fmt"

const myu = 1
const sgm = 2

type SW struct {
	x string
	y string
	c [][]int
	b [][]string
}

func NewSW(y string,x string) *SW {
	xlen := len(x)+1
	ylen := len(y)+1
	c := make([][]int,xlen,xlen)
	b := make([][]string,xlen,xlen)
	for y:=0; y<xlen ; y++{
		c[y] = make([]int,ylen,ylen)
		b[y] = make([]string,ylen)
	}
	SW := &SW{x:x,y:y,c:c,b:b}
	return SW
}

func (l SW) Length() {
	var m = len(l.x)+1
	var n = len(l.y)+1
	for i:=1; i<m ;i++{
		l.c[i][0] = 0
	}
	for j:=1; j<n ;j++{
		l.c[0][j] = 0
	}

	for i:=1; i<m ;i++{
		for j:=1; j<n ;j++{
			if l.x[i-1]==l.y[j-1] {
				l.c[i][j] = l.c[i-1][j-1]+1
				l.b[i][j] = "!"
			} else {
				l.c[i][j] = l.c[i-1][j-1]-myu
				l.b[i][j] = " "
			}
			if l.c[i-1][j]-s>=l.c[i][j] {
				l.c[i][j] = l.c[i-1][j]-sgm
				l.b[i][j] = "|"
			}
			if l.c[i][j-1]-s>=l.c[i][j] {
				l.c[i][j] = l.c[i][j-1]-sgm
				l.b[i][j] = "-"
			}
			if 0>=l.c[i][j] {
				l.c[i][j] = 0
				l.b[i][j] = "aborted"
			}
		}
	}
}

func (l SW) Print(b [][]string,x string,y string,i int,j int) (string,string,string) {
	var p,q,r string
	if i==0 || j==0 || b[i][j]=="aborted" {
		return "","",""
	}
	if b[i][j] == " " || b[i][j]=="!" {
		p,q,r = l.Print(b,x,y,i-1,j-1)
		//fmt.Printf("%c",x[i-1])
		p += fmt.Sprintf("%c",x[i-1])
		q += b[i][j]
		r += fmt.Sprintf("%c",y[j-1])
	} else if b[i][j] == "|" {
	    	p,q,r =l.Print(b,x,y,i-1,j)
		p += fmt.Sprintf("%c",x[i-1])
		q += " "
		r += "-"

	} else {
		p,q,r =l.Print(b,x,y,i,j-1)
		p += "-"
		q += " "
		r += fmt.Sprintf("%c",y[j-1])
	}
	//fmt.Println(rt)
	return p,q,r
}
/*
func main() {
	var lcs = NewSW("gctagg","aattgaagg") //stringのGoにおける実装上、半角英数でなければならない。
	lcs.Length()
	fmt.Println(lcs.b)
	fmt.Println(lcs.c)
	var p,q,r =lcs.Print(lcs.b,lcs.x,lcs.y,len(lcs.x),len(lcs.y))
	fmt.Println(p)
fmt.Println(q)
fmt.Println(r)
}
*/

