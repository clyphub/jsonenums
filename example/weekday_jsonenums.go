// generated by jsonenums -type=WeekDay; DO NOT EDIT

package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var (
	_WeekDayNameToValue = map[string]WeekDay{
		"Monday":    Monday,
		"Tuesday":   Tuesday,
		"Wednesday": Wednesday,
		"Thursday":  Thursday,
		"Friday":    Friday,
		"Saturday":  Saturday,
		"Sunday":    Sunday,
	}

	_WeekDayValueToName = map[WeekDay]string{
		Monday:    "Monday",
		Tuesday:   "Tuesday",
		Wednesday: "Wednesday",
		Thursday:  "Thursday",
		Friday:    "Friday",
		Saturday:  "Saturday",
		Sunday:    "Sunday",
	}
)

func init() {
	var v WeekDay
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_WeekDayNameToValue = map[string]WeekDay{
			interface{}(Monday).(fmt.Stringer).String():    Monday,
			interface{}(Tuesday).(fmt.Stringer).String():   Tuesday,
			interface{}(Wednesday).(fmt.Stringer).String(): Wednesday,
			interface{}(Thursday).(fmt.Stringer).String():  Thursday,
			interface{}(Friday).(fmt.Stringer).String():    Friday,
			interface{}(Saturday).(fmt.Stringer).String():  Saturday,
			interface{}(Sunday).(fmt.Stringer).String():    Sunday,
		}
	}
}

func (r WeekDay) getString() (string, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return s.String(), nil
	}

	s, ok := _WeekDayValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid WeekDay: %d", r)
	}
	return s, nil

}

func (r *WeekDay) setValue(str string) error {
	v, ok := _WeekDayNameToValue[str]
	if !ok {
		return fmt.Errorf("invalid WeekDay %q", str)
	}
	*r = v
	return nil
}

// MarshalJSON is generated so WeekDay satisfies json.Marshaler.
func (r WeekDay) MarshalJSON() ([]byte, error) {
	s, err := r.getString()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so WeekDay satisfies json.Unmarshaler.
func (r *WeekDay) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("WeekDay should be a string, got %s", data)
	}
	return r.setValue(s)
}

//Scan an input string into this structure for use with GORP
func (r *WeekDay) Scan(i interface{}) error {
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

func (r WeekDay) Value() (driver.Value, error) {
	return r.getString()
}
