package alignment

import (
	. "./interface"
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe7(t *testing.T) {
	Describe(t, "We have the other strings", func() {
		arr := [][]int{{1, -1, -1, -1}, {-1, 1, -1, -1}, {-1, -1, 1, -1}, {-1, -1, -1, 1}}
		charlist := "acgt"
		var settings = NewConstants(2, 1, arr, charlist)
		var lcs = NewRGotoh("ggatgcatgcatgc", "atgcatgcatgccc", *settings)
		Context("and we calculate by RealLGotoh-algo", func() {
			lcs.Length()
			var lx, ly = lcs.Strlen()
			var p, q, r = lcs.Print(lx, ly)
			It("should be the correct string of 1st line", func() {
				Expect(p).To(Equal, "--atgcatgcatgccc")
			})
			It("should be the correct string of 2nd line", func() {
				Expect(q).To(Equal, "  |||||||||||  |")
			})
			It("should be the correct string of 3rd line", func() {
				Expect(r).To(Equal, "ggatgcatgcatg--c")
			})
			It("should be the correct score", func() {
				var s = lcs.Score()
				Expect(s).To(Equal, 6)
			})
		})
	})
}
