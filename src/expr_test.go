package golambda

import (
	"testing";
)

func TestString(t *testing.T) {
	test(t, []testCase{
		fmtTest{"variable", "x", Variable{"x"}},
		fmtTest{"group", "(x)", Group{ Variable{"x"} }},
		fmtTest{"application", "f x", Application{ Variable{"f"}, Variable{"x"} }},
		fmtTest{"abstraction", "fn x. y", Abstraction{ "x", Variable{"y"} }}
	})
}