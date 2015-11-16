package alignment

import (
	. "./interface"
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe5(t *testing.T) {
	Describe(t, "We have the other strings", func() {
		arr := [][]int{{1, 2, 3}, {7, 5, 6}, {4, 8, 9}}
		charlist := "acgt"
		var matrix = NewConstants(2, 1, arr, charlist)
		var lcs DPMatrix = NewMEA("gcg", "atg", *matrix)
		Context("and we calculate MEA", func() {
			lcs.Length()
			var lx, ly = lcs.Strlen()
			var p, q, r = lcs.Print(lx, ly)
			It("should be the correct string of 1st line", func() {
				Expect(p).To(Equal, "at-g")
			})
			It("should be the correct string of 2nd line", func() {
				Expect(q).To(Equal, " ! |")
			})
			It("should be the correct string of 3rd line", func() {
				Expect(r).To(Equal, "-gcg")
			})
			It("should be the correct score", func() {
				var s int = lcs.Score()
				Expect(s).To(Equal, 16)
			})
		})
	})
}
