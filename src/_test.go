package golambda

import (
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

func expectBoolean(t *testing.T, expected, actual bool, message string) {
	if expected != actual {
		t.Errorf("%v: expected %v to be %v", message, actual, expected);
	}
}

func expectString(t *testing.T, expected, actual, message string) {
	if expected != actual {
		t.Errorf("%v: expected `%v` to be `%v`", message, actual, expected);
	}
}
