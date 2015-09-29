package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
	"flag"
	"regexp"
	. "./src"
)

//fasta形式で併記されている文字列を読み取り、配列に返す。
func readfasta(i string) []string {

	var reader *bufio.Reader
	var line []byte
	var err error
	var ary []string
	// ファイルを読み込みモードでオープン
	read_file, _ := os.OpenFile(i, os.O_RDONLY, 0600)
	// Readerを用意
	reader = bufio.NewReader(read_file)

	line, _ , err = reader.ReadLine()
	var j = -1
	for {
		// EOFなら終了
		if err == io.EOF {
			break
		}
		// 1行読み出す
		line,_, err = reader.ReadLine()
		if line[0] == '>' {
      j++
  	}else{
		  ary[j] += string(line)
	  }
	}
	fmt.Println(ary)
	return ary
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

	line, _ , err = reader.ReadLine()
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
	var ary string
	var ary2 string
	var array []string

	if flag.Arg(1)==""{ //正規表現で、ドットを含むのであれば。 
		ary = readfile("sequence.fasta")
		ary2 = readfile("sequence2.fasta")
	} else if m,_ := regexp.MatchString("\\.",flag.Arg(2)); m {
		ary = readfile(flag.Arg(1))
		ary2 = readfile(flag.Arg(2))
	} else if m,_ := regexp.MatchString("\\.",flag.Arg(1)); m {
		array = readfasta(flag.Arg(1))
		ary = array[0]
		ary2 = array[1]
	} else {
		ary = flag.Arg(1)
		ary2 = flag.Arg(2)
	}

	var lcs DPMatrix
	switch flag.Arg(0) {
		case "1": lcs = NewLCS(ary,ary2)
		case "2": lcs = NewSW(ary,ary2)
	  //case "3": lcs = NewGotoh(ary,ary2)
		default : lcs = NewNW(ary,ary2)
	}

	lcs.Length() // Exec alignment
	var lx,ly = lcs.Strlen()
	var p,q,r = lcs.Print(lx,ly)
	j:=0
	for i:=50;i<=len(p);i+=50 {
		fmt.Println("from",j,"to",i)
		fmt.Println(p[j:i])
		fmt.Println(q[j:i])
		fmt.Println(r[j:i])
		j=i+1
	}
}
