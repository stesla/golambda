package golambda

import (
	"fmt";
)

type Expression interface {
	fmt.Stringer;
}

type Abstraction struct {
	variable string;
	body Expression;
}

func (a Abstraction) String() string {
	return "fn " + a.variable + ". " + a.body.String();
}

type Application struct {
	function Expression;
	argument Expression;
}

func (a Application) String() string {
	return a.function.String() + " " + a.argument.String();
}

type Group struct {
	inner Expression;
}

func (g Group) String() string {
	return "(" + g.inner.String() + ")";
}

type Variable struct {
	name string;
}

func (v Variable) String() string {
	return v.name;
}