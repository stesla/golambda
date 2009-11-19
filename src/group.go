package golambda

type Group struct {
	inner Expression;
}

func (g Group) String() string {
	return "(" + g.inner.String() + ")";
}

func (g Group) OccursFree(ident string) bool {
	return g.inner.OccursFree(ident);
}

func (g Group) Substitute(ident string, subst Expression) Expression {
	return Group{g.inner.Substitute(ident, subst)};
}

func (g Group) Reduce() Expression {
	return g.inner.Reduce();
}
