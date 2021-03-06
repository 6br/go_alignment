package alignment

//package main

import . "./interface"
import "fmt"
import "math"

type Rgotoh struct {
	x     string
	y     string
	h     [][][]int
	vague bool
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
	l.RegionAlign(0, len(l.x), 0, len(l.y), true)
}

func (l *Rgotoh) RegionAlign(i1 int, i2 int, j1 int, j2 int, order bool) (int, int, int) {
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

	for j := 1; j <= n; j++ {
		l.h[0][0][1] = math.MinInt32
		l.h[1][0][1] = math.MinInt32
		l.h[2][0][1] = l.Cost(j)
		for i := 1; i <= m; i++ {
			var nexth int
			//Update H^1 (l.h[0])
			nexth, _ = Max(l.h[0][i-1][0], l.h[1][i-1][0], l.h[2][i-1][0])
			//fmt.Println(i, j)
			if order {
				l.h[0][i][1] = nexth + l.Substitution(i+i1, j+j1)
			} else {
				l.h[0][i][1] = nexth + l.Substitution(i2-i+1, j2-j+1)
			}
			//Update H^2 (l.h[1])
			e, d := l.Geted()
			nexth, _ = Max(l.h[0][i-1][1]-d, l.h[1][i-1][1]-e, l.h[2][i-1][1]-d)
			l.h[1][i][1] = nexth
			//Update H^3 (l.h[2])
			nexth, _ = Max(l.h[0][i][0]-d, math.MinInt64, l.h[2][i][0]-e)
			l.h[2][i][1] = nexth
		}
		for i := 0; i <= m; i++ {
			l.h[0][i][0] = l.h[0][i][1]
			l.h[1][i][0] = l.h[1][i][1]
			l.h[2][i][0] = l.h[2][i][1]
		}
	}
	return l.h[0][m][0], l.h[1][m][0], l.h[2][m][0] //l.ScoreArgs(m)
}

func (l Rgotoh) ScoreArgs(x int) (e int, k int) {
	e, k = Max(l.h[0][x][0], l.h[1][x][0], l.h[2][x][0])
	return
}

func (l Rgotoh) Print(i int, j int) (string, string, string) {
	//fmt.Println(l.LinearSpaceAlign(120, 120, 120, 120))
	//fmt.Println(l.RegionAlign(0, 7, 0, 0, true))
	if len(l.x) > 200 {
		l.vague = true
	}
	return l.LinearSpace(0, 0, i-1, j-1, l.Score())
	//return l.LinearBinaryAlign(0, i-1, 0, j-1, l.Score())
	//return l.linearSpaceAlign(0, len(l.x), 0, len(l.y))
}

func (l Rgotoh) pointMaxScore(i1 int, j1 int, i2 int, j2 int, ih int, jh int) int {
	tmp11, tmp12, tmp13 := l.RegionAlign(i1, ih, j1, jh, true)
	tmp21, tmp22, tmp23 := l.RegionAlign(ih, i2, jh, j2, false)
	_, d := l.Geted()
	array := make([]int, 8)
	array[0] = tmp11 + tmp21
	array[1] = tmp11 + tmp22
	array[2] = tmp11 + tmp23
	array[3] = tmp12 + tmp21
	array[4] = tmp12 + tmp22 + d
	array[5] = tmp12 + tmp23
	array[6] = tmp13 + tmp21
	array[7] = tmp13 + tmp23 + d

	tmp := math.MinInt64
	for _, v := range array {
		if tmp < v {
			tmp = v
		}
	}
	return tmp
}

func (l Rgotoh) LinearSpace(i0 int, j0 int, i int, j int, score int) (p string, q string, r string) {
	if (i < i0 && j == j0) || (j < j0 && i == i0) {
		return
	} else if i < i0 {
		for jt := j0; jt <= j; jt++ {
			p += "-"
			q += " "
			r += fmt.Sprintf("%c", l.y[jt])
		}
		return
	} else if j < j0 {
		for it := i0; it <= i; it++ {
			p += fmt.Sprintf("%c", l.x[it])
			q += " "
			r += "-"
		}
		return
	}
	maxj := j + 1 //To use determine whethere i have a suitable solution.
	min := 0
	if l.vague && j-70 >= 0 {
		min = j - 70
	}
	for jh := j; jh >= min; jh-- {
		tmp := l.pointMaxScore(0, 0, len(l.x), len(l.y), i, jh)
		//fmt.Println(tmp1, tmp2, err1, err2, tmp, jh)
		if tmp == score {
			maxj = jh
			break
		}
	}
	//fmt.Println(i, j, maxj, "result")
	p, q, r = l.LinearSpace(i0, j0, i-1, maxj-1, score)
	if maxj == j+1 {
		//When no suitable solution is not found for i.
		p += fmt.Sprintf("%c", l.x[i])
		q += " "
		r += "-"
	} else if maxj < j {
		//When no suitable solution is not found for j.
		p += fmt.Sprintf("%c", l.x[i])
		q += "|"
		r += fmt.Sprintf("%c", l.y[maxj])
		for maxj++; maxj <= j; maxj++ {
			p += "-"
			q += " "
			r += fmt.Sprintf("%c", l.y[maxj])
		}
	} else {
		//fmt.Println(i)
		p += fmt.Sprintf("%c", l.x[i])
		q += "|"
		r += fmt.Sprintf("%c", l.y[j])
	}
	return
}

func (l Rgotoh) LinearBinaryAlign(i1 int, i2 int, j1 int, j2 int, score int) (p string, q string, r string) {
	if i2-i1 <= 100 || j2-j1 <= 100 {
		return l.LinearSpace(i1, j1, i2, j2, score)
	}
	ih := ((i1 + i2) / 2)

	maxscore, maxj := -1000, j1+1
	for jh := j2; jh >= j1; jh-- {
		tmp := l.pointMaxScore(i1, j1, i2, j2, ih, jh)

		//fmt.Println(i1, i2, j1, j2, ih, jh, tmp)
		if tmp > maxscore {
			maxscore = tmp
			maxj = jh
		}
	}
	fmt.Println(i1, ih, i2, j1, maxj, j2)
	p, q, r = l.LinearBinaryAlign(i1, ih-1, j1, maxj-1, score)
	//p += fmt.Sprintf("%c", l.x[ih])
	//q += "|"
	//r += fmt.Sprintf("%c", l.y[maxj])
	p2, q2, r2 := l.LinearBinaryAlign(ih, i2, maxj, j2, score)
	p += p2
	q += q2
	r += r2
	return
}
