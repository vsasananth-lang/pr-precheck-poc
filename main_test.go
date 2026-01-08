package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 6 {
		t.Fatalf("Expected 5, got %d", result)
	}
}
