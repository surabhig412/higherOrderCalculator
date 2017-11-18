package code

import (
	"fmt"
	"log"
	"math"

	"github.com/surabhig412/hoc/symbol"
)

type Datum struct {
	Val float64
	Sym symbol.Symbol
}

type Inst func() interface{}

var stack []Datum
var prog []Inst
var pc Inst
var pCounter int = 0

var STOP Inst = nil

func (d Datum) Push() {
	stack = append(stack, d)
}

func Pop() interface{} {
	length := len(stack)
	if length == 0 {
		log.Fatalf("stack underflow")
	}
	poppedDatum := stack[length-1]
	stack = stack[:length-1]
	return poppedDatum
}

func (inst Inst) Code() Inst {
	var oldProg Inst
	if len(prog) != 0 {
		oldProg = prog[len(prog)-1]
	}
	prog = append(prog, inst)
	return oldProg
}

func Execute() {
	pc = prog[pCounter]
	if pc == nil {
		pCounter++
		pc = prog[pCounter]
	}
	for pc != nil {
		pc()
		pCounter++
		pc = prog[pCounter]
	}
}

func Constpush() interface{} {
	pCounter++
	pc = prog[pCounter]
	d := &Datum{Sym: *pc().(*symbol.Symbol)}
	d.Val = d.Sym.Val
	d.Push()
	return nil
}

func Varpush() interface{} {
	pCounter++
	pc = prog[pCounter]
	d := &Datum{Sym: *pc().(*symbol.Symbol)}
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
	if d.Sym.Type == symbol.UNDEF {
		log.Fatalf("undefined variable ", d.Sym.Name)
	}
	d.Val = d.Sym.Val
	d.Push()
	return nil
}

func Assign() interface{} {
	d1 := Pop().(Datum)
	d2 := Pop().(Datum)
	if d1.Sym.Type != symbol.VAR && d1.Sym.Type != symbol.UNDEF {
		log.Fatalf("assignment to non-variable ", d1.Sym.Name)
	}
	d1.Sym.Val = d2.Val
	d1.Sym.Type = symbol.VAR
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
	pc = prog[pCounter]
	d.Val = (*pc().(*symbol.Symbol)).F(d.Val)
	d.Push()
	return nil
}