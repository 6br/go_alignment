package main

import (
	. "./src"
	. "./src/interface"
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//fasta形式で併記されている文字列を読み取り、配列に返す。
func readfasta(i string) (string, string) {

	var reader *bufio.Reader
	var line []byte
	var err error
	var ary [2]string
	// ファイルを読み込みモードでオープン
	read_file, _ := os.OpenFile(i, os.O_RDONLY, 0600)
	// Readerを用意
	reader = bufio.NewReader(read_file)

	//line, _ , err = reader.ReadLine()
	var j = -1
	for {
		line, _, err = reader.ReadLine()
		// EOFなら終了
		if err == io.EOF {
			break
		}
		// 1行読み出す
		if line[0] == 62 {
			j++
			ary[j] = ""
		} else {
			//遅い
			ary[j] += strings.ToLower(string(line))
		}
	}
	return ary[0], ary[1]
}

//configファイルを読み取り、Constants構造体に格納
func readconfig(i string) Constants {
	var reader *bufio.Reader
	var line []byte
	var err error
	var ary [4][]string
	var ary2int [][]int

	// ファイルを読み込みモードでオープン
	read_file, _ := os.OpenFile(i, os.O_RDONLY, 0600)
	// Readerを用意
	reader = bufio.NewReader(read_file)

	line, _, err = reader.ReadLine()
	tmp := strings.Split(string(line), " ")

	d, _ := strconv.Atoi(tmp[0])
	e, _ := strconv.Atoi(tmp[1])

	for j := 0; ; j++ {
		line, _, err = reader.ReadLine()
		// EOFなら終了
		if err == io.EOF {
			break
		}
		// 1行読み出す
		ary[j] = strings.Fields(string(line))
		var tem []int
		for _, value := range ary[j] {
			temp, _ := strconv.Atoi(value)
			tem = append(tem, temp)
		}
		ary2int = append(ary2int, tem)
	}
	charlist := "acgt"
	var settings *Constants
	settings = NewConstants(d, e, ary2int, charlist)
	return *settings
}

func readfile(i string) string {
	var reader *bufio.Reader
	var line []byte
	var err error

	var ary string
	// ファイルを読み込みモードでオープン
	read_file, _ := os.OpenFile(i, os.O_RDONLY, 0600)
	// Readerを用意
	reader = bufio.NewReader(read_file)

	line, _, err = reader.ReadLine()
	for {
		// EOFなら終了
		if err == io.EOF {
			break
		}
		// 1行読み出す
		line, _, err = reader.ReadLine()
		//遅い
		ary += strings.ToLower(string(line))
	}
	return ary
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	var config string
	var interval int
	var debug bool
	flag.StringVar(&config, "c", "empty", "Write a pass of the config file including gap costs and a substitution matrix")
	flag.StringVar(&config, "config", "empty", "Write a pass of the config file including gap costs and a substitution matrix")
	flag.IntVar(&interval, "i", 50, "Write an interval of alignment result sequences (default: 50)")
	flag.BoolVar(&debug, "d", false, "Set true if you want to examine memory usage")
	flag.BoolVar(&debug, "debug", false, "Set true if you want to examine memory usage")
	flag.Parse()
	var ary string
	var ary2 string
	var settings Constants

	if config != "empty" {
		settings = readconfig(config)
	}

	if flag.Arg(1) == "" { //正規表現で、ドットを含むのであれば。
		ary = readfile("sequence.fasta")
		ary2 = readfile("sequence2.fasta")
	} else if m, _ := regexp.MatchString("\\.txt", flag.Arg(1)); m {
		ary, ary2 = readfasta(flag.Arg(1))
	} else if m, _ := regexp.MatchString("\\.", flag.Arg(2)); m {
		ary = readfile(flag.Arg(1))
		ary2 = readfile(flag.Arg(2))
	} else {
		ary = flag.Arg(1)
		ary2 = flag.Arg(2)
	}
	if debug {
		fmt.Println("Waiting...")
		fmt.Scanln()
	}
	var lcs DPMatrix
	switch flag.Arg(0) {
	case "1":
		lcs = NewLCS(ary, ary2)
	case "2":
		lcs = NewSW(ary, ary2)
	case "3":
		lcs = NewGotoh(ary, ary2, settings)
	case "4":
		lcs = NewMEA(ary, ary2, settings)
	case "5":
		lcs = NewLGotoh(ary, ary2, settings)
	case "6":
		lcs = NewRGotoh(ary, ary2, settings)
	default:
		lcs = NewNW(ary, ary2)
	}

	lcs.Length() // Exec alignment
	score := lcs.Score()
	var lx, ly = lcs.Strlen()
	if debug {
		fmt.Println("Waiting...")
		fmt.Scanln()
	}
	var p, q, r = lcs.Print(lx, ly)
	if debug {
		fmt.Println("Waiting...")
		fmt.Scanln()
	}
	j := 0
	fmt.Println("Score:", score)
	if interval >= 1 && interval < len(p) {
		for i := interval - 1; i < len(p)+interval; i += interval {
			if i > len(p) {
				i = len(p)
			}
			fmt.Println("from", j, "to", i)
			fmt.Println(r[j:i])
			fmt.Println(q[j:i])
			fmt.Println(p[j:i])
			j = i + 1
		}
	} else {
		fmt.Println(r)
		fmt.Println(q)
		fmt.Println(p)
	}
}
