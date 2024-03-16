package lab2

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	reader := bufio.NewReader(ch.Input)
	input, _ := reader.ReadString('\n')
	formattedInput := strings.TrimRight(input, "\n")
	result, err := PostfixToInfix(formattedInput)
	if err != nil {
		fmt.Fprintln(ch.Output, err)
	} else {
		fmt.Fprintln(ch.Output, result)
	}
	return err
}
