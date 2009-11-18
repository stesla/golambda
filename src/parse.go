package golambda

import(
	"fmt";
	"container/list";
)

var numErrors int;

func listPopFront(list *list.List) interface{} {
	elt := list.Front();
	list.Remove(elt);
	return elt.Value;
}

func makeAbstraction(idents *list.List, body Expression) Expression {
	for idents.Len() > 0 {
		ident,_ := listPopFront(idents).(string);
		body = Abstraction{ident, body};
	}
	return body;
}

func makeAexprs(expr Expression, exprs *list.List) *list.List {
	exprs.PushFront(expr);
	return exprs;
}

func makeApplication(exprs *list.List) (result Expression) {
	switch len := exprs.Len(); true {
	case len >= 2:
		x1,_ := listPopFront(exprs).(Expression);
		x2,_ := listPopFront(exprs).(Expression);
		exprs.PushFront(Application{x1,x2});
		result = makeApplication(exprs);
	case len == 1:
		result,_ = listPopFront(exprs).(Expression);
	}
	return;
}

func makeIdents(idents *list.List, ident string) *list.List {
	idents.PushFront(ident);
	return idents;
}

func Error(s string, v ...) {
	numErrors += 1;
	fmt.Printf(s, v);
	fmt.Printf("\n");
}

func ParseString(input string) (result Expression, ok bool) {
	line = input;
	linep = 0;
	lookahead = getrune();
	if Parse() == 0 {
		result = YYVAL.expr;
		ok = true;
	}
	return;
}
