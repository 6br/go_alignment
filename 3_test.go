package main

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe(t *testing.T){
	Describe(t,"We have strings", func() {
		var lcs = NewLCS("acccagcagttaga","atatgcgggatgcg") //stringのGoにおける実装上、半角英数でなければならない。
		Context("and we calculate LCS",func() {
			lcs.Length()
			It("should be the correct string",func() {
				var s string =	lcs.Print(lcs.b,lcs.x,len(lcs.x),len(lcs.y))
				Expect(s).To(Equal,"aagcgga")
			})
		})
	})
}
