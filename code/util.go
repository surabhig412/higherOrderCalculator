package code

import "github.com/surabhig412/hoc/symbol"

func Opr(f func() interface{}) *Inst {
	return (Inst(f)).Code()
}

func Val(s *symbol.Symbol) {
	(Inst(func() interface{} { return s })).Code()
}
