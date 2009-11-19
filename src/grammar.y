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
	list *list.List;
}

%type <expr> expr aexpr
%type <list> idents aexprs

%token <ident> IDENT
%token FN

%%

expr: aexprs { $$ = makeApplication($1); }
| FN idents '.' expr { $$ = makeAbstraction($2, $4); }
;

aexpr: IDENT { $$ = Variable{$1}; }
| '(' expr ')' { $$ = Group{$2}; }
;

aexprs: aexpr { $$ = push($1, list.New()); }
| aexpr aexprs { $$ = push($1, $2); }
;

idents: idents IDENT { $$ = push($2, $1); }
      | IDENT { $$ = push($1, list.New()); }
      ;

%%
