package golambda

import (
	"fmt";
	"testing";
)

func expectExpression(t *testing.T, expected string, actual fmt.Stringer) {
	if pass, err := testing.MatchString(expected, actual.String()); !pass {
		t.Error(err);
	}
}

func TestVariable(t *testing.T) {
	varX := Variable{"x"};
	expectExpression(t, "x", varX);
}

func TestGroup(t *testing.T) {
	group := Group{ Variable{"x"} };
	expectExpression(t, "(x)", group);
}

func TestApplication(t *testing.T) {
	application := Application{ Variable{"f"}, Variable{"x"} };
	expectExpression(t, "f x", application);
}

func TestAbstraction(t *testing.T) {
	abstraction := Abstraction{ "x", Variable{"y"} };
	expectExpression(t, "fn x. y", abstraction);
}