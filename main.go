package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"flag"
	. "./go"
)

func readfile(i string) string {

	var reader *bufio.Reader
	var line []byte
	var err error

	var ary string
	// ファイルを読み込みモードでオープン
	read_file, _ := os.OpenFile(i, os.O_RDONLY, 0600)
	// Readerを用意
	reader = bufio.NewReader(read_file)

		line,_, err = reader.ReadLine()
	for {
		// EOFなら終了
		if err == io.EOF {
			break
		}
		// 1行読み出す
		line,_, err = reader.ReadLine()
	ary += string(line)
	}
	return ary
}

func main() {
	flag.Parse()
	var ary string = readfile("sequence.fasta")
	var ary2 string = readfile("sequence2.fasta")

	var lcs DPMatrix
	switch flag.Arg(0) {
	case "1": lcs = NewLCS(ary,ary2)
	case "2": lcs = NewSW(ary,ary2)
	default : lcs = NewNW(ary,ary2)
	}

	//var lcs DPMatrix = NewLCS("attataatgtgct","ggattgtac") //stringのGoにおける実装上、半角英数でなければならない。
	lcs.Length()
	var lx,ly = lcs.Strlen()
	var p,q,r = lcs.Print(lx,ly)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
}
