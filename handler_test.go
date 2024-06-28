package lab2

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	type fields struct {
		Input  io.Reader
		Output io.Writer
	}

	var writer bytes.Buffer
	mockError := "mockError"

	tests := []struct {
		name       string
		fields     fields
		calculator PrefixCalculatorSpy
		isError    bool
	}{
		{"Should pass", fields{strings.NewReader("4 2 - 3 * 5 +"), &writer}, PrefixCalculatorSpy{"(+(*(-42)3)5)", nil}, false},
		{"Should pass", fields{strings.NewReader("+5 - 10 * 2  6"), &writer}, PrefixCalculatorSpy{"", errors.New(mockError)}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			computeHandler := &ComputeHandler{
				Input:      test.fields.Input,
				Output:     test.fields.Output,
				Calculator: &test.calculator,
			}
			err := computeHandler.Compute()

			if test.isError {
				assert.Error(t, err, mockError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
