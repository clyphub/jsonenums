// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:generate jsonenums -type=TestCasing -snake_case_json=true
type TestCasing int

const (
	caseMadnessA TestCasing = iota
	caseMaDnEEsB
	normalCaseExample
)

//go:generate jsonenums -type=TestPrefixDrop -prefix_to_drop=prefixDrop
type TestPrefixDrop int

const (
	prefixDropThisS TestPrefixDrop = iota
	prefixDrophould
	prefixDropWork
)

//go:generate jsonenums -type=TestUnderscoreToSpace -prefix_to_drop=TestUnderscoreToSpace_ -underscore_to_space=true
type TestUnderscoreToSpace int

const (
	TestUnderscoreToSpace_Foo TestUnderscoreToSpace = iota
	TestUnderscoreToSpace_FooBar
	TestUnderscoreToSpace_Foo_Baz
	TestUnderscoreToSpace_Foo_Baz_Qwerty
)

//go:generate jsonenums -type=TestAllCaps -all_caps=true -snake_case_json=true
type TestAllCaps int

const (
	someCamel TestAllCaps = iota
	some_snake
	SoMMaDnEss
)

type TestFailWithCaps int

const (
	FaIL TestFailWithCaps = iota
	FAiL
)

//go:generate jsonenums -type=ShirtSize

type ShirtSize byte

const (
	NA ShirtSize = iota
	XS
	S
	M
	L
	XL
)

//go:generate jsonenums -type=WeekDay

type WeekDay int

const (
	Monday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (d WeekDay) String() string {
	switch d {
	case Monday:
		return "Dilluns"
	case Tuesday:
		return "Dimarts"
	case Wednesday:
		return "Dimecres"
	case Thursday:
		return "Dijous"
	case Friday:
		return "Divendres"
	case Saturday:
		return "Dissabte"
	case Sunday:
		return "Diumenge"
	default:
		return "invalid WeekDay"
	}
}

func main() {
	v := struct {
		Size ShirtSize
		Day  WeekDay
	}{M, Friday}
	if err := json.NewEncoder(os.Stdout).Encode(v); err != nil {
		log.Fatal(err)
	}

	input := `{"Size":"XL", "Day":"Dimarts"}`
	if err := json.NewDecoder(strings.NewReader(input)).Decode(&v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoded %s as %+v\n", input, v)
}
