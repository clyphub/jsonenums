package main

import (
	"testing"
)

func Test_ToSnake(t *testing.T) {
	sample1 := "aBcDEfg"
	if "a_bc_d_efg" != ToSnake(sample1) {
		t.Fail()
	}
}
