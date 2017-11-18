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
      | list asgn '\n'  {code.Opr(code.Print); code.STOP.Code(); return 1;}
      | list expr '\n'  {code.Opr(code.Print); code.STOP.Code(); return 1;}
      | list error '\n' {fmt.Printf("error occurred")}
      ;

asgn: VAR '=' expr {code.Opr(code.Varpush); code.Val($1); code.Opr(code.Assign)}
      ;

expr:   '('expr')'    {$$ = $2}
      | expr '%' expr {code.Opr(code.Mod)}
      | expr '+' expr {code.Opr(code.Add)}
      | expr '-' expr {code.Opr(code.Sub)}
      | expr '*' expr {code.Opr(code.Mul)}
      | expr '/' expr {code.Opr(code.Div)}
      | expr '^' expr {code.Opr(code.Power)}
      | NUMBER        {code.Opr(code.Constpush); code.Val($1)}
      | '-' expr %prec UNARYMINUS {code.Opr(code.Negate)}
      | VAR {code.Opr(code.Varpush); code.Val($1); code.Opr(code.Eval)}
      | asgn
      | BLTIN '('expr')' {code.Opr(code.Bltin); code.Val($1)}
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
