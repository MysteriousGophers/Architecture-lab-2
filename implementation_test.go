package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

type testCase struct {
	input    string
	expected string
}

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestBob(c *C) {
	c.Assert("bob", Equals, "bob")
}

func (s *MySuite) TestPostfixToInfixValidationEmptyInput(c *C) {
	_, err := PostfixToInfix("")
	c.Assert(err, NotNil)
}

func (s *MySuite) TestPostfixToInfixValidationBlankInput(c *C) {
	_, err := PostfixToInfix(" \n")
	c.Assert(err, NotNil)
}

func (s *MySuite) TestPostfixToInfixValidationUnsupportedInput(c *C) {
	_, err := PostfixToInfix("5 4 &")
	c.Assert(err, NotNil)
}

func (s *MySuite) TestPostfixToInfixSimpleExpressions(c *C) {
	testCases := []testCase{
		{"7 2 + 4 *", "(7 + 2) * 4"},
		{"256 128 - 4 /", "(256 - 128) / 4"},
		{"256 128 / 4 *", "256 / 128 * 4"},
		{"7 2 / 4 +", "7 / 2 + 4"},
		{"7 2 - 4 3 + ^", "(7 - 2) ^ (4 + 3)"},
		{"7 2 * 4 2 / ^", "(7 * 2) ^ (4 / 2)"},
	}

	for _, e := range testCases {
		result, err := PostfixToInfix(e.input)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, e.expected)
	}
}

func (s *MySuite) TestPostfixToInfixComplexExpressions(c *C) {
	testCases := []testCase{
		{"3 4 + 5 * 6 - 7 / 8 + 9 *", "(((3 + 4) * 5 - 6) / 7 + 8) * 9"},
		{"4 5 * 3 2 - 1 2 / + - 6 7 + *", "(4 * 5 - (3 - 2 + 1 / 2)) * (6 + 7)"},
		{"7 3 / 6 2 * + 4 2 * - 8 9 + /", "(7 / 3 + 6 * 2 - 4 * 2) / (8 + 9)"},
		{"9 3 / 2 1 * 7 + 8 4 / - + 5 6 * -", "9 / 3 + (2 * 1 + 7 - 8 / 4) - 5 * 6"},
	}

	for _, e := range testCases {
		result, err := PostfixToInfix(e.input)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, e.expected)
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 2 + 2
}
