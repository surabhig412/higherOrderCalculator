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
  )
%}

%union{
	inst *Inst
  sym *Symbol
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
      | list asgn '\n'  {Opr(Print); STOP.Code(); return 1;}
      | list stmt '\n'  {STOP.Code(); return 1;}
      | list expr '\n'  {Opr(Print); STOP.Code(); return 1;}
      | list error '\n' {fmt.Println("error occurred")}
      ;

asgn: VAR '=' expr {$$ = $3; Opr(Varpush); Val($1); Opr(Assign)}
      ;

stmt:   expr          {fmt.Println("while expr called"); Opr(Print)}
      | PRINT expr   {Opr(PrExpr); $$=$2;}
      | while cond stmt end {
          fmt.Println("while stmt called, $2: ", $2, " $3: ", $3, " $4: ", $4);
          WhileBodyCounter = len(Prog)
          ($3).Code()
          WhileNextCounter = len(Prog)
          if $4 == nil {
            STOP.Code()
          } else {
            ($4).Code()
          }
      }
      | if cond stmt end {
          fmt.Println("if stmt called, $2: ", $2, " $3: ", $3, " $4: ", $4);
          ($3).Code()
          IfNextCounter = len(Prog)
          if $4 == nil {
            STOP.Code()
          } else {
            ($4).Code()
          }
      }
      | if cond stmt end else stmt end {
          fmt.Println("if else stmt called, $2: ", $2, " $3: ", $3, " $4: ", $4, " $6: ", $6, " $7: ", $7);
          ($3).Code()

          ($6).Code()
          IfNextCounter = len(Prog)
          if $7 == nil {
            STOP.Code()
          } else {
            ($7).Code()
          }
      }
      | '{'stmtlist'}'    {$$ = $2}
      ;
cond:   '('expr')'    {fmt.Println("while cond called"); STOP.Code(); $$ = $2; IfBodyCounter = len(Prog)}
      ;

while:  WHILE         {fmt.Println("while called"); $$=Opr(Whilecode); STOP.Code(); STOP.Code(); CondCounter = len(Prog);}
      ;

else:   ELSE          {fmt.Println("else called"); IfElseCounter = len(Prog)}
      ;

if:     IF            {$$=Opr(Ifcode); STOP.Code(); STOP.Code(); STOP.Code(); CondCounter = len(Prog);}
      ;

end:    /* empty */   {STOP.Code()} //$$ = len(Prog)
      ;

stmtlist: /* empty */  {}   //{$$ = len(Prog)}
        | stmtlist '\n'
        | stmtlist stmt
        ;

expr:   '('expr')'    {$$ = $2}
      | expr '%' expr {Opr(Mod)}
      | expr '+' expr {Opr(Add)}
      | expr '-' expr {Opr(Sub)}
      | expr '*' expr {Opr(Mul)}
      | expr '/' expr {Opr(Div)}
      | expr '^' expr {Opr(Power)}
      | NUMBER        {$$=Opr(Constpush); Val($1)}
      | '-' expr %prec UNARYMINUS {$$=$2; Opr(Negate)}
      | VAR {$$=Opr(Varpush); Val($1); Opr(Eval)}
      | asgn
      | BLTIN '('expr')' {$$=$3; Opr(Bltin); Val($1)}
      | expr GT expr   {Opr(Gt)}
      | expr GE expr   {Opr(Ge)}
      | expr LT expr   {Opr(Lt)}
      | expr LE expr   {Opr(Le)}
      | expr EQ expr   {Opr(Eq)}
      | expr NE expr   {Opr(Ne)}
      | expr AND expr  {Opr(And)}
      | expr OR expr   {Opr(Or)}
      | NOT expr       {$$=$2; Opr(Not)}
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
