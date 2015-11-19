package main

import (
	. "./src"
	"testing"
)

func Benchmark3A(b *testing.B) {
	//ary, ary2 := readfasta("sequences.txt")
	ary := readfile("NM_000690.fasta")
	ary2 := readfile("XM_0062493585.2.fasta")
	settings := readconfig("config.txt")
	b.ResetTimer()
	lcs := NewGotoh(ary, ary2, settings)
	lcs.Length() // Exec alignment
	var lx, ly = lcs.Strlen()
	var _, _, _ = lcs.Print(lx, ly)
}

func Benchmark5A(b *testing.B) {
	//ary, ary2 := readfasta("sequences.txt")
	ary := readfile("NM_000690.fasta")
	ary2 := readfile("XM_0062493585.2.fasta")
	settings := readconfig("config.txt")
	b.ResetTimer()
	lcs := NewLGotoh(ary, ary2, settings)
	lcs.Length() // Exec alignment
	var lx, ly = lcs.Strlen()
	var _, _, _ = lcs.Print(lx, ly)
}

func Benchmark3B(b *testing.B) {
	ary, ary2 := readfasta("sequences.txt")
	//ary := readfile("NM_000690.fasta")
	//ary2 := readfile("XM_0062493585.2.fasta")
	settings := readconfig("config.txt")
	b.ResetTimer()
	lcs := NewGotoh(ary, ary2, settings)
	lcs.Length() // Exec alignment
	var lx, ly = lcs.Strlen()
	var _, _, _ = lcs.Print(lx, ly)
}

func Benchmark5B(b *testing.B) {
	ary, ary2 := readfasta("sequences.txt")
	//ary := readfile("NM_000690.fasta")
	//ary2 := readfile("XM_0062493585.2.fasta")
	settings := readconfig("config.txt")
	b.ResetTimer()
	lcs := NewLGotoh(ary, ary2, settings)
	lcs.Length() // Exec alignment
	var lx, ly = lcs.Strlen()
	var _, _, _ = lcs.Print(lx, ly)
}

func Benchmark6B(b *testing.B) {
	ary, ary2 := readfasta("sequences.txt")
	//ary := readfile("NM_000690.fasta")
	//ary2 := readfile("XM_0062493585.2.fasta")
	settings := readconfig("config.txt")
	b.ResetTimer()
	lcs := NewRGotoh(ary, ary2, settings)
	lcs.Length() // Exec alignment
	var lx, ly = lcs.Strlen()
	var _, _, _ = lcs.Print(lx, ly)
}
