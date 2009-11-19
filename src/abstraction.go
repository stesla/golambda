package golambda

import "fmt";

var unique_id int;

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
