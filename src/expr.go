package golambda

import "fmt";

type Expression interface {
	fmt.Stringer;
	OccursFree(ident string) bool;
	Reduce() Expression;
	Substitute(ident string, subst Expression) Expression;
}
