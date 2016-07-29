package main

import (
	"testing"

	"_vendor/github.com/stretchr/testify/assert"
)

func Test_ToSnake(t *testing.T) {
	sample1 := "aBcDEfg"
	assert.Equal(t, "a_bc_d_efg", ToSnake(sample1))
}
