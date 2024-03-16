package lab2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCorrectInput(t *testing.T) {
	output := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader("3 5 +"),
		Output: output,
	}
	err := handler.Compute()

	assert.Nil(t, err)
}

func TestIncorrectInput(t *testing.T) {
	output := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader("3 5 ||"),
		Output: output,
	}
	err := handler.Compute()

	assert.NotNil(t, err)
}

func TestEmptyInput(t *testing.T) {
	output := bytes.NewBuffer(nil)
	handler := ComputeHandler{
		Input:  strings.NewReader(""),
		Output: output,
	}
	err := handler.Compute()

	assert.NotNil(t, err)
}
