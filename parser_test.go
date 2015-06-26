package parser

import (
	"testing"
)

func TestInt32(t *testing.T) {
	p := New()
	var v int32
	p.Add(&v)
	if err := p.Parse([]interface{}{"32"}); err != nil {
		t.Errorf("parse error: %v", err)
	}
	if v != 32 {
		t.Errorf("want: 32, got %d", v)
	}
}

func TestInt64(t *testing.T) {
	p := New()
	var v int64
	p.Add(&v)
	if err := p.Parse([]interface{}{"64"}); err != nil {
		t.Errorf("parse error: %v", err)
	}
	if v != 64 {
		t.Errorf("want: 64, got %d", v)
	}
}

func TestFloat32(t *testing.T) {
	p := New()
	var v float32
	p.Add(&v)
	if err := p.Parse([]interface{}{"3.14"}); err != nil {
		t.Errorf("parse error: %v", err)
	}
	if v != 3.14 {
		t.Errorf("want: 3.14, got %f", v)
	}
}

func TestSliceOfIfc(t *testing.T) {
	p := New()
	var f float32
	var s string
	ifc := []interface{}{&f, &s}
	p.Add(ifc)
	if err := p.Parse([]interface{}{"3.14", "napis"}); err != nil {
		t.Errorf("parse error: %v", err)
	}
	if f != 3.14 {
		t.Errorf("want: 3.14, got %f", f)
	}

	if s != "napis" {
		t.Errorf("want: `napis`, got %q", s)
	}
}
