package lab2

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

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

	_, err = fmt.Fprintln(ch.Writer, result)
	if err != nil {
		return err
	}
	return nil
}
