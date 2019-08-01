package build

// type Mux interface {
// 	HandleFunc(pattern string)
// }
type ServeStack struct {
	stk []func(func())
}

func New(f ...func(func())) *ServeStack {
	return &ServeStack{f}
}

func (p *ServeStack) HFunc(h func()) func() {
	return BuildFunc(h, p.stk...)
}

func BuildFunc(h func(), stk ...func(func())) func() {

	if len(stk) == 0 {
		return h
	}
	return BuildFunc(func() {
		stk[0](h)
	}, stk[1:]...)
}
