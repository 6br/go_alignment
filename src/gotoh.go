
//package alignment
package main

import . "./interface"
import "fmt"
import "math"

type Gotoh struct {
	x string
	y string
	phi [][][]int
	h [][][]int
	maxscore int
	xmax int
	ymax int
	settings Constants
}

func NewGotoh(y string,x string,settings Constants) *Gotoh {
	xlen := len(x)+1
	ylen := len(y)+1
	phi := make([][][]int,3)
	h := make([][][]int,3)
	for i:=0; i<3; i++{
		phi[i] = make([][]int,xlen)
		h[i] = make([][]int,xlen)
	for y:=0; y<xlen ; y++{
		phi[i][y] = make([]int,ylen)
		h[i][y] = make([]int,ylen)
	}
}
	Gotoh := &Gotoh{x:x,y:y,phi:phi,h:h,maxscore:0,settings:settings}
	return Gotoh
}
func (l Gotoh) Strlen() (int,int){
	//return len(l.x),len(l.y)
	return l.xmax,l.ymax
}
func (l Gotoh) Score() int {
	return l.maxscore
}

func (l *Gotoh) Length() {
	var m = len(l.x)+1
	var n = len(l.y)+1
	l.h[1][0][0] = math.MinInt64
	l.h[2][0][0] = math.MinInt64
	for i:=1; i<m ;i++{
		l.h[1][i][0] = l.settings.Cost(i)
	}
	for j:=1; j<n ;j++{
		l.h[2][0][j] = l.settings.Cost(j) 
	}

/*
	for i:=1; i<m ;i++{
		for j:=1; j<n ;j++{
			if l.x[i-1]==l.y[j-1] {
				l.c[i][j] = l.c[i-1][j-1]+1
				l.b[i][j] = "!"
			} else {
				l.c[i][j] = l.c[i-1][j-1]-myu
				l.b[i][j] = " "
			}
			if l.c[i-1][j]-sgm>=l.c[i][j] {
				l.c[i][j] = l.c[i-1][j]-sgm
				l.b[i][j] = "|"
			}
			if l.c[i][j-1]-sgm>=l.c[i][j] {
				l.c[i][j] = l.c[i][j-1]-sgm
				l.b[i][j] = "-"
			}
			if 0>=l.c[i][j] {
				l.c[i][j] = 0
				l.b[i][j] = "aborted"
			}
			if l.maxscore<l.c[i][j] {
				l.maxscore = l.c[i][j]
				l.xmax = i
				l.ymax = j
			}
		}
	}
*/
}


func (l Gotoh) Print(int,int)(string, string, string){
  return "a","e","e"
}

func main() {
	//arr := make([][]int,0)
	arr := [][]int{{1,-1,-1,-1},{-1,1,-1,-1},{-1,-1,1,-1},{-1,-1,-1,1}}
	var settings = NewConstants(7,1,arr)
	var lcs = NewGotoh("gctaggaa","aattgaag",*settings) //stringのGoにおける実装上、半角英数でなければならない。
	lcs.Length()
	fmt.Println(lcs.h)
	fmt.Println(lcs.phi)
	fmt.Println(lcs.xmax)
	var lx,ly = lcs.Strlen()
	var p,q,r =lcs.Print(lx,ly)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
}


