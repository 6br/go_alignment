package alignment

import (
	. "github.com/r7kamura/gospel"
	"testing"
	. "./interface"
)

func TestDescribe2(t *testing.T){
	Describe(t,"We have the other strings", func() {
		var lcs DPMatrix = NewNW("gctagg","aattgaagg") //stringのGoにおける実装上、半角英数でなければならない。
		Context("and we calculate NW",func() {
			lcs.Length()
			var lx,ly = lcs.Strlen() 
			var p,q,r =lcs.Print(lx,ly)
			It("should be the correct string of 1st line",func() {
				Expect(p).To(Equal,"aattgaagg")
			})
			It("should be the correct string of 2nd line",func() {
				Expect(q).To(Equal,"  !  ! !!")
			})
			It("should be the correct string of 3rd line",func() {
				Expect(r).To(Equal,"gct--a-gg")
			})
			It("should be the correct score",func() {
				var s int = lcs.Score()
				Expect(s).To(Equal,-4)
			})
		})
	})
}
