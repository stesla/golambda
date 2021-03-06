package golambda

import (
	"testing";
)

type parseTest struct {
	message string;
	input string;
	output string;
}

func (test parseTest) run(t *testing.T) {
	ast, ok := ParseString(test.input);
	if ok {
		expectString(t, test.output, ast.String(), test.message);
	} else {
		t.Errorf("%v: `%v` does not parse", test.message, test.input);
	}
}

func TestParse(t *testing.T) {
	test(t, []testCase{
		parseTest{"ident - single chart", "x", "x"},
		parseTest{"ident - lowers", "foo", "foo"},
		parseTest{"ident - all", "aZ_1", "aZ_1"},
		parseTest{"abstraction", "fn x. y", "fn x. y"},
		parseTest{"application", "f x", "f x"},
		parseTest{"group", "(x)", "(x)"},
		parseTest{"Y combinator",
            "fn g. (fn x. g (x x)) (fn x. g (x x))",
            "fn g. (fn x. g (x x)) (fn x. g (x x))"},
		parseTest{"var applied to abstraction", "x (fn y. y)", "x (fn y. y)"},
		parseTest{"currying", "fn f x. f x", "fn f. fn x. f x"}
	});
}