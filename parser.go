package parser

import (
	"fmt"
	"reflect"
	"strconv"
)

type ParserBuilder struct {
	parsers []func(interface{}) error
}

func New() *ParserBuilder {
	return &ParserBuilder{}
}

func (p *ParserBuilder) Add(v interface{}) {
	switch t := v.(type) {
	case *int32:
		// fmt.Printf("Integer32: %v\n", *t)
		p.parsers = append(p.parsers, Int32Parser(v.(*int32)))
	case *int64:
		// fmt.Printf("Integer32: %v\n", *t)
		p.parsers = append(p.parsers, Int64Parser(v.(*int64)))
	case *float32:
		// fmt.Printf("Float32: %v\n", *t)
		p.parsers = append(p.parsers, Float32Parser(v.(*float32)))
	case *float64:
		// fmt.Printf("Float64: %v\n", *t)
		p.parsers = append(p.parsers, Float64Parser(v.(*float64)))
	case *string:
		// fmt.Printf("String: %v\n", *t)
		p.parsers = append(p.parsers, StringParser(v.(*string)))
	case *bool:
		// fmt.Printf("Bool: %v\n", *t)
		p.parsers = append(p.parsers, BoolParser(v.(*bool)))
	case []interface{}:
		for _, n := range t {
			p.Add(n)
		}
	default:
		var r = reflect.TypeOf(t)
		fmt.Printf("Other:%v\n", r)
	}
}

func Int32Parser(dest *int32) func(interface{}) error {
	return func(v interface{}) error {
		d, err := strconv.ParseInt(v.(string), 10, 32)
		*dest = int32(d)
		return err
	}
}

func Int64Parser(dest *int64) func(interface{}) error {
	return func(v interface{}) (err error) {
		*dest, err = strconv.ParseInt(v.(string), 10, 64)
		return err
	}
}

func Float32Parser(dest *float32) func(interface{}) error {
	return func(v interface{}) error {
		d, err := strconv.ParseFloat(v.(string), 32)
		*dest = float32(d)
		return err
	}
}

func Float64Parser(dest *float64) func(interface{}) error {
	return func(v interface{}) (err error) {
		*dest, err = strconv.ParseFloat(v.(string), 64)
		return err
	}
}

func BoolParser(dest *bool) func(interface{}) error {
	return func(v interface{}) (err error) {
		*dest, err = strconv.ParseBool(v.(string))
		return err
	}
}

func StringParser(dest *string) func(interface{}) error {
	return func(v interface{}) error {
		*dest = v.(string)
		return nil
	}
}

func (p *ParserBuilder) Parse(row []interface{}) error {
	if len(row) != len(p.parsers) {
		return fmt.Errorf("row len, want: %d, got: %d", len(p.parsers), len(row))
	}
	var err error
	for i, parser := range p.parsers {
		err = parser(row[i])
		if err != nil {
			return err
		}
	}
	return nil
}
