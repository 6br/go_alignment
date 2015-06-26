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
    q := make([]int,0)
    r := make([]int,0)
    sum_q := 0
    sum_r := 0
    
    for i := 0; i < n; i++ {
        q = append(q,nextInt())
        sum_q += q[len(q)-1]
        r = append(r,nextInt())
        sum_r += r[len(r)-1]
    }
    //fmt.Println(m,q,r,sum_q,sum_r)

    //calculate
    dp_m := sum_q - m
    dp := make([]int,dp_m+1)
    for i:=0; i<n ;i++ {
        r_tmp := r[i]
        q_tmp := q[i]
        //fmt.Println(dp,r[i],q[i])
        for j:=dp_m; j>= q_tmp; j-- {
            if dp[j] < dp[j-q_tmp]+r_tmp{
                dp[j] = dp[j-q_tmp]+r_tmp
            }
        }
    }
    fmt.Println(sum_r-dp[dp_m])
}

