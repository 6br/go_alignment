import "fmt"

type Dictionary interface {
	insert(int,string)
	search(int) string
	delete(int)
}

type Dict struct{
	Dict []string
	Stack []int
}

func (dict Dict) insert(key int,value string) {
	dict[key] = value
}

func (dict Dict) search(key int) {
	return dict[key]
}

func (dict Dict) delete(key int) {
  dict[key] = null
}

func (dict Dict) init(length int){
	dict
}

func main(){


}
