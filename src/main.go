package main

import (
	"bufio";
	"fmt";
	"os";
	"./golambda";
)

func main() {
	in := bufio.NewReader(os.Stdin);
	for {
		fmt.Print("> ");
		input,_ := in.ReadString('\n');
		if ast,ok := golambda.ParseString(input); ok {
			output := ast.Reduce();
			fmt.Println(output);
		}
	}
}
