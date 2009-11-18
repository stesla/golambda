package golambda

import(
	"fmt";
	"container/list";
)

var numErrors int;

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
