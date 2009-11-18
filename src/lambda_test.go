package golambda

import (
	"testing";
)

type occursFree struct {
	message string;
	expected bool;
	ident string;
	expression string;
}

func (test occursFree) run(t *testing.T) {
	if ast, ok := ParseString(test.expression); ok {
		expectBoolean(t, test.expected, ast.OccursFree(test.ident), test.message);
	} else {
		t.Errorf("%v: `%v` does not parse", test.message, test.expression);
	}
}

func TestOccursFree(t *testing.T) {
	test(t, []testCase{
		occursFree{"same variable", true, "x", "x"},
		occursFree{"different variable", false, "x", "y"},
		occursFree{"bound", false, "x", "fn x. x"},
		occursFree{"not bound", true, "x", "fn y. x"},
		occursFree{"application", true, "x", "(x y)"},
		occursFree{"bound application", false, "x", "((fn x. x) y)"},
		occursFree{"bound application to free var", true, "x", "((fn x. x) x)"}
	});
}