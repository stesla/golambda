package golambda

import(
	"fmt";
	"container/list";
)

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

func push(item interface{}, l *list.List) *list.List {
	l.PushFront(item);
	return l;
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

func Error(s string, v ...) {
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
