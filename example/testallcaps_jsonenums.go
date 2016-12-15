// generated by jsonenums -type=TestAllCaps -all_caps=true -snake_case_json=true; DO NOT EDIT

package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var (
	_TestAllCapsNameToValue = map[string]TestAllCaps{
		"SOME_CAMEL":     someCamel,
		"SOME_SNAKE":     some_snake,
		"SO_M_MA_DN_ESS": SoMMaDnEss,
	}

	_TestAllCapsValueToName = map[TestAllCaps]string{
		someCamel:  "SOME_CAMEL",
		some_snake: "SOME_SNAKE",
		SoMMaDnEss: "SO_M_MA_DN_ESS",
	}
)

func init() {
	var v TestAllCaps
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TestAllCapsNameToValue = map[string]TestAllCaps{
			interface{}(someCamel).(fmt.Stringer).String():  someCamel,
			interface{}(some_snake).(fmt.Stringer).String(): some_snake,
			interface{}(SoMMaDnEss).(fmt.Stringer).String(): SoMMaDnEss,
		}
	}
}

func (r TestAllCaps) toString() (string, error) {
	s, ok := _TestAllCapsValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid TestAllCaps: %d", r)
	}
	return s, nil
}

func (r TestAllCaps) getString() (string, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return s.String(), nil
	}
	return r.toString()
}

func (r *TestAllCaps) setValue(str string) error {
	v, ok := _TestAllCapsNameToValue[str]
	if !ok {
		return fmt.Errorf("invalid TestAllCaps %q", str)
	}
	*r = v
	return nil
}

// MarshalJSON is generated so TestAllCaps satisfies json.Marshaler.
func (r TestAllCaps) MarshalJSON() ([]byte, error) {
	s, err := r.getString()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so TestAllCaps satisfies json.Unmarshaler.
func (r *TestAllCaps) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TestAllCaps should be a string, got %s", data)
	}
	return r.setValue(s)
}

//Scan an input string into this structure for use with GORP
func (r *TestAllCaps) Scan(i interface{}) error {
	switch t := i.(type) {
	case []byte:
		return r.setValue(string(t))
	case string:
		return r.setValue(t)
	default:
		return fmt.Errorf("Can't scan %T into type %T", i, r)
	}
	return nil
}

func (r TestAllCaps) Value() (driver.Value, error) {
	return r.getString()
}
