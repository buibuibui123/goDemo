package build

import "fmt"

type Fone struct {
}

// type Ftow struct {
// }
// type Fthird struct {
// }

func (*Fone) Pone(f func()) {
	fmt.Println("this is one")
	f()
}
func (*Fone) Ptow(f func()) {
	fmt.Println("this is tow")
	f()
}
func (*Fone) Pthird(f func()) {
	fmt.Println("this is thrid")
	f()
}
func (*Fone) Zero() {
	fmt.Println("this is Zero")

}
