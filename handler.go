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
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buffer := new(bytes.Buffer)
	_, readErr := buffer.ReadFrom(ch.Reader)
	if readErr != nil {
		return readErr
	}

	formattedInput := strings.TrimRight(buffer.String(), "\n")
	result, err := PostfixToInfix(formattedInput)
	if err != nil {
		return err
	}

	fmt.Fprintln(ch.Writer, result)
	return nil
}
