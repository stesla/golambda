%{
package golambda

import(
	"fmt";
	"container/list";
	"utf8";
)

var line string;
var linep int;
var lookahead int;
var numErrors int;
%}

%union {
	expr Expression;
	ident string;
	idents *list.List;
}

%type <expr> expr
%type <idents> idents

%token <ident> IDENT
%token FN

%%

expr: IDENT { $$ = Variable{$1}; }
    | FN idents '.' expr { $$ = makeAbstraction($2, $4); }
	| expr expr { $$ = Application{$1,$2}; }
    | '(' expr ')' { $$ = $2; }
    ;

idents: idents IDENT { $$ = makeIdents($1, $2); }
      | IDENT { $$ = makeIdents(list.New(), $1); }
      ;

%%

func makeAbstraction(idents *list.List, body Expression) Expression {
	for idents.Len() > 0 {
		elt := idents.Front();
		idents.Remove(elt);
		ident,_ := elt.Value.(string);
		body = Abstraction{ident, body};
	}
	return body;
}

func makeIdents(idents *list.List, ident string) *list.List {
	idents.PushFront(ident);
	return idents;
}

func getrune() (result int) {
	var n int;

	if linep >= len(line) {
		return 0;
	}
	result,n = utf8.DecodeRuneInString(line[linep:len(line)]);
	linep += n;
	if result == '\n' {
		result = 0;
	}
	return result;
}

func Lex() int {
	c := lookahead;
	lookahead = getrune();
	switch true {
	case c == ' ' || c == '\t':
		return Lex();
	case c == 'f' && lookahead == 'n':
		lookahead = getrune();
		return FN;
	case isIdent(c):
		str := string(c);
		for isIdent(lookahead) {
			str += string(lookahead);
			lookahead = getrune();
		}
		yylval.ident = str;
		return IDENT;
	}
	return c;
}

func isIdent(c int) bool {
	return
		('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z') ||
		('0' <= c && c <= '9') ||
		c == '_';
}

func Error(s string, v ...)
{
	numErrors += 1;
	fmt.Printf(s, v);
	fmt.Printf("\n");
}

func ParseString(input string) (e Expression, ok bool) {
	line = input;
	linep = 0;
	lookahead = getrune();
	if Parse() == 0 {
		e = YYVAL.expr;
		ok = true;
    }
	return;
}
