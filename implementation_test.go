package lab2

import (
	"fmt"
	"testing"
)

func TestPrefixToPostfix(t *testing.T) {
	PrefixToPostfix("+ 5 * - 4 2 3")
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
