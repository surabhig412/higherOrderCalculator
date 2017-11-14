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
    "math"
    "github.com/surabhig412/hoc/symbol"
  )
%}

%union{
	val float64
  sym *symbol.Symbol
}
%type <val> expr asgn
%token <val> NUMBER
%token <sym> VAR BLTIN UNDEF
%right '='
%left '%'
%left '+' '-'
%left '*' '/'
%left UNARYMINUS
%left UNARYPLUS
%right '^'
%%
list:   /* empty */
      | list '\n'
      | list asgn '\n'
      | list expr '\n'  {fmt.Printf("%v\n", $2)}
      | list error '\n' {fmt.Printf("error occurred")}
      ;

asgn: VAR '=' expr {$1.Val = $3; $$ = $1.Val; $1.Type = VAR;}
      ;

expr:   '('expr')'    {$$ = $2}
      | expr '%' expr {$$ = math.Mod($1, $3)}
      | expr '+' expr {$$ = $1 + $3}
      | expr '-' expr {$$ = $1 - $3}
      | expr '*' expr {$$ = $1 * $3}
      | expr '/' expr {
              if($3 == 0.0) {
                log.Fatalf("division by zero")
              }
              $$ = $1 / $3}
      | expr '^' expr {$$ = math.Pow($1, $3)}
      | NUMBER        {$$ = $1}
      | '-' expr %prec UNARYMINUS {$$ = -$2}
      | '+' expr %prec UNARYPLUS {$$ = $2}
      | VAR {
          if $1.Type == UNDEF {
          log.Fatalf("undefined variable: ", $1.Name)
          }
          $$ = $1.Val
      }
      | asgn
      | BLTIN '('expr')' {$$ = (($1.F))($3)}
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
    lval.val = f
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
  symbol.Init(VAR, BLTIN)
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
