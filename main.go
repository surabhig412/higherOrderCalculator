package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var lineNo int
var yylval float64

type Token int

const (
	INVALID Token = iota
	NUMBER
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	OPEN_PARENS
	CLOSE_PARENS
	NEWLINE
)

type Buffer struct {
	token Token
	size  int
}
type Parser struct {
	reader *bufio.Reader
	buf    Buffer
}

func NewParser(reader *bufio.Reader) *Parser {
	return &Parser{reader: reader}
}

func (p *Parser) Parse() {
	for {
		rune, _, err := p.reader.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if string(rune) == " " || string(rune) == "\t" {
			continue
		}
		token := Lexer(rune)
		fmt.Println("Parsed string: ", token)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	parser := NewParser(reader)
	parser.Parse()
}

// lexical analyser
func Lexer(r rune) Token {
	num, err := strconv.ParseFloat(string(r), 32)
	if string(r) == "." || err == nil {
		yylval = num
		fmt.Println("yylval : ", yylval)
		return NUMBER
	}
	if r == '\n' {
		lineNo++
		return NEWLINE
	}
	switch string(r) {
	case "+":
		return PLUS
	case "-":
		return MINUS
	case "*":
		return MULTIPLY
	case "/":
		return DIVIDE
	case "(":
		return OPEN_PARENS
	case ")":
		return CLOSE_PARENS
	}
	return INVALID
}

func yyerror(s string) {
	fmt.Fprintf(os.Stderr, "%s: %s near line %d\n", "HOC", s, lineNo)
}
