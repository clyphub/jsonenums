// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// JSONenums is a tool to automate the creation of methods that satisfy the
// fmt.Stringer, json.Marshaler and json.Unmarshaler interfaces.
// Given the name of a (signed or unsigned) integer type T that has constants
// defined, jsonenums will create a new self-contained Go source file implementing
//
//  func (t T) String() string
//  func (t T) MarshalJSON() ([]byte, error)
//  func (t *T) UnmarshalJSON([]byte) error
//
// The file is created in the same package and directory as the package that defines T.
// It has helpful defaults designed for use with go generate.
//
// JSONenums is a simple implementation of a concept and the code might not be
// the most performant or beautiful to read.
//
// For example, given this snippet,
//
//	package painkiller
//
//	type Pill int
//
//	const (
//		Placebo Pill = iota
//		Aspirin
//		Ibuprofen
//		Paracetamol
//		Acetaminophen = Paracetamol
//	)
//
// running this command
//
//	jsonenums -type=Pill
//
// in the same directory will create the file pill_jsonenums.go, in package painkiller,
// containing a definition of
//
//  func (r Pill) String() string
//  func (r Pill) MarshalJSON() ([]byte, error)
//  func (r *Pill) UnmarshalJSON([]byte) error
//
// That method will translate the value of a Pill constant to the string representation
// of the respective constant name, so that the call fmt.Print(painkiller.Aspirin) will
// print the string "Aspirin".
//
// Typically this process would be run using go generate, like this:
//
//	//go:generate jsonenums -type=Pill
//
// If multiple constants have the same value, the lexically first matching name will
// be used (in the example, Acetaminophen will print as "Paracetamol").
//
// With no arguments, it processes the package in the current directory.
// Otherwise, the arguments must name a single directory holding a Go package
// or a set of Go source files that represent a single Go package.
//
// The -type flag accepts a comma-separated list of types so a single run can
// generate methods for multiple types. The default output file is
// t_jsonenums.go, where t is the lower-cased name of the first type listed.
// The suffix can be overridden with the -suffix flag and a prefix may be added
// with the -prefix flag.
//
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/clyphub/jsonenums/parser"
)

var (
	typeNames              = flag.String("type", "", "comma-separated list of type names; must be set")
	outputPrefix           = flag.String("prefix", "", "prefix to be added to the output file")
	outputSuffix           = flag.String("suffix", "_jsonenums", "suffix to be added to the output file")
	exportSnakeCaseJSON    = flag.Bool("snake_case_json", false, "Map camel case variable names to snake case json?")
	serializedPrefixToDrop = flag.String("prefix_to_drop", "", "string to drop from beginning of each iota const name when converting to string")
	allCaps                = flag.Bool("all_caps", false, "convert the serialized string to uppercase?")
	generateStringer       = flag.Bool("to_string", false, "generating ToString() function for iota?")
)

func ToSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

type CammelSnakePair struct {
	CammelRep string
	SnakeRep  string
}

func dropPrefix(str, prefix string) (string, error) {
	if strings.Index(str, prefix) != 0 {
		return "", errors.New(fmt.Sprintf("%s is not a prefix of %s", prefix, str))
	}
	return str[len(prefix):], nil
}

func main() {
	flag.Parse()
	if len(*typeNames) == 0 {
		log.Fatalf("the flag -type must be set")
	}
	types := strings.Split(*typeNames, ",")

	// Only one directory at a time can be processed, and the default is ".".
	dir := "."
	if args := flag.Args(); len(args) == 1 {
		dir = args[0]
	} else if len(args) > 1 {
		log.Fatalf("only one directory at a time")
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("unable to determine absolute filepath for requested path %s: %v",
			dir, err)
	}

	pkg, err := parser.ParsePackage(dir)
	if err != nil {
		log.Fatalf("parsing package: %v", err)
	}

	var analysis = struct {
		Command        string
		PackageName    string
		Stringer       bool
		TypesAndValues map[string][]CammelSnakePair
	}{
		Command:        strings.Join(os.Args[1:], " "),
		PackageName:    pkg.Name,
		Stringer:       *generateStringer,
		TypesAndValues: make(map[string][]CammelSnakePair),
	}

	// Run generate for each type.
	for _, typeName := range types {
		values, err := pkg.ValuesOfType(typeName)
		if err != nil {
			log.Fatalf("finding values for type %v: %v", typeName, err)
		}

		cammelSnakePairs := make([]CammelSnakePair, len(values))
		serializedNamesUsed := map[string]bool{}

		for i, rawValue := range values {
			value := rawValue
			if serializedPrefixToDrop != nil {
				var err error
				value, err = dropPrefix(rawValue, *serializedPrefixToDrop)
				if err != nil {
					log.Fatalf("Error removing prefix: %v", err)
				}
			}

			cammelSnakePairs[i].CammelRep = rawValue
			if exportSnakeCaseJSON != nil && *exportSnakeCaseJSON {
				cammelSnakePairs[i].SnakeRep = ToSnake(value)
			} else {
				cammelSnakePairs[i].SnakeRep = value
			}

			if allCaps != nil && *allCaps {
				cammelSnakePairs[i].SnakeRep = strings.ToUpper(cammelSnakePairs[i].SnakeRep)
			}
			if _, ok := serializedNamesUsed[cammelSnakePairs[i].SnakeRep]; ok {
				log.Fatalf("Multiple iota consts map to serialized value %s", cammelSnakePairs[i].SnakeRep)
			}
			serializedNamesUsed[cammelSnakePairs[i].SnakeRep] = true
		}

		analysis.TypesAndValues[typeName] = cammelSnakePairs

		var buf bytes.Buffer
		if err := generatedTmpl.Execute(&buf, analysis); err != nil {
			log.Fatalf("generating code: %v", err)
		}

		src, err := format.Source(buf.Bytes())
		if err != nil {
			// Should never happen, but can arise when developing this code.
			// The user can compile the output to see the error.
			log.Printf("warning: internal error: invalid Go generated: %s", err)
			log.Printf("warning: compile the package to analyze the error")
			src = buf.Bytes()
		}

		output := strings.ToLower(*outputPrefix + typeName +
			*outputSuffix + ".go")
		outputPath := filepath.Join(dir, output)
		if err := ioutil.WriteFile(outputPath, src, 0644); err != nil {
			log.Fatalf("writing output: %s", err)
		}
	}
}
