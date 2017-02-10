package main

import "testing"

func TestTwice(t *testing.T) {
	if x := twice(2); x != 4 {
		t.Error("Expected 4 but got", x)
	}
}
