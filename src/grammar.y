%{
package golambda

import(
	"fmt";
	"container/list";
)
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
