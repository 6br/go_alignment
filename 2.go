package main

import "fmt"

type LCS struct {
	x string
	y string
	c [][]int
	b [][]string
}

func NewLCS(y string,x string) *LCS {
	xlen := len(x)+1
	ylen := len(y)+1
	c := make([][]int,xlen,xlen)
	b := make([][]string,xlen,xlen)
	for y:=0; y<xlen ; y++{
		c[y] = make([]int,ylen,ylen)
		b[y] = make([]string,ylen)
	}
	LCS := &LCS{x:x,y:y,c:c,b:b}
	return LCS
}

func (l LCS) Length() {
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
				l.b[i][j] = "\\"
			} else if l.c[i-1][j]>=l.c[i][j-1] {
				l.c[i][j] = l.c[i-1][j]
				l.b[i][j] = "|"
			} else {
				l.c[i][j] = l.c[i][j-1]
				l.b[i][j] = "-"
			}
		}
	}
}

func (l LCS) Print(b [][]string,x string,i int,j int,rt string) {
	if i==0 || j==0 {
		return
	}
	if b[i][j] == "\\" {
	//	fmt.Printf("%c",x[i-1])
		l.Print(b,x,i-1,j-1,rt)
		fmt.Printf("%c",x[i-1])
	} else if b[i][j] == "|" {
	    l.Print(b,x,i-1,j,rt)
	} else {
		l.Print(b,x,i,j-1,rt)
	}
}

func main() {
	var lcs = NewLCS("acccagcagttaga","atatgcgggatgcg") //stringのGoにおける実装上、半角英数でなければならない。
	lcs.Length()
	fmt.Println(lcs.b)
	fmt.Println(lcs.c)
	lcs.Print(lcs.b,lcs.x,len(lcs.x),len(lcs.y),"")
}


