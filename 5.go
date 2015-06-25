package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func main() {
		//scan
    sc.Split(bufio.ScanWords)
    m := nextInt()
    n := nextInt()
		p := make([]int,0)
		q := make([]int,0)
    for i := 0; i < n; i++ {
        p = append(p,nextInt())
        q = append(q,nextInt())
    }
    //fmt.Println(m,p,q)

		//calculate

}

