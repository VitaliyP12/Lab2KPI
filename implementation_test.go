package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t testing.T) {
    // Прості тести
    result, err := PostfixToPrefix("4 2 +")
    assert.Nil(t, err)
    assert.Equal(t, "+ 4 2", result)

    result, err = PostfixToPrefix("4 2 5 + 7 +")
    assert.Nil(t, err)
    assert.Equal(t, "+ + 4 * 2 5 7", result)

    // Складні тести
    result, err = PostfixToPrefix("5 6 2 + * 12 4 / -")
    assert.Nil(t, err)
    assert.Equal(t, "- * 5 + 6 2 / 12 4", result)

    // Тести на неправильні дані
    , err = PostfixToPrefix("")
    assert.NotNil(t, err)

    , err = PostfixToPrefix("4 2 + +")
    assert.NotNil(t, err)
}
