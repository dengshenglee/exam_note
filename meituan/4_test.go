package main

import "testing"

func TestCandy2(t *testing.T) {
	candies := []int{3, 1, 2, 7, 10, 2, 4}
	//choose 3, 7, 4, break rule at 10
	k := 1
	if candy2(candies, k) != 24 {
		t.Fatal("Expected 24, got ", candy2(candies, k))
	}
}
