//line hoc.y:2
package main

import __yyfmt__ "fmt"

//line hoc.y:2
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

//line hoc.y:15
type yySymType struct {
	yys  int
	inst *Inst
	sym  *Symbol
}

const NUMBER = 57346
const VAR = 57347
const BLTIN = 57348
const UNDEF = 57349
const PRINT = 57350
const WHILE = 57351
const IF = 57352
const ELSE = 57353
const OR = 57354
const AND = 57355
const GT = 57356
const GE = 57357
const LT = 57358
const LE = 57359
const EQ = 57360
const NE = 57361
const UNARYMINUS = 57362
const NOT = 57363

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"VAR",
	"BLTIN",
	"UNDEF",
	"PRINT",
	"WHILE",
	"IF",
	"ELSE",
	"'='",
	"OR",
	"AND",
	"GT",
	"GE",
	"LT",
	"LE",
	"EQ",
	"NE",
	"'%'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"UNARYMINUS",
	"NOT",
	"'^'",
	"'\\n'",
	"'{'",
	"'}'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line hoc.y:121

type Lexer struct {
	s   string
	pos int
}

func (l *Lexer) Lex(lval *yySymType) int {
	var c rune = ' '
	for c == ' ' || c == '\t' {
		if l.pos == len(l.s) {
			return 0
		}
		c = rune(l.s[l.pos])
		l.pos += 1
	}

	_, err := strconv.ParseFloat(string(c), 32)
	if string(c) == "." || err == nil {
		re := regexp.MustCompile("[0-9]+")
		locations := re.FindStringIndex(l.s[l.pos-1:])
		str := re.FindString(l.s[l.pos-1:])
		l.pos += locations[1] - 1
		f, _ := strconv.ParseFloat(str, 64)
		s := &Symbol{Type: NUMBER, Val: f}
		lval.sym = s
		fmt.Println("NUMBER: ", s)
		return NUMBER
	}

	if unicode.IsLetter(c) {
		name := string(c)
		for unicode.IsLetter(rune(l.s[l.pos])) && l.pos != len(l.s) {
			name += string(l.s[l.pos])
			l.pos += 1
		}
		s := Lookup(name)
		if s == nil {
			s = &Symbol{Name: name, Type: UNDEF, Val: 0.0}
		}
		lval.sym = s
		if s.Type == UNDEF {
			return VAR
		}
		return s.Type
	}

	switch string(c) {
	case ">":
		return l.follow("=", GE, GT)
	case "<":
		return l.follow("=", LE, LT)
	case "=":
		return l.follow("=", EQ, int(c))
	case "!":
		return l.follow("=", NE, NOT)
	case "|":
		return l.follow("|", OR, int(c))
	case "&":
		return l.follow("&", AND, int(c))
	case "\n":
		return int(c)
	default:
		return int(c)
	}
	return int(c)
}

func (l *Lexer) follow(expect string, ifyes, ifno int) int {
	nextChar := rune(l.s[l.pos])
	if string(nextChar) == expect {
		l.pos += 1
		return ifyes
	}
	return ifno
}

func (l *Lexer) Error(s string) {
	fmt.Fprintf(os.Stderr, "%s: %s", "HOC", s)
}

func main() {
	Init()
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		yyParse(&Lexer{s: str, pos: 0})
		Execute()
	}
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 262

var yyAct = [...]int{

	72, 4, 46, 41, 36, 64, 20, 5, 22, 23,
	24, 25, 26, 19, 38, 27, 27, 37, 44, 77,
	45, 76, 47, 25, 26, 1, 10, 27, 48, 49,
	50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
	60, 61, 63, 62, 66, 69, 39, 65, 3, 9,
	43, 0, 71, 35, 34, 28, 29, 30, 31, 32,
	33, 22, 23, 24, 25, 26, 40, 74, 27, 23,
	24, 25, 26, 75, 0, 27, 0, 42, 78, 79,
	35, 34, 28, 29, 30, 31, 32, 33, 22, 23,
	24, 25, 26, 0, 0, 27, 0, 0, 0, 0,
	73, 35, 34, 28, 29, 30, 31, 32, 33, 22,
	23, 24, 25, 26, 0, 0, 27, 0, 13, 7,
	15, 70, 8, 17, 18, 34, 28, 29, 30, 31,
	32, 33, 22, 23, 24, 25, 26, 14, 0, 27,
	0, 16, 0, 68, 11, 67, 12, 35, 34, 28,
	29, 30, 31, 32, 33, 22, 23, 24, 25, 26,
	0, 0, 27, 21, 6, 0, 13, 7, 15, 0,
	8, 17, 18, 0, 13, 7, 15, 0, 8, 17,
	18, 0, 0, 0, 0, 14, 0, 0, 0, 16,
	0, 2, 11, 14, 12, 0, 0, 16, 0, 0,
	11, 0, 12, 35, 34, 28, 29, 30, 31, 32,
	33, 22, 23, 24, 25, 26, 0, 0, 27, 28,
	29, 30, 31, 32, 33, 22, 23, 24, 25, 26,
	0, 0, 27, 13, 7, 15, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 14, 0, 0, 0, 16, 0, 0, 0,
	0, 12,
}
var yyPact = [...]int{

	-1000, 162, -1000, -16, -23, 134, -25, 5, 229, -29,
	-29, -1000, 229, -1000, 229, -30, 229, -1000, -1000, -1000,
	-1000, -1000, 229, 229, 229, 229, 229, 229, 229, 229,
	229, 229, 229, 229, 229, 229, -1000, 229, 190, -1000,
	170, 229, 170, 114, 88, -12, 229, -12, 47, -1,
	-1, -12, -12, -12, -13, -13, -13, -13, -13, -13,
	204, 111, 190, -1000, 190, 67, -1000, -1000, -1000, -1000,
	-1000, 40, -1000, -1000, 8, -1000, 170, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 5, 46, 1, 50, 66, 49, 26, 0, 25,
	21,
}
var yyR1 = [...]int{

	0, 9, 9, 9, 9, 9, 9, 2, 3, 3,
	3, 3, 3, 3, 5, 6, 10, 7, 8, 4,
	4, 4, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1,
}
var yyR2 = [...]int{

	0, 0, 2, 3, 3, 3, 3, 3, 1, 2,
	4, 4, 7, 3, 3, 1, 1, 1, 0, 0,
	2, 2, 3, 3, 3, 3, 3, 3, 3, 1,
	2, 1, 1, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 2,
}
var yyChk = [...]int{

	-1000, -9, 29, -2, -3, -1, 2, 5, 8, -6,
	-7, 30, 32, 4, 23, 6, 27, 9, 10, 29,
	29, 29, 21, 22, 23, 24, 25, 28, 15, 16,
	17, 18, 19, 20, 14, 13, 29, 12, -1, -2,
	-5, 32, -5, -4, -1, -1, 32, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -3, -1, -1, -3, 31, 29, -3,
	33, -1, -8, 33, -8, 33, -10, 11, -3, -8,
}
var yyDef = [...]int{

	1, -2, 2, 32, 0, 0, 0, 31, 0, 0,
	0, 19, 0, 29, 0, 0, 0, 15, 17, 3,
	4, 5, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 6, 0, 9, 32,
	0, 0, 0, 0, 0, 30, 0, 42, 23, 24,
	25, 26, 27, 28, 34, 35, 36, 37, 38, 39,
	40, 41, 7, 18, 8, 0, 18, 13, 20, 21,
	22, 0, 10, 14, 11, 33, 0, 16, 18, 12,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	29, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 21, 3, 3,
	32, 33, 24, 22, 3, 23, 3, 25, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 12, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 28, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 30, 3, 31,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	13, 14, 15, 16, 17, 18, 19, 20, 26, 27,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:33
		{
			Opr(Print)
			STOP.Code()
			return 1
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:34
		{
			STOP.Code()
			return 1
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:35
		{
			Opr(Print)
			STOP.Code()
			return 1
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:36
		{
			fmt.Println("error occurred")
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:39
		{
			yyVAL.inst = yyDollar[3].inst
			Opr(Varpush)
			Val(yyDollar[1].sym)
			Opr(Assign)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hoc.y:42
		{
			fmt.Println("while expr called")
			Opr(Print)
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hoc.y:43
		{
			Opr(PrExpr)
			yyVAL.inst = yyDollar[2].inst
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hoc.y:44
		{
			fmt.Println("while stmt called, $2: ", yyDollar[2].inst, " $3: ", yyDollar[3].inst, " $4: ", yyDollar[4].inst)
			WhileBodyCounter = len(Prog)
			(yyDollar[3].inst).Code()
			WhileNextCounter = len(Prog)
			if yyDollar[4].inst == nil {
				STOP.Code()
			} else {
				(yyDollar[4].inst).Code()
			}
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hoc.y:55
		{
			fmt.Println("if stmt called, $2: ", yyDollar[2].inst, " $3: ", yyDollar[3].inst, " $4: ", yyDollar[4].inst)
			(yyDollar[3].inst).Code()
			IfNextCounter = len(Prog)
			if yyDollar[4].inst == nil {
				STOP.Code()
			} else {
				(yyDollar[4].inst).Code()
			}
		}
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line hoc.y:65
		{
			fmt.Println("if else stmt called, $2: ", yyDollar[2].inst, " $3: ", yyDollar[3].inst, " $4: ", yyDollar[4].inst, " $6: ", yyDollar[6].inst, " $7: ", yyDollar[7].inst)
			(yyDollar[3].inst).Code()

			(yyDollar[6].inst).Code()
			IfNextCounter = len(Prog)
			if yyDollar[7].inst == nil {
				STOP.Code()
			} else {
				(yyDollar[7].inst).Code()
			}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:77
		{
			yyVAL.inst = yyDollar[2].inst
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:79
		{
			fmt.Println("while cond called")
			STOP.Code()
			yyVAL.inst = yyDollar[2].inst
			IfBodyCounter = len(Prog)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hoc.y:82
		{
			fmt.Println("while called")
			yyVAL.inst = Opr(Whilecode)
			STOP.Code()
			STOP.Code()
			CondCounter = len(Prog)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hoc.y:85
		{
			fmt.Println("else called")
			IfElseCounter = len(Prog)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hoc.y:88
		{
			yyVAL.inst = Opr(Ifcode)
			STOP.Code()
			STOP.Code()
			STOP.Code()
			CondCounter = len(Prog)
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hoc.y:91
		{
			STOP.Code()
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line hoc.y:94
		{
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:99
		{
			yyVAL.inst = yyDollar[2].inst
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:100
		{
			Opr(Mod)
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:101
		{
			Opr(Add)
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:102
		{
			Opr(Sub)
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:103
		{
			Opr(Mul)
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:104
		{
			Opr(Div)
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:105
		{
			Opr(Power)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hoc.y:106
		{
			yyVAL.inst = Opr(Constpush)
			Val(yyDollar[1].sym)
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hoc.y:107
		{
			yyVAL.inst = yyDollar[2].inst
			Opr(Negate)
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hoc.y:108
		{
			yyVAL.inst = Opr(Varpush)
			Val(yyDollar[1].sym)
			Opr(Eval)
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line hoc.y:110
		{
			yyVAL.inst = yyDollar[3].inst
			Opr(Bltin)
			Val(yyDollar[1].sym)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:111
		{
			Opr(Gt)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:112
		{
			Opr(Ge)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:113
		{
			Opr(Lt)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:114
		{
			Opr(Le)
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:115
		{
			Opr(Eq)
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:116
		{
			Opr(Ne)
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:117
		{
			Opr(And)
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:118
		{
			Opr(Or)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hoc.y:119
		{
			yyVAL.inst = yyDollar[2].inst
			Opr(Not)
		}
	}
	goto yystack /* stack new state and value */
}
