package main

import (
	"errors"
	"strings"
)

// isOperator перевіряє, чи є символ оператором
func isOperator(c string) bool {
	operators := "+-*/^"
	return strings.Contains(operators, c)
}

// PostfixToPrefix перетворює постфіксний вираз у префіксний
func PostfixToPrefix(postfix string) (string, error) {
	if postfix == "" {
		return "", errors.New("порожній вхідний вираз")
	}
	stack := []string{}
	tokens := strings.Fields(postfix)

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return "", errors.New("недостатньо операндів у виразі")
			}
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			newExpr := token + " " + op2 + " " + op1
			stack = append(stack, newExpr)
		} else {
			stack = append(stack, token)
		}
	}
	if len(stack) != 1 {
		return "", errors.New("некоректний вираз")
	}
	return stack[0], nil
}
