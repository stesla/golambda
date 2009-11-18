package golambda

import (
	"testing";
)

func TestParse(t *testing.T) {
	test(t, []testCase{
		&parseTest{"ident - single chart", "x", "x"},
		&parseTest{"ident - lowers", "foo", "foo"},
		&parseTest{"ident - all", "aZ_1", "aZ_1"},
		&parseTest{"abstraction", "fn x. y", "fn x. y"}
	});
}