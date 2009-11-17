package golambda

import (
	"testing";
)

func TestVariable(t *testing.T) {
	expectExpression(t,
		"x",
		Variable{"x"});
}

func TestGroup(t *testing.T) {
	expectExpression(t,
		"(x)",
		Group{ Variable{"x"} });
}

func TestApplication(t *testing.T) {
	expectExpression(t,
		"f x",
		Application{ Variable{"f"}, Variable{"x"} });
}

func TestAbstraction(t *testing.T) {
	expectExpression(t, 
		"fn x. y",
		Abstraction{ "x", Variable{"y"} });
}