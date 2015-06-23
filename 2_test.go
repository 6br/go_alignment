package 2

import (
	"gospel"
	"testing"
)

func TestDescribe(t *testing.T){
	Describe(t,"We have strings", func() {
		var lcs = NewLCS("acccagcagttaga","atatgcgggatgcg") //stringのGoにおける実装上、半角英数でなければならない。
		Context("and we calcucalted "),func() {
			lcs.Length()
			It("should be the a",func() {
				Expect().To(Equal,"aagcgga")
				
				lcs.Print(lcs.b,lcs.x,len(lcs.x),len(lcs.y),""))
			})
		})
	})
}
