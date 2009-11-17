package golambda

import (
	"fmt";
	"testing";
)

func expectFmt(t *testing.T, expected string, actual fmt.Stringer) {
	if pass,_ := testing.MatchString(expected, actual.String()); !pass {
		t.Errorf("expected `%v` to be `%v`", actual, expected)
	}
}
