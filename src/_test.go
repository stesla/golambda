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

type fmtTest struct {
	message string;
	expected string;
	actual fmt.Stringer;
}

func (test fmtTest) run(t *testing.T) {
	if pass,_ := testing.MatchString(test.expected, test.actual.String()); !pass {
		t.Errorf("%v: expected `%v` to be `%v`", test.message, test.actual, test.expected);
	}
}
