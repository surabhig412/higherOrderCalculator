%{
  package main
  import(
    "fmt"
    "bufio"
    "os"
    "log"
    "io"
    "strconv"
    "regexp"
    "unicode"
    "github.com/surabhig412/hoc/symbol"
    "github.com/surabhig412/hoc/code"
  )
%}

%union{
	inst *code.Inst
  sym *symbol.Symbol
}
%type <inst> expr asgn
%token <sym> NUMBER VAR BLTIN UNDEF
%right '='
%left '%'
%left '+' '-'
%left '*' '/'
%left UNARYMINUS
%right '^'
%%
list:   /* empty */
      | list '\n'
      | list asgn '\n'  {(code.Inst(code.Print)).Code(); code.STOP.Code(); return 1;}
      | list expr '\n'  {(code.Inst(code.Print)).Code(); code.STOP.Code(); return 1;}
      | list error '\n' {fmt.Printf("error occurred")}
      ;

asgn: VAR '=' expr {(code.Inst(code.Varpush)).Code(); s := $1; (code.Inst(func()interface{}{return s})).Code(); (code.Inst(code.Assign)).Code()}
      ;

expr:   '('expr')'    {$$ = $2}
      | expr '%' expr {(code.Inst(code.Mod)).Code()}
      | expr '+' expr {(code.Inst(code.Add)).Code()}
      | expr '-' expr {(code.Inst(code.Sub)).Code()}
      | expr '*' expr {(code.Inst(code.Mul)).Code()}
      | expr '/' expr {(code.Inst(code.Div)).Code()}
      | expr '^' expr {(code.Inst(code.Power)).Code()}
      | NUMBER        {(code.Inst(code.Constpush)).Code(); s := $1; (code.Inst(func()interface{}{return s})).Code()}
      | '-' expr %prec UNARYMINUS {(code.Inst(code.Negate)).Code()}
      | VAR {(code.Inst(code.Varpush)).Code(); s := $1; (code.Inst(func()interface{}{return s})).Code(); (code.Inst(code.Eval)).Code()}
      | asgn
      | BLTIN '('expr')' {(code.Inst(code.Bltin)).Code(); s := $1; (code.Inst(func()interface{}{return s})).Code()}
      ;
%%

type Lexer struct {
  s string
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
    s := &symbol.Symbol{Type: NUMBER, Val: f}
    lval.sym = s
    return NUMBER
  }

  if unicode.IsLetter(c) {
    name := string(c)
    for unicode.IsLetter(rune(l.s[l.pos])) && l.pos != len(l.s) {
      name += string(l.s[l.pos])
      l.pos += 1
    }
    s := symbol.Lookup(name)
    if s == nil {
      s = &symbol.Symbol{Name: name, Type: UNDEF, Val: 0.0}
    }
    lval.sym = s
    if s.Type == UNDEF {
      return VAR
    }
    return s.Type
  }
  return int(c)
}

func (l *Lexer) Error(s string) {
  fmt.Fprintf(os.Stderr, "%s: %s", "HOC", s)
}

func main() {
  symbol.Init(VAR, BLTIN, UNDEF)
  reader := bufio.NewReader(os.Stdin)
  for {
  str, err := reader.ReadString('\n')
  if err == io.EOF {
			break
	} else if err != nil {
			log.Fatal(err)
	}
  yyParse(&Lexer{s: str, pos: 0})
  code.Execute()
  }
}
