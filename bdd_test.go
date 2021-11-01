package bdd

import "testing"

func TestParams(t *testing.T) {
	params := Params([]string{"a: 123", "b: abc"})
	if params.Get("a") != "123" || params.Get("b") != "abc" {
		t.Fatal("error")
	}

	if params.Get("c", "xxx") != "xxx" {
		t.Fatal("bad default value")
	}
}
