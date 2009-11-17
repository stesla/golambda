package golambda

import (
	"testing";
)

func TestVariable(t *testing.T) {
	expectFmt(t,
		"x",
		Variable{"x"});
}

func TestGroup(t *testing.T) {
	expectFmt(t,
		"(x)",
		Group{ Variable{"x"} });
}

func TestApplication(t *testing.T) {
	expectFmt(t,
		"f x",
		Application{ Variable{"f"}, Variable{"x"} });
}

func TestAbstraction(t *testing.T) {
	expectFmt(t, 
		"fn x. y",
		Abstraction{ "x", Variable{"y"} });
}