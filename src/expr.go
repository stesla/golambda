package golambda

import (
	"fmt";
)

var unique_id int;

type Expression interface {
	fmt.Stringer;
	OccursFree(ident string) bool;
	Reduce() Expression;
	Substitute(ident string, subst Expression) Expression;
}

type Abstraction struct {
	variable string;
	body Expression;
}

func (a Abstraction) String() string {
	return "fn " + a.variable + ". " + a.body.String();
}

func (a Abstraction) OccursFree(ident string) bool {
	if a.variable == ident {
		return false;
	}
	return a.body.OccursFree(ident);
}

func (a Abstraction) Substitute(ident string, subst Expression) (result Expression) {
	if a.variable == ident {
		result = a;
	} else if subst.OccursFree(a.variable) {
		newVar := fmt.Sprintf("%v%v", a.variable, unique_id);
		unique_id++;
		body := a.body.Substitute(newVar, Variable{ident});
		result = Abstraction{newVar, body.Substitute(ident, subst)};
	} else {
		result = Abstraction{a.variable, a.body.Substitute(ident, subst)};
	}
	return;
}

func (a Abstraction) Reduce() Expression {
	return a;
}

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
