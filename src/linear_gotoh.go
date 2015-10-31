package alignment

//package main

import . "./interface"
import "fmt"
import "math"

type LGotoh struct {
	x string
	y string
	h [][][]int
	Constants
}

func NewLGotoh(y string, x string, settings Constants) *LGotoh {
	xlen := len(x) + 1
	ylen := len(y) + 1
	h := make([][][]int, 3)
	for i := 0; i < 3; i++ {
		h[i] = make([][]int, xlen)
		for y := 0; y < xlen; y++ {
			h[i][y] = make([]int, ylen)
		}
	}
	LGotoh := &LGotoh{x: x, y: y, h: h}
	LGotoh.Constants = settings
	return LGotoh
}

func (l LGotoh) Strlen() (int, int) {
	return len(l.x), len(l.y)
}

func (l LGotoh) Substitution(x int, y int) int {
	return l.Constants.Substitution(l.x[x-1], l.y[y-1])
}

func (l LGotoh) Score() int {
	e, _ := l.ScoreArgs(len(l.x), len(l.y))
	return e
}

func (l *LGotoh) Length() {
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
			var nexth int
			//Update H^1 (l.h[0])
			nexth, _ = Max(l.h[0][i-1][j-1], l.h[1][i-1][j-1], l.h[2][i-1][j-1])
			//fmt.Println(i,j)
			l.h[0][i][j] = nexth + l.Substitution(i, j)
			//Update H^2 (l.h[1])
			e, d := l.Geted()
			nexth, _ = Max(l.h[0][i-1][j]-d, l.h[1][i-1][j]-e, l.h[2][i-1][j]-d)
			l.h[1][i][j] = nexth
			//Update H^3 (l.h[2])
			nexth, _ = Max(l.h[0][i][j-1]-d, math.MinInt64, l.h[2][i][j-1]-e)
			l.h[2][i][j] = nexth
		}
	}
}

func (l LGotoh) ScoreArgs(x int, y int) (int, int) {
	e, k := Max(l.h[0][x][y], l.h[1][x][y], l.h[2][x][y])
	return e, k
}

func (l LGotoh) Print(i int, j int) (string, string, string) {
	_, arg := l.ScoreArgs(i, j)
	return l.Print_iter(i, j, arg)
}

func (l LGotoh) Print_iter(i int, j int, arg int) (string, string, string) {
	var p, q, r string
	if i <= 0 && j <= 0 {
		return "", "", ""
	} else if i <= 0 {
		//p,q,r = l.Print(i,j-1)
		for ; j > 0; j-- {
			p += "-"
			q += " "
			r += fmt.Sprintf("%c", l.y[j-1])
		}
		return p, q, r
	} else if j <= 0 {
		//p,q,r = l.Print(i-1,j)
		for ; i > 0; i-- {
			p += fmt.Sprintf("%c", l.x[i-1])
			q += " "
			r += "-"
		}
		return p, q, r
	}
	e, d := l.Geted()
	if arg == 0 {
		_, arg = l.ScoreArgs(i-1, j-1)
		p, q, r = l.Print_iter(i-1, j-1, arg)
		p += fmt.Sprintf("%c", l.x[i-1])
		q += "|"
		r += fmt.Sprintf("%c", l.y[j-1])
	} else if arg == 1 {
		_, arg = Max(l.h[0][i-1][j]-d, l.h[1][i-1][j]-e, l.h[2][i-1][j]-d)
		p, q, r = l.Print_iter(i-1, j, arg)
		p += fmt.Sprintf("%c", l.x[i-1])
		q += " "
		r += "-"
	} else {
		_, arg = Max(l.h[0][i][j-1]-d, math.MinInt64, l.h[2][i][j-1]-e)
		p, q, r = l.Print_iter(i, j-1, arg)
		p += "-"
		q += " "
		r += fmt.Sprintf("%c", l.y[j-1])
	}
	return p, q, r
}

/*
func main() {
	arr := [][]int{{1, -1, -1, -1}, {-1, 1, -1, -1}, {-1, -1, 1, -1}, {-1, -1, -1, 1}}
	charlist := "acgt"
	var settings = NewConstants(2, 1, arr, charlist)
	var lcs = NewLGotoh("ggatgcatgcatgc", "atgcatgcatgccc", *settings)
	lcs.Length()
	fmt.Println(lcs.h)
	fmt.Println(lcs.Score())
	fmt.Println(lcs.h[0][8][8])
	var lx, ly = lcs.Strlen()
	var p, q, r = lcs.Print(lx, ly)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
}
*/
