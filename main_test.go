package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Display(t *testing.T) {
	n := randomNumber()
	assert.GreaterOrEqual(t, n, 0)
	assert.LessOrEqual(t, n, 1000)
}
