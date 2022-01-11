package main

import (
	"testing"
)

type output struct {
	check  bool
	person emp
}

func TestEmp(t *testing.T) {
	tables := []struct {
		input emp
		out   output
	}{
		{input: emp{name: "Rakesh", age: 21}, out: output{}},
		{input: emp{name: "satya", age: 24}, out: output{true, emp{"satya", 24}}},
	}
	for _, v := range tables {
		b, result := checker(v.input)

		if b != v.out.check {
			t.Errorf("accepted %v but got %v", v.out, b)
		}
		if result != v.out.person {
			t.Errorf("accepted %v but got %v", v.out, b)
		}
	}
}
