package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestBob(c *C) {
	c.Assert("bob", Equals, "bob")
}

func (s *MySuite) TestPostfixToInfixSimpleAddition(c *C) {
	result, err := PostfixToInfix("254 256 +")
	c.Assert(err, IsNil)
	c.Assert(result, Equals, "254 + 256")
}

func (s *MySuite) TestPostfixToInfixSimpleSubtraction(c *C) {
	result, err := PostfixToInfix("4 1 -")
	c.Assert(err, IsNil)
	c.Assert(result, Equals, "4 - 1")
}

func (s *MySuite) TestPostfixToInfixSimpleMultiplication(c *C) {
	result, err := PostfixToInfix("4 5 *")
	c.Assert(err, IsNil)
	c.Assert(result, Equals, "4 * 5")
}

func (s *MySuite) TestPostfixToInfixSimpleDivision(c *C) {
	result, err := PostfixToInfix("4 2 /")
	c.Assert(err, IsNil)
	c.Assert(result, Equals, "4 / 2")
}

func (s *MySuite) TestPostfixToInfixSimpleExponentiation(c *C) {
	result, err := PostfixToInfix("4 2 ^")
	c.Assert(err, IsNil)
	c.Assert(result, Equals, "4 ^ 2")
}

func ExamplePrefixToPostfix() {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 2 + 2
}
