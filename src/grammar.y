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

aexprs: aexpr { $$ = makeAexprs($1, list.New()); }
| aexpr aexprs { $$ = makeAexprs($1, $2); }
;

idents: idents IDENT { $$ = makeIdents($1, $2); }
      | IDENT { $$ = makeIdents(list.New(), $1); }
      ;

%%
