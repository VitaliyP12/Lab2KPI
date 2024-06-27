package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"4 2 +", "+ 4 2"},
		{"4 2 5 * + 7 +", "+ + 4 * 2 5 7"},
	}

	for _, tt := range tests {
		input := strings.NewReader(tt.input)
		output := new(strings.Builder)
		handler := &ComputeHandler{Input: input, Output: output}

		err := handler.Compute()
		assert.Nil(t, err)
		assert.Equal(t, tt.expected, strings.TrimSpace(output.String()))
	}

	// Тест на неправильні дані
	input := strings.NewReader("")
	output := new(strings.Builder)
	handler := &ComputeHandler{Input: input, Output: output}

	err := handler.Compute()
	assert.NotNil(t, err)
}
