package alignment

//package main

import . "./interface"
import "fmt"
import "math"

//import "runtime"

type Gotoh struct {
	x   string
	y   string
	phi [][][]int
	h   [][][]int
	Constants
}

func NewGotoh(y string, x string, settings Constants) *Gotoh {
	xlen := len(x) + 1
	ylen := len(y) + 1
	phi := make([][][]int, 3)
	h := make([][][]int, 3)
	for i := 0; i < 3; i++ {
		phi[i] = make([][]int, xlen)
		h[i] = make([][]int, xlen)
		for y := 0; y < xlen; y++ {
			phi[i][y] = make([]int, ylen)
			h[i][y] = make([]int, ylen)
		}
	}
	Gotoh := &Gotoh{x: x, y: y, phi: phi, h: h}
	Gotoh.Constants = settings
	return Gotoh
}

func (l Gotoh) Strlen() (int, int) {
	return len(l.x), len(l.y)
}

func (l Gotoh) Substitution(x int, y int) int {
	return l.Constants.Substitution(l.x[x-1], l.y[y-1])
}

func (l Gotoh) Score() int {
	e, _ := l.ScoreArgs(len(l.x), len(l.y))
	return e
}

func (l Gotoh) ScoreArgs(x int, y int) (int, int) {
	e, k := Max(l.h[0][x][y], l.h[1][x][y], l.h[2][x][y])
	return e, k
}

func (l *Gotoh) Length() {
	//var r runtime.MemStats
	var m = len(l.x)
	var n = len(l.y)
	l.h[1][0][0] = math.MinInt64
	l.h[2][0][0] = math.MinInt64
	for i := 1; i <= m; i++ {
		l.h[0][i][0] = math.MinInt32
		l.h[1][i][0] = l.Cost(i)
		l.h[2][i][0] = math.MinInt32
	}
	for j := 1; j <= n; j++ {
		l.h[0][0][j] = math.MinInt32
		l.h[1][0][j] = math.MinInt32
		l.h[2][0][j] = l.Cost(j)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//runtime.ReadMemStats(&r)
			//fmt.Printf("%d,%d,%d,%d\n", r.HeapSys, r.HeapAlloc,r.HeapIdle, r.HeapReleased)
			var nexth int
			var nextphi int
			//Update H^1 (l.h[0])
			nexth, nextphi = Max(l.h[0][i-1][j-1], l.h[1][i-1][j-1], l.h[2][i-1][j-1])
			//fmt.Println(i,j)
			l.h[0][i][j] = nexth + l.Substitution(i, j)
			l.phi[0][i][j] = nextphi
			//Update H^2 (l.h[1])
			e, d := l.Geted()
			nexth, nextphi = Max(l.h[0][i-1][j]-d, l.h[1][i-1][j]-e, l.h[2][i-1][j]-d)
			l.h[1][i][j] = nexth
			l.phi[1][i][j] = nextphi
			//Update H^3 (l.h[2])
			nexth, nextphi = Max(l.h[0][i][j-1]-d, math.MinInt64, l.h[2][i][j-1]-e)
			l.h[2][i][j] = nexth
			l.phi[2][i][j] = nextphi
		}
	}
}

func (l Gotoh) Print(i int, j int) (string, string, string) {
	_, arg := l.ScoreArgs(i, j)
	return l.Print_iter(i, j, arg)
}

func (l Gotoh) Print_iter(i int, j int, arg int) (string, string, string) {
	var p, q, r string
	if i <= 0 && j <= 0 {
		return "", "", ""
	} else if i <= 0 {
		//p,q,r = l.Print(i,j-1)
		for jt := 0; jt < j; jt++ {
			p += "-"
			q += " "
			r += fmt.Sprintf("%c", l.y[jt])
		}
		return p, q, r
	} else if j <= 0 {
		//p,q,r = l.Print(i-1,j)
		for it := 0; it < i; it++ {
			p += fmt.Sprintf("%c", l.x[it])
			q += " "
			r += "-"
		}
		return p, q, r
	}

	if arg == 0 {
		arg = l.phi[arg][i][j]
		p, q, r = l.Print_iter(i-1, j-1, arg)
		p += fmt.Sprintf("%c", l.x[i-1])
		q += "|"
		r += fmt.Sprintf("%c", l.y[j-1])
	} else if arg == 1 {
		arg = l.phi[arg][i][j]
		p, q, r = l.Print_iter(i-1, j, arg)
		p += fmt.Sprintf("%c", l.x[i-1])
		q += " "
		r += "-"
	} else {
		arg = l.phi[arg][i][j]
		p, q, r = l.Print_iter(i, j-1, arg)
		p += "-"
		q += " "
		r += fmt.Sprintf("%c", l.y[j-1])
	}
	return p, q, r
}

/*
func main() {
	arr := [][]int{{1,-1,-1,-1},{-1,1,-1,-1},{-1,-1,1,-1},{-1,-1,-1,1}}
	charlist := "acgt"
	var settings = NewConstants(2, 1, arr, charlist)
	var lcs = NewGotoh("ggatgcatgcatgc","atgcatgcatgccc",*settings)
	lcs.Length()
	fmt.Println(lcs.h)
	fmt.Println(lcs.phi)
	fmt.Println(lcs.Score())
	fmt.Println(lcs.h[0][8][8])
	var lx,ly = lcs.Strlen()
	var p,q,r =lcs.Print(lx,ly)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
}
*/
