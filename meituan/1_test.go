package main

import "testing"

func TestTrain(t *testing.T) {
	in_order := []int{1, 2, 3, 4, 5}
	out_order := []int{4, 5, 3, 2, 1}
	if !train(in_order, out_order) {
		t.Error("Expected true, got false")
	}
}
