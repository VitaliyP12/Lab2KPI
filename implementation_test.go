package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPrefixToPostfix(t *testing.T) {
	tests := []struct {
		name        string
		expression  string
		expected    string
		expectedErr bool
	}{
		{
			name:        "Simple expression",
			expression:  "2 3 +",
			expected:    "(+23)",
			expectedErr: false,
		},
		{
			name:        "Complex expression",
			expression:  "2 3 * 4 +",
			expected:    "(+(*23)4)",
			expectedErr: false,
		},
	}

	prefixCalc := PrefixCalculator{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := prefixCalc.ConvertPrefixToPostfix(test.expression)
			if test.expectedErr {
				assert.Error(t, err, "expected error")
			} else {
				assert.NoError(t, err, "unexpected error")
				assert.Equal(t, test.expected, result, "result not as expected")
			}
		})
	}
}

func ExamplePrefixCalculator_ConvertPrefixToPostfix() {
	prefixCalc := PrefixCalculator{}

	result, err := prefixCalc.ConvertPrefixToPostfix("+ 2 3")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)

	result, err = prefixCalc.ConvertPrefixToPostfix("+ * 2 3 4")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)

	result, err = prefixCalc.ConvertPrefixToPostfix("")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)

	result, err = prefixCalc.ConvertPrefixToPostfix("+ 0")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
