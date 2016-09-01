package main

//go:generate jsonenums -type=CustomStatus -prefix_to_drop=CustomStatus -all_caps -snake_case_json

type CustomStatus int

const (
	CustomStatusError CustomStatus = iota
	CustomStatusOK
	CustomStatusFun
	CustomStatusNoFun
)
