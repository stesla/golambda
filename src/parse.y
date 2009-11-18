%{
package golambda

import(
	"fmt";
	"utf8";
)

var line string;
var linep int;
var numErrors int;
%}

%union {
	expr Expression;
	ident string;
}

%type <expr> expr
%type <ident> 'x'

%%

expr: 'x' { $$ = Variable{$1}; }

%%

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
	c := getrune();
	yylval.ident = string(c);
	return c;
}

func
Error(s string, v ...)
{
	numErrors += 1;
	fmt.Printf(s, v);
	fmt.Printf("\n");
}

func ParseString(input string) (e Expression, ok bool) {
	line = input;
	linep = 0;
    if Parse() == 0 {
		e = YYVAL.expr;
		ok = true;
	}
	return;
}
