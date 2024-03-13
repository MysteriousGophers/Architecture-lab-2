package lab2

import (
	"fmt"
	"testing"
)

func TestPrefixToPostfix(t *testing.T) {
	PostfixToInfix("+ 5 * - 4 2 3")
}

func ExamplePrefixToPostfix() {
	res, _ := PostfixToInfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
