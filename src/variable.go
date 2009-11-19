package golambda

type Variable struct {
	name string;
}

func (v Variable) String() string {
	return v.name;
}

func (v Variable) OccursFree(ident string) bool {
	return v.name == ident;
}

func (v Variable) Substitute(ident string, subst Expression) Expression {
	if v.name == ident {
		return subst;
	}
	return v;
}

func (v Variable) Reduce() Expression {
	return v;
}
