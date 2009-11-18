package golambda

import (
	"fmt";
)

type Expression interface {
	fmt.Stringer;
	OccursFree(ident string) bool;
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

type Group struct {
	inner Expression;
}

func (g Group) String() string {
	return "(" + g.inner.String() + ")";
}

func (g Group) OccursFree(ident string) bool {
	return g.inner.OccursFree(ident);
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
