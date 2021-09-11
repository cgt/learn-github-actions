package main

import "testing"

func Test2Plus2Equals4(t *testing.T) {
	actual := add(2, 2)
	if expected := 4; actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
