package main

import "testing"

func TestTour(t *testing.T) {
	chocolate := []int{1, 2, 2, 4, 5}
	questions := []int{1, 3, 7, 9, 15}
	correct := []int{1, 1, 2, 3, 3}
	for i, x := range questions {
		answer := tour(x, chocolate)
		if answer != correct[i] {
			t.Error("Expected ", correct[i], ", got ", answer)
		}
	}
}
