package golambda

type Application struct {
	function Expression;
	argument Expression;
}

func (a Application) String() string {
	return a.function.String() + " " + a.argument.String();
}

func (a Application) OccursFree(ident string) bool {
	return a.function.OccursFree(ident) || a.argument.OccursFree(ident); 
}

func (a Application) Substitute(ident string, subst Expression) Expression {
	return Application{a.function.Substitute(ident, subst), a.argument.Substitute(ident, subst)};
}

func (a Application) Reduce() Expression {
	switch f := a.function.(type) {
	case Group:
		return Application{f.inner, a.argument}.Reduce();
	case Abstraction:
		expr := f.body.Substitute(f.variable, a.argument);
		return expr.Reduce();
	case Application:
		if _,isVar := f.function.(Variable); !isVar {
			return Application{f.Reduce(), a.argument}.Reduce();
		}
	}
	return a;
}
