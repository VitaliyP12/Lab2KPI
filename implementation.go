package lab2

import (
	"errors"
	"strings"
)

type PrefixCalculator struct{}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func (ptic *PrefixCalculator) ConvertPrefixToPostfix(str string) (string, error) {
	tokens := strings.Fields(str)
	stack := []string{}

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return "", errors.New("invalid postfix expression")
			}

			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			prefixExpr := "(" + token + op1 + op2 + ")"
			stack = append(stack, prefixExpr)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", errors.New("invalid postfix expression")
	}

	return stack[0], nil
}
