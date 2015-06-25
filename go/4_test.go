package alignment

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe3(t *testing.T){
	Describe(t,"We have the other strings", func() {
		var lcs = NewSW("gctagg","aattgaagg") //stringのGoにおける実装上、半角英数でなければならない。
		Context("and we calculate LCS",func() {
			lcs.Length()
			var p,q,r =lcs.Print(lcs.b,lcs.x,lcs.y,len(lcs.x),len(lcs.y))
			It("should be the correct string of 1st line",func() {
				Expect(p).To(Equal,"agg")
			})
			It("should be the correct string of 2nd line",func() {
				Expect(q).To(Equal,"!!!")
			})
			It("should be the correct string of 3rd line",func() {
				Expect(r).To(Equal,"agg")
			})
		})
	})
}
