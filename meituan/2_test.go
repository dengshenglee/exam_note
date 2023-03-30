package main

import "testing"

func TestCandy(t *testing.T) {
	candies := []int{3, 1, 2, 7, 10, 2, 4}
	if candy(candies) != 14 {
		t.Fatal("Expected 14, got ", candy(candies))
	}
}
