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

func testFmt(t *testing.T, tests []fmtTest) {
	for _,test := range tests {
		if pass,_ := testing.MatchString(test.expected, test.actual.String()); !pass {
			t.Errorf("%v: expected `%v` to be `%v`", test.message, test.actual, test.expected);
		}
	}
}