package main

import "testing"

func TestFibOne(t *testing.T) {
	n := fibOne(6)

	if n != 8 {
		t.Errorf("expected 8 but got %d", n)
	}
}

func TestFibTwo(t *testing.T) {
	n := fibTwo(6)

	if n != 8 {
		t.Errorf("expected 8 but got %d", n)
	}
}
