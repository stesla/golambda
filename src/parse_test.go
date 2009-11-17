package golambda

import (
	"testing";
)

func TestParse(t *testing.T) {
	test(t, []testCase{
		&parseTest{"variable", "x", "x"}
	});
}