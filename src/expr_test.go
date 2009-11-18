package golambda

import (
	"fmt";
	"testing";
)

type fmtTest struct {
	message string;
	expected string;
	actual fmt.Stringer;
}

func (test fmtTest) run(t *testing.T) {
	expectString(t, test.expected, test.actual.String(), test.message);
}

func TestString(t *testing.T) {
	test(t, []testCase{
		fmtTest{"variable", "x", Variable{"x"}},
		fmtTest{"group", "(x)", Group{ Variable{"x"} }},
		fmtTest{"application", "f x", Application{ Variable{"f"}, Variable{"x"} }},
		fmtTest{"abstraction", "fn x. y", Abstraction{ "x", Variable{"y"} }}
	})
}