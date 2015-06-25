package alignment

type DPMatrix interface {
	Length()
	Print(int,int) (string,string,string)
	Strlen() (int,int)
}
