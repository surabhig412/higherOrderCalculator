package main

import (
	"fmt"
	"log"
	"math"
)

type Datum struct {
	Val float64
	Sym Symbol
}

type Inst func() interface{}

var stack []Datum
var Prog []Inst
var pc Inst
var pCounter int = 0
var WhileBodyCounter, WhileNextCounter, IfBodyCounter, IfNextCounter, IfElseCounter, CondCounter int

var STOP Inst = nil

func (d Datum) Push() {
	fmt.Println("push called")
	stack = append(stack, d)
}

func Pop() interface{} {
	fmt.Println("pop called")
	length := len(stack)
	if length == 0 {
		log.Fatalf("stack underflow")
	}
	poppedDatum := stack[length-1]
	stack = stack[:length-1]
	return poppedDatum
}

func (inst Inst) Code() *Inst {
	fmt.Println("code called: ", inst)
	var oldProg *Inst
	if len(Prog) != 0 {
		oldProg = &Prog[len(Prog)-1]
	}
	Prog = append(Prog, inst)
	return oldProg
}

func Execute() {
	fmt.Println("execute called: prog: ", Prog)
	pc = Prog[pCounter]
	if pc == nil {
		pCounter++
		pc = Prog[pCounter]
	}
	for pc != nil {
		pc()
		pCounter++
		pc = Prog[pCounter]
	}
	fmt.Println("execute finished")
}

func Constpush() interface{} {
	pCounter++
	pc = Prog[pCounter]
	fmt.Println("constpush: ", pc, pc())
	d := &Datum{Sym: *pc().(*Symbol)}
	d.Val = d.Sym.Val
	d.Push()
	return nil
}

func Varpush() interface{} {
	pCounter++
	pc = Prog[pCounter]
	d := &Datum{Sym: *pc().(*Symbol)}
	d.Push()
	return nil
}

func Add() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	d1.Val += d2.Val
	d1.Push()
	return nil
}

func Mod() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	d1.Val = math.Mod(d1.Val, d2.Val)
	d1.Push()
	return nil
}

func Sub() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	d1.Val -= d2.Val
	d1.Push()
	return nil
}

func Mul() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	d1.Val *= d2.Val
	d1.Push()
	return nil
}

func Div() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	d1.Val /= d2.Val
	d1.Push()
	return nil
}

func Power() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	d1.Val = math.Pow(d1.Val, d2.Val)
	d1.Push()
	return nil
}

func Negate() interface{} {
	d1 := Pop().(Datum)
	d1.Val = -d1.Val
	d1.Push()
	return nil
}

func Eval() interface{} {
	d := Pop().(Datum)
	if d.Sym.Type == UNDEF {
		log.Fatalf("undefined variable ", d.Sym.Name)
	}
	d.Val = d.Sym.Val
	d.Push()
	return nil
}

func Assign() interface{} {
	d1 := Pop().(Datum)
	d2 := Pop().(Datum)
	if d1.Sym.Type != VAR && d1.Sym.Type != UNDEF {
		log.Fatalf("assignment to non-variable ", d1.Sym.Name)
	}
	// d1.Sym.Val = d2.Val
	// d1.Sym.Type = symbol.VAR
	d2.Push()
	return nil
}

func Print() interface{} {
	d := Pop().(Datum)
	fmt.Printf("\t%v\n", d.Val)
	return nil
}

func Bltin() interface{} {
	d := Pop().(Datum)
	pCounter++
	pc = Prog[pCounter]
	d.Val = (*pc().(*Symbol)).F(d.Val)
	d.Push()
	return nil
}

func Le() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	if d1.Val <= d2.Val {
		d1.Val = 1
	} else {
		d1.Val = 0
	}
	d1.Push()
	return nil
}

func Gt() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	if d1.Val > d2.Val {
		d1.Val = 1
	} else {
		d1.Val = 0
	}
	d1.Push()
	return nil
}

func Lt() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	if d1.Val < d2.Val {
		d1.Val = 1
	} else {
		d1.Val = 0
	}
	d1.Push()
	return nil
}

func Eq() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	if d1.Val == d2.Val {
		d1.Val = 1
	} else {
		d1.Val = 0
	}
	d1.Push()
	return nil
}

func Ge() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	if d1.Val >= d2.Val {
		d1.Val = 1
	} else {
		d1.Val = 0
	}
	d1.Push()
	return nil
}

func Ne() interface{} {
	d2 := Pop().(Datum)
	d1 := Pop().(Datum)
	if d1.Val != d2.Val {
		d1.Val = 1
	} else {
		d1.Val = 0
	}
	d1.Push()
	return nil
}

func And() interface{} {
	// 	d2 := Pop().(Datum)
	// 	d1 := Pop().(Datum)
	// 	if d1.Val && d2.Val {
	// 		d1.Val = 1
	// 	} else {
	// 		d1.Val = 0
	// 	}
	// 	d1.Push()
	return nil
}

//
func Or() interface{} {
	// 	d2 := Pop().(Datum)
	// 	d1 := Pop().(Datum)
	// 	if d1.Val || d2.Val {
	// 		d1.Val = 1
	// 	} else {
	// 		d1.Val = 0
	// 	}
	// 	d1.Push()
	return nil
}

//
func Not() interface{} {
	// 	d1 := Pop().(Datum)
	// 	if !d1.Val {
	// 		d1.Val = 1
	// 	} else {
	// 		d1.Val = 0
	// 	}
	// 	d1.Push()
	return nil
}

func Whilecode() interface{} {
	fmt.Println("In Whilecode: ", pc)
	fmt.Println("WhileBodyCounter, WhileNextCounter, CondCounter: ", WhileBodyCounter, WhileNextCounter, CondCounter)
	savepCounter := pCounter
	pCounter = CondCounter
	Execute()
	d := Pop().(Datum)
	fmt.Println("condition data: ", d)
	for d.Val == 1 {
		pCounter = savepCounter
		Execute()
		pCounter = CondCounter
		Execute()
		d = Pop().(Datum)
	}

	return nil
}

func Ifcode() interface{} {
	fmt.Println("In Ifcode: ", pc)
	fmt.Println("IfBodyCounter, IfElseCounter, CondCounter, IfNextCounter: ", IfBodyCounter, IfElseCounter, CondCounter, IfNextCounter)
	pCounter = CondCounter
	Execute()
	d := Pop().(Datum)
	fmt.Println("condition data: ", d)
	if d.Val == 1 {
		pCounter = IfBodyCounter
		Execute()
	} else {
		pCounter = IfElseCounter
		Execute()
	}
	pCounter = IfNextCounter
	Execute()
	pCounter--
	return nil
}

func PrExpr() interface{} {
	fmt.Println("In PrExpr: ", pc)
	d := Pop().(Datum)
	fmt.Printf("\t%v\n", d.Val)
	return nil
}
