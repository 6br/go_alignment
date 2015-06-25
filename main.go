package main

import (
	"fmt"
	. "./go"
)

func main() {
	var lcs DPMatrix = NewLCS("attataatgtgct","ggattgtac") //stringのGoにおける実装上、半角英数でなければならない。
	lcs.Length()
	var lx,ly = lcs.Strlen() 
	fmt.Println(lcs.Print(lx,ly))
}
