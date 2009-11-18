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

type substTest struct {
	message string;
	expected string;
	input string;
	ident string;
	subst string;
}

func (test substTest) run(t *testing.T) {
	if ast_input, ok := ParseString(test.input); ok {
		if ast_subst, ok := ParseString(test.subst); ok {
			actual := ast_input.Substitute(test.ident, ast_subst);
			expectString(t, test.expected, actual.String(), test.message);
		} else {
			t.Errorf("%v: `%v` does not parse", test.message, test.subst);
		}
	} else {
		t.Errorf("%v: `%v` does not parse", test.message, test.input);
	}
}

func TestSubstitute(t *testing.T) {
	test(t, []testCase{
		substTest{"matching variable", "y", "x", "x", "y"},
		substTest{"substitute an abstraction", "fn x. x", "x", "x", "fn x. x"},
		substTest{"non-matching variable", "x", "x", "y", "z"},
		substTest{"application", "y z", "x z", "x", "y"},
		substTest{"bound var", "fn x. x", "fn x. x", "x", "y"},
		substTest{"bound var not free in substitution", "fn x. z", "fn x. y", "y", "z"},
		substTest{"bound var free in substitution", "fn x0. x", "fn x. y", "y", "x"},
		substTest{"group", "(y)", "(x)", "x", "y"}
	});
}

type reduceTest struct {
	message string;
	expected string;
	input string;
}

func (test reduceTest) run(t *testing.T) {
	if ast, ok := ParseString(test.input); ok {
		actual := ast.Reduce();
		expectString(t, test.expected, actual.String(), test.message);
	} else {
		t.Errorf("%v: `%v` does not parse", test.message, test.input);
	}
}

func TestReduce(t *testing.T) {
	test(t, []testCase{
		reduceTest{"variable", "x", "x"},
		reduceTest{"abstraction", "fn x. x", "fn x. x"},
		reduceTest{"application of non-abstraction", "x y", "x y"},
		reduceTest{"application of abstraction", "y", "(fn x. x) y"},
		reduceTest{"higher-order function", "fn x. x", "(fn z. z) (fn x. x)"},
		reduceTest{"sequence", "y", "(fn z. z) (fn x. x) y"},
		reduceTest{"nested", "a", "(fn x y. x y) (fn z. z) a"},
		reduceTest{"several non-abstractions", "x y z", "x y z"},
		reduceTest{"Church if true", "a", "(fn p. p) (fn x y. x) a b"},
		reduceTest{"Church if false", "b", "(fn p. p) (fn x y. y) a b"}
	});
}