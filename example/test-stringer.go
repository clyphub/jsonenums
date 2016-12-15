package main

//go:generate jsonenums -to_string -type=toString
type toString int

const (
	elemA toString = iota
	elemB
	elemC
)
