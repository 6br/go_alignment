package alignment

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe(t *testing.T){
	Describe(t,"We have strings", func() {
		var lcs DPMatrix = NewLCS("acccagcagttaga","atatgcgggatgcg") //stringのGoにおける実装上、半角英数でなければならない。
		Context("and we calculate LCS",func() {
			lcs.Length()
			It("should be the correct string",func() {
				var lx,ly = lcs.Strlen()
				var s,_,_ string =	lcs.Print(lx,ly)
				Expect(s).To(Equal,"aagcgga")
			})
			It("should be the correct score",func() {
				var s int = lcs.Score()
				Expect(s).To(Equal,7)
			})
		})
	})
}
