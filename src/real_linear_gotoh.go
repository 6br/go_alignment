package alignment

//package main

import . "./interface"
import "fmt"
import "math"

type Rgotoh struct {
	x string
	y string
	h [][][]int
	Constants
}

func NewRGotoh(y string, x string, settings Constants) *Rgotoh {
	xlen := len(x) + 1
	h := make([][][]int, 3)
	for i := 0; i < 3; i++ {
		h[i] = make([][]int, xlen)
		for y := 0; y < xlen; y++ {
			h[i][y] = make([]int, 2)
		}
	}
	Rgotoh := &Rgotoh{x: x, y: y, h: h}
	Rgotoh.Constants = settings
	return Rgotoh
}

func (l Rgotoh) Strlen() (int, int) {
	return len(l.x), len(l.y)
}

func (l Rgotoh) Substitution(x int, y int) int {
	return l.Constants.Substitution(l.x[x-1], l.y[y-1])
}

func (l Rgotoh) Score() int {
	e, _ := l.ScoreArgs(len(l.x))
	return e
}

func (l *Rgotoh) Length() {
	var m = len(l.x)
	var n = len(l.y)
	l.h[0][0][0] = 0
	l.h[1][0][0] = math.MinInt64
	l.h[2][0][0] = math.MinInt64
	for i := 1; i <= m; i++ {
		l.h[0][i][0] = math.MinInt32
		l.h[1][i][0] = l.Cost(i)
		l.h[2][i][0] = math.MinInt32
	}
	for j := 1; j <= 1; j++ {
		l.h[0][0][j] = math.MinInt32
		l.h[1][0][j] = math.MinInt32
		l.h[2][0][j] = l.Cost(j)
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			var nexth int
			//Update H^1 (l.h[0])
			nexth, _ = Max(l.h[0][i-1][0], l.h[1][i-1][0], l.h[2][i-1][0])
			//fmt.Println(i,j)
			l.h[0][i][1] = nexth + l.Substitution(i, j)
			//Update H^2 (l.h[1])
			e, d := l.Geted()
			nexth, _ = Max(l.h[0][i-1][1]-d, l.h[1][i-1][1]-e, l.h[2][i-1][1]-d)
			l.h[1][i][1] = nexth
			//Update H^3 (l.h[2])
			nexth, _ = Max(l.h[0][i][0]-d, math.MinInt64, l.h[2][i][0]-e)
			l.h[2][i][1] = nexth
		}
		for i := 1; i <= m; i++ {
			l.h[0][i][0] = l.h[0][i][1]
			l.h[1][i][0] = l.h[1][i][1]
			l.h[2][i][0] = l.h[2][i][1]
		}
		l.h[0][0][0] = math.MinInt32
		l.h[1][0][0] = math.MinInt32
		l.h[2][0][0] = l.Cost(j)
	}
}

func (l *Rgotoh) RegionAlign(i1 int, i2 int, j1 int, j2 int) (result int, class int) {
	var m = i2 - i1
	var n = j2 - j1
	l.h[0][0][0] = 0
	l.h[1][0][0] = math.MinInt64
	l.h[2][0][0] = math.MinInt64
	for i := 1; i <= m; i++ {
		l.h[0][i][0] = math.MinInt32
		l.h[1][i][0] = l.Cost(i)
		l.h[2][i][0] = math.MinInt32
	}
	for j := 1; j <= 1; j++ {
		l.h[0][0][j] = math.MinInt32
		l.h[1][0][j] = math.MinInt32
		l.h[2][0][j] = l.Cost(j)
	}

	for j := 1; j <= n; j++ {
		for i := 1; i <= m; i++ {
			var nexth int
			//Update H^1 (l.h[0])
			nexth, _ = Max(l.h[0][i-1][0], l.h[1][i-1][0], l.h[2][i-1][0])
			//fmt.Println(i, j)
			l.h[0][i][1] = nexth + l.Substitution(i+i1, j+j1)
			//Update H^2 (l.h[1])
			e, d := l.Geted()
			nexth, _ = Max(l.h[0][i-1][1]-d, l.h[1][i-1][1]-e, l.h[2][i-1][1]-d)
			l.h[1][i][1] = nexth
			//Update H^3 (l.h[2])
			nexth, _ = Max(l.h[0][i][0]-d, math.MinInt64, l.h[2][i][0]-e)
			l.h[2][i][1] = nexth
		}
		for i := 1; i <= m; i++ {
			l.h[0][i][0] = l.h[0][i][1]
			l.h[1][i][0] = l.h[1][i][1]
			l.h[2][i][0] = l.h[2][i][1]
		}
		l.h[0][0][0] = math.MinInt32
		l.h[1][0][0] = math.MinInt32
		l.h[2][0][0] = l.Cost(j)
	}

	//e, k := Max(l.h[0][m][0], l.h[1][m][0], l.h[2][m][0])
	return l.ScoreArgs(m)
}

func (l Rgotoh) ScoreArgs(x int) (e int, k int) {
	e, k = Max(l.h[0][x][0], l.h[1][x][0], l.h[2][x][0])
	return
}

func (l Rgotoh) Print(i int, j int) (string, string, string) {
	//fmt.Println(l.LinearSpaceAlign(120, 120, 120, 120))
	//fmt.Println(l.RegionAlign(0, len(l.x), 0, len(l.y)))
	return l.LinearSpace(i-1, j-1, l.Score())
}

func (l Rgotoh) LinearSpace(i int, j int, score int) (p string, q string, r string) {
	if (i < 0 && j == 0) || (j < 0 && i == 0) {
		return
	} else if i < 0 {
		for jt := 0; jt <= j; jt++ {
			p += "-"
			q += " "
			r += fmt.Sprintf("%c", l.y[jt])
		}
		return
	} else if j < 0 {
		for it := 0; it <= i; it++ {
			p += fmt.Sprintf("%c", l.x[it])
			q += " "
			r += "-"
		}
		return
	}
	maxj := j + 1 //To use determine whethere i have a suitable solution.
	for jh := 0; jh <= j; jh++ {
		tmp1, err1 := l.RegionAlign(0, i, 0, jh)
		tmp2, err2 := l.RegionAlign(i, len(l.x), jh, len(l.y))
		tmp := tmp1 + tmp2
		_, d := l.Geted()
		if err1 >= 1 && err2 >= 1 {
			tmp += d
		}
		if tmp == score {
			maxj = jh
		}
	}
	//fmt.Println(i, j, maxj, "result")
	p, q, r = l.LinearSpace(i-1, maxj-1, score)
	if maxj == j+1 {
		//iに適解が見つからなかった時
		p += fmt.Sprintf("%c", l.x[i])
		q += " "
		r += "-"
	} else if maxj < j {
		//jに適解が見つからなかった時
		p += fmt.Sprintf("%c", l.x[i])
		q += "|"
		r += fmt.Sprintf("%c", l.y[maxj])
		for maxj++; maxj <= j; maxj++ {
			p += "-"
			q += " "
			r += fmt.Sprintf("%c", l.y[maxj])
		}
	} else {
		p += fmt.Sprintf("%c", l.x[i])
		q += "|"
		r += fmt.Sprintf("%c", l.y[j])
	}
	return
}

func (l Rgotoh) LinearSpaceAlign(i1 int, i2 int, j1 int, j2 int) (p string, q string, r string) {
	i := i2 - i1
	j := j2 - j1
	fmt.Println(i1, i2, j1, j2, i, j, "hantei")
	if i <= 0 && j <= 0 {
		return "", "", ""
	} else if i <= 0 {
		//p,q,r = l.Print(i,j-1)
		for ; j2 > j1; j2-- {
			p += "-"
			q += " "
			r += fmt.Sprintf("%c", l.y[j2])
		}
		return p, q, r
	} else if j <= 0 {
		//p,q,r = l.Print(i-1,j)
		for ; i2 > i1; i2-- {
			p += fmt.Sprintf("%c", l.x[i2])
			q += " "
			r += "-"
		}
		return p, q, r
	} else if i == 1 && j == 1 {
		//	p += fmt.Sprintf("%c", l.x[i1])
		//	q += "|"
		//	r += fmt.Sprintf("%c", l.y[i1])
		return p, q, r
	}
	ih := ((i1 + i2) / 2)
	if ih >= len(l.x) || ih < i1 {
		panic(i1 + i2)
	}
	maxscore, maxj := math.MinInt64, j1+1
	for jh := j1 + 1; jh < j2; jh++ {
		tmp1, err1 := l.RegionAlign(i1, ih, j1, jh)
		tmp2, err2 := l.RegionAlign(ih, i2, jh, j2)
		fmt.Println(tmp1, tmp2)
		tmp := tmp1 + tmp2
		_, d := l.Geted()
		if err1 >= 1 && err2 >= 1 {
			tmp -= d
		}
		fmt.Println(i1, i2, j1, j2, ih, jh, tmp)
		if tmp > maxscore {
			maxscore = tmp
			maxj = jh
		}
	}
	fmt.Println(i1, ih, i2, j1, maxj, j2)
	p, q, r = l.LinearSpaceAlign(i1, ih, j1, maxj)
	p += fmt.Sprintf("%c", l.x[ih])
	q += "!"
	r += fmt.Sprintf("%c", l.y[maxj])
	p2, q2, r2 := l.LinearSpaceAlign(ih, i2, maxj, j2)
	p += p2
	q += q2
	r += r2
	return
}

/*
func main() {
	arr := [][]int{{1, -1, -1, -1}, {-1, 1, -1, -1}, {-1, -1, 1, -1}, {-1, -1, -1, 1}}
	charlist := "acgt"
	var settings = NewConstants(2, 1, arr, charlist)
	var lcs = NewRgotoh("ggatgcatgcatgc", "atgcatgcatgccc", *settings)
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
