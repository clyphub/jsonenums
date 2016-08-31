package main

import (
	"testing"
)

type SnakeTest struct {
	input  string
	output string
}

var tests = []SnakeTest{
	{"a", "a"},
	{"snake", "snake"},
	{"A", "a"},
	{"ID", "id"},
	{"MOTD", "motd"},
	{"Snake", "snake"},
	{"SnakeTest", "snake_test"},
	{"SnakeID", "snake_id"},
	{"SnakeIDGoogle", "snake_id_google"},
	{"LinuxMOTD", "linux_motd"},
	{"OMGWTFBBQ", "omgwtfbbq"},
	{"omg_wtf_bbq", "omg_wtf_bbq"},
}

func TestToSnake(t *testing.T) {
	for _, test := range tests {
		if ToSnake(test.input) != test.output {
			t.Errorf(`ToSnake("%s"), wanted "%s", got \%s"`, test.input, test.output, ToSnake(test.input))
		}
	}
}

func TestDropPrefix(t *testing.T) {
	_, err := dropPrefix("abcd", "bc")
	if err == nil {
		t.Errorf("Expected error when dropping string that is not a prefix")
	}

	result, err := dropPrefix("bababa", "ba")
	if err != nil {
		t.Errorf("Expected no error, received %v", err)
	}

	if result != "baba" {
		t.Errorf("Expected `baba`, received %s", result)
	}
}
