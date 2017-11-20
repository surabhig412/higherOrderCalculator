package main

import "math"

var consts = map[string]float64{
	"PI":    3.14159265358979323846,
	"E":     2.71828182845904523536,
	"GAMMA": 0.57721566490153286060,  // Euler’s constant
	"DEG":   57.29577951308232087680, // degrees/radian
	"PHI":   1.61803398874989484820,  // golden ratio
}

var keywords = map[string]func() int{
	"if":    func() int { return IF },
	"else":  func() int { return ELSE },
	"while": func() int { return WHILE },
	"print": func() int { return PRINT },
}

func sin(x float64) float64 {
	return math.Sin(x)
}

func cos(x float64) float64 {
	return math.Cos(x)
}

func atan(x float64) float64 {
	return math.Atan(x)
}

func elog(x float64) float64 {
	return math.Log(x)
}

func log10(x float64) float64 {
	return math.Log10(x)
}

func exp(x float64) float64 {
	return math.Exp(x)
}

func sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func fabs(x float64) float64 {
	return math.Abs(x)
}

var builtins = map[string](func(float64) float64){
	"sin":   sin,
	"cos":   cos,
	"atan":  atan,
	"log":   elog,
	"log10": log10,
	"exp":   exp,
	"sqrt":  sqrt,
	"abs":   fabs,
}

type Symbol struct {
	Name string
	Type int
	Val  float64
	F    func(float64) float64
}

var symMap map[string]Symbol

func Lookup(name string) *Symbol {
	value, ok := symMap[name]
	if ok {
		return &value
	}
	return nil
}

func (symbol *Symbol) Install() {
	symMap[symbol.Name] = *symbol
}

func Init() {
	symMap = make(map[string]Symbol, 100)
	for key, value := range consts {
		s := &Symbol{Name: key, Type: VAR, Val: value}
		s.Install()
	}
	for key, value := range builtins {
		s := &Symbol{Name: key, Type: BLTIN, F: value}
		s.Install()
	}
	for key, value := range keywords {
		s := &Symbol{Name: key, Type: value()}
		s.Install()
	}
}
