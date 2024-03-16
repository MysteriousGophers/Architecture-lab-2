package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
)

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestCorrectInput(c *C) {
	expected := "3 + 5\n"
	output := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader("3 5 +"),
		Output: output,
	}
	err := handler.Compute()

	c.Assert(err, IsNil)
	c.Assert(output.String(), Equals, expected)
}

func (s *HandlerSuite) TestIncorrectInput(c *C) {
	output := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader("3 5 ||"),
		Output: output,
	}
	err := handler.Compute()

	c.Assert(err, NotNil)
}
