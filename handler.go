package lab2

import (
	"bytes"
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
	buffer := new(bytes.Buffer)
	_, readErr := buffer.ReadFrom(ch.Input)

	if readErr != nil {
		fmt.Printf("Error reading input: %v\n", readErr)
		return readErr
	}

	formattedInput := strings.TrimRight(buffer.String(), "\n")
	result, err := PostfixToInfix(formattedInput)
	if err != nil {
		fmt.Fprintln(ch.Output, err)
	} else {
		fmt.Fprintln(ch.Output, result)
	}
	return err
}
