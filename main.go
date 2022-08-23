package main

import (
	"fmt"
	functionsandpointers "funcoes-e-ponteiros/functions-and-pointers"
)

func main() {
	var n1 int = 20

	fmt.Printf("n1 before: %d\n", n1)

	functionsandpointers.ReverseSignal(&n1)

	fmt.Printf("n1 after: %d\n", n1)
}
