//line hoc.y:2
package main

import __yyfmt__ "fmt"

//line hoc.y:2
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var mem [26]float64

//line hoc.y:18
type yySymType struct {
	yys   int
	val   float64
	index int
}

const NUMBER = 57346
const VAR = 57347
const UNARYMINUS = 57348
const UNARYPLUS = 57349

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"VAR",
	"'='",
	"'%'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"UNARYMINUS",
	"UNARYPLUS",
	"'\\n'",
	"'('",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line hoc.y:55

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
		lval.val = f
		return NUMBER
	}

	if unicode.IsLower(c) {
		lval.index = int(c - 'a')
		return VAR
	}
	return int(c)
}

func (l *Lexer) Error(s string) {
	fmt.Fprintf(os.Stderr, "%s: %s", "HOC", s)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		yyParse(&Lexer{s: str, pos: 0})
	}

}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 52

var yyAct = [...]int{

	3, 11, 12, 13, 14, 15, 17, 16, 18, 19,
	26, 20, 21, 22, 23, 24, 25, 4, 1, 6,
	9, 27, 0, 8, 7, 0, 0, 6, 9, 2,
	5, 8, 7, 11, 12, 13, 14, 15, 5, 0,
	10, 11, 12, 13, 14, 15, 12, 13, 14, 15,
	14, 15,
}
var yyPact = [...]int{

	-1000, 15, -1000, 26, -7, 23, -1000, 23, 23, 5,
	-1000, 23, 23, 23, 23, 23, -1000, -6, -1000, -1000,
	23, 38, 40, 40, -1000, -1000, -1000, 34,
}
var yyPgo = [...]int{

	0, 0, 18,
}
var yyR1 = [...]int{

	0, 2, 2, 2, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1,
}
var yyR2 = [...]int{

	0, 0, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 1, 2, 2, 1, 3,
}
var yyChk = [...]int{

	-1000, -2, 14, -1, 2, 15, 4, 9, 8, 5,
	14, 7, 8, 9, 10, 11, 14, -1, -1, -1,
	6, -1, -1, -1, -1, -1, 16, -1,
}
var yyDef = [...]int{

	1, -2, 2, 0, 0, 0, 11, 0, 0, 14,
	3, 0, 0, 0, 0, 0, 4, 0, 12, 13,
	0, 6, 7, 8, 9, 10, 5, 15,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 7, 3, 3,
	15, 16, 10, 8, 3, 9, 3, 11, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 6,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 12, 13,
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
		//line hoc.y:34
		{
			fmt.Printf("%v\n", yyDollar[2].val)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:35
		{
			fmt.Printf("error occurred")
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:37
		{
			yyVAL.val = yyDollar[2].val
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:38
		{
			yyVAL.val = math.Mod(yyDollar[1].val, yyDollar[3].val)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:39
		{
			yyVAL.val = yyDollar[1].val + yyDollar[3].val
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:40
		{
			yyVAL.val = yyDollar[1].val - yyDollar[3].val
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:41
		{
			yyVAL.val = yyDollar[1].val * yyDollar[3].val
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:42
		{
			if yyDollar[3].val == 0.0 {
				log.Fatalf("division by zero")
			}
			yyVAL.val = yyDollar[1].val / yyDollar[3].val
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hoc.y:47
		{
			yyVAL.val = yyDollar[1].val
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hoc.y:48
		{
			yyVAL.val = -yyDollar[2].val
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line hoc.y:49
		{
			yyVAL.val = yyDollar[2].val
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line hoc.y:50
		{
			yyVAL.val = mem[yyDollar[1].index]
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line hoc.y:51
		{
			yyVAL.val = mem[yyDollar[1].index]
			mem[yyDollar[1].index] = yyDollar[3].val
		}
	}
	goto yystack /* stack new state and value */
}
