package main

import(
	"testing";
	golambda "./test_golambda";
)

func main() {
	testing.Main([]testing.Test{
		testing.Test{"String", golambda.TestString},
		testing.Test{"Parse", golambda.TestParse},
		testing.Test{"OccursFree", golambda.TestOccursFree},
		testing.Test{"Substitute", golambda.TestSubstitute},
		testing.Test{"Reduce", golambda.TestReduce}
	});
}