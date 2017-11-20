package main

func Opr(f func() interface{}) *Inst {
	return (Inst(f)).Code()
}

func Val(s *Symbol) {
	(Inst(func() interface{} { return s })).Code()
}
