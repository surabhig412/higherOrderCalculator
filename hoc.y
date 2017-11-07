%{
  package main
  import(
    "fmt"
    "bufio"
    "os"
    "log"
    "io"
    "strconv"
  )
%}
%union{
	val int
}
%type <val> expr
%token <val> NUMBER
%left '%'
%left '+' '-'
%left '*' '/'
%left UNARYMINUS
%left UNARYPLUS
%%
list:   /* empty */
      | list '\n'
      | list expr '\n'  {fmt.Printf("%d\n", $2)}
      ;
expr:   '('expr')'    {$$ = $2}
      | expr '%' expr {$$ = $1 % $3}
      | expr '+' expr {$$ = $1 + $3}
      | expr '-' expr {$$ = $1 - $3}
      | expr '*' expr {$$ = $1 * $3}
      | expr '/' expr {$$ = $1 / $3}
      | NUMBER        {$$ = $1}
      | '-' expr %prec UNARYMINUS {$$ = -$2}
      | '+' expr %prec UNARYPLUS {$$ = $2}
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
    lval.val = int(c) - '0'
    return NUMBER
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
