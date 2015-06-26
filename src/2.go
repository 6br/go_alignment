package alignment

import (
	"fmt"
	"testing"
)

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

func (l LCS) Strlen() (int,int){
	return len(l.x),len(l.y)
}

func (l LCS) Print(i int,j int) (string,string,string) {
	if i==0 || j==0 {
		return "","",""
	}
	var p string;
	if l.b[i][j] == "\\" {
	//	fmt.Printf("%c",x[i-1])
		p,_,_=l.Print(i-1,j-1)
		//fmt.Printf("%c",x[i-1])
		p += fmt.Sprintf("%c",l.x[i-1])
	} else if l.b[i][j] == "|" {
	    p,_,_=l.Print(i-1,j)
	} else {
		p,_,_=l.Print(i,j-1)
	}
	return p,"",""
}
func (l LCS) Score() int {
	var lx,ly = l.Strlen() 
	p,_,_ := l.Print(lx,ly)
	return len(p)
}

func Benchmark(b *testing.B){
var lcs = NewLCS("jkahkncjknewrkfiljsklhlsfhskujejjflwjklnvmxcnvlcsdjfjjsdljdslfjsljfhedjwljkshfuejcklsjs","* klvnwoihwoihtewkllnxcnvsmvmjsdnjkjnshuvhsuiujeijwiodkakcopjnsdvsvbfsfvxjhsduifjkshskfrf")
	lcs.Length()
/*	fmt.Println(lcs.b)
	fmt.Println(lcs.c)*/
	fmt.Println(lcs.Print(len(lcs.x),len(lcs.y)))
}
/*
func main() {
	var lcs DPMatrix = NewLCS("attataatgtgct","ggattgtac") //stringのGoにおける実装上、半角英数でなければならない。
	lcs.Length()
//	fmt.Println(lcs.b)
//	fmt.Println(lcs.c)
	var lx,ly = lcs.Strlen() 
	fmt.Println(lcs.Print(lx,ly))
}*/


