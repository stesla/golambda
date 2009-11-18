package golambda

import (
	"fmt";
	"testing";
)

type testCase interface {
	run(*testing.T);
}

func test(t *testing.T, tests []testCase) {
	for _,test := range tests {
		test.run(t);
	}
}

func expectMatchString(t *testing.T, expected, actual, message string) {
	if pass, _ := testing.MatchString(expected, actual); !pass {
		t.Errorf("%v: expected `%v` to be `%v`", message, actual, expected);
	}
}

type fmtTest struct {
	message string;
	expected string;
	actual fmt.Stringer;
}

func (test fmtTest) run(t *testing.T) {
	expectMatchString(t, test.expected, test.actual.String(), test.message);
}

type parseTest struct {
	message string;
	input string;
	output string;
}

func (test parseTest) run(t *testing.T) {
	ast, ok := ParseString(test.input);
	if ok {
		expectMatchString(t, test.output, ast.String(), test.message);
	} else {
		t.Errorf("%v: `%v` does not parse", test.message, test.input);
	}
}