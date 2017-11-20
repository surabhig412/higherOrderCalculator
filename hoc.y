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
%type <inst> expr asgn stmt stmtlist cond while if end
%token <sym> NUMBER VAR BLTIN UNDEF PRINT WHILE IF ELSE
%right '='
%left OR
%left AND
%left GT GE LT LE EQ NE
%left '%'
%left '+' '-'
%left '*' '/'
%left UNARYMINUS NOT
%right '^'
%%
list:   /* empty */
      | list '\n'
      | list asgn '\n'  {code.Opr(code.Print); code.STOP.Code(); return 1;}
      | list stmt '\n'  {code.STOP.Code(); return 1;}
      | list expr '\n'  {code.Opr(code.Print); code.STOP.Code(); return 1;}
      | list error '\n' {fmt.Println("error occurred")}
      ;

asgn: VAR '=' expr {$$ = $3; code.Opr(code.Varpush); code.Val($1); code.Opr(code.Assign)}
      ;

stmt:   expr          {fmt.Println("while expr called"); code.Opr(code.Print)}
      | PRINT expr   {code.Opr(code.PrExpr); $$=$2;}
      | while cond stmt end {
          fmt.Println("while stmt called, $2: ", $2, " $3: ", $3, " $4: ", $4);
          code.WhileBodyCounter = len(code.Prog)
          ($3).Code()
          code.WhileNextCounter = len(code.Prog)
          if $4 == nil {
            code.STOP.Code()
          } else {
            ($4).Code()
          }
      }
      | if cond stmt end {
          fmt.Println("if stmt called, $2: ", $2, " $3: ", $3, " $4: ", $4);
          ($3).Code()
          code.IfNextCounter = len(code.Prog)
          if $4 == nil {
            code.STOP.Code()
          } else {
            ($4).Code()
          }
      }
      | if cond stmt end else stmt end {
          fmt.Println("if else stmt called, $2: ", $2, " $3: ", $3, " $4: ", $4, " $6: ", $6, " $7: ", $7);
          ($3).Code()

          ($6).Code()
          code.IfNextCounter = len(code.Prog)
          if $7 == nil {
            code.STOP.Code()
          } else {
            ($7).Code()
          }
      }
      | '{'stmtlist'}'    {$$ = $2}
      ;
cond:   '('expr')'    {fmt.Println("while cond called"); code.STOP.Code(); $$ = $2; code.IfBodyCounter = len(code.Prog)}
      ;

while:  WHILE         {fmt.Println("while called"); $$=code.Opr(code.Whilecode); code.STOP.Code(); code.STOP.Code(); code.CondCounter = len(code.Prog);}
      ;

else:   ELSE          {fmt.Println("else called"); code.IfElseCounter = len(code.Prog)}
      ;

if:     IF            {$$=code.Opr(code.Ifcode); code.STOP.Code(); code.STOP.Code(); code.STOP.Code(); code.CondCounter = len(code.Prog);}
      ;

end:    /* empty */   {code.STOP.Code()} //$$ = len(code.Prog)
      ;

stmtlist: /* empty */  {}   //{$$ = len(code.Prog)}
        | stmtlist '\n'
        | stmtlist stmt
        ;

expr:   '('expr')'    {$$ = $2}
      | expr '%' expr {code.Opr(code.Mod)}
      | expr '+' expr {code.Opr(code.Add)}
      | expr '-' expr {code.Opr(code.Sub)}
      | expr '*' expr {code.Opr(code.Mul)}
      | expr '/' expr {code.Opr(code.Div)}
      | expr '^' expr {code.Opr(code.Power)}
      | NUMBER        {$$=code.Opr(code.Constpush); code.Val($1)}
      | '-' expr %prec UNARYMINUS {$$=$2; code.Opr(code.Negate)}
      | VAR {$$=code.Opr(code.Varpush); code.Val($1); code.Opr(code.Eval)}
      | asgn
      | BLTIN '('expr')' {$$=$3; code.Opr(code.Bltin); code.Val($1)}
      | expr GT expr   {code.Opr(code.Gt)}
      | expr GE expr   {code.Opr(code.Ge)}
      | expr LT expr   {code.Opr(code.Lt)}
      | expr LE expr   {code.Opr(code.Le)}
      | expr EQ expr   {code.Opr(code.Eq)}
      | expr NE expr   {code.Opr(code.Ne)}
      | expr AND expr  {code.Opr(code.And)}
      | expr OR expr   {code.Opr(code.Or)}
      | NOT expr       {$$=$2; code.Opr(code.Not)}
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
    fmt.Println("NUMBER: ", s)
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

  switch string(c) {
  case ">": return l.follow("=", GE, GT);
  case "<": return l.follow("=", LE, LT);
  case "=": return l.follow("=", EQ, int(c));
  case "!": return l.follow("=", NE, NOT);
  case "|": return l.follow("|", OR, int(c));
  case "&": return l.follow("&", AND, int(c));
  case "\n": return int(c);
  default:  return int(c);
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
  symbol.Init(VAR, BLTIN, UNDEF, IF, ELSE, WHILE, PRINT)
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
