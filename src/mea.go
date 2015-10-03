package alignment

//package main

import "fmt"
import . "./interface"

type MEA struct {
	x string
	y string
	c [][]int
	b [][]uint8
	Constants
}

func NewMEA(y string, x string, matrix Constants) *MEA {
	xlen := len(x) + 1
	ylen := len(y) + 1
	c := make([][]int, xlen, xlen)
	b := make([][]uint8, xlen, xlen)
	for y := 0; y < xlen; y++ {
		c[y] = make([]int, ylen, ylen)
		b[y] = make([]uint8, ylen)
	}
	MEA := &MEA{x: x, y: y, c: c, b: b}
	MEA.Constants = matrix
	return MEA
}

func (l MEA) Length() {
	var m = len(l.x) + 1
	var n = len(l.y) + 1
	for i := 1; i < m; i++ {
		l.c[i][0] = 0
	}
	for j := 1; j < n; j++ {
		l.c[0][j] = 0
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			l.c[i][j] = l.c[i-1][j-1] + l.Getscore(i, j)
			l.b[i][j] = 0
			if l.c[i-1][j] >= l.c[i][j] {
				l.c[i][j] = l.c[i-1][j]
				l.b[i][j] = 1
			}
			if l.c[i][j-1] >= l.c[i][j] {
				l.c[i][j] = l.c[i][j-1]
				l.b[i][j] = 2
			}
		}
	}
}

func (l MEA) Score() int {
	x, y := l.Strlen()
	return l.c[x][y]
}
func (l MEA) Strlen() (int, int) {
	return len(l.x), len(l.y)
}

func (l MEA) Print(i int, j int) (string, string, string) {
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
	if l.b[i][j] == 0 {
		p, q, r = l.Print(i-1, j-1)
		//fmt.Printf("%c",x[i-1])
		p += fmt.Sprintf("%c", l.x[i-1])
		r += fmt.Sprintf("%c", l.y[j-1])
		if l.x[i-1] == l.y[i-1] {
			q += "|"
		} else {
			q += "!"
		}
	} else if l.b[i][j] == 1 {
		p, q, r = l.Print(i-1, j)
		p += fmt.Sprintf("%c", l.x[i-1])
		q += " "
		r += "-"
	} else {
		p, q, r = l.Print(i, j-1)
		p += "-"
		q += " "
		r += fmt.Sprintf("%c", l.y[j-1])
	}
	//fmt.Println(rt)
	return p, q, r
}

/*
func main() {
  var matrix = [][]int{{1, 2, 3}, {7, 5, 6}, {4, 8, 9}}
	var lcs = NewMEA("abc","def", matrix) //stringのGoにおける実装上、半角英数でなければならない。
	lcs.Length()
	fmt.Println(lcs.b)
	fmt.Println(lcs.c)
	var lx,ly = lcs.Strlen()
	var p,q,r =lcs.Print(lx,ly)
	fmt.Println(p)
  fmt.Println(q)
  fmt.Println(r)
}
*/
