package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(1, 2)
	if result != -1 {
		t.Errorf("Expected -1, but got %d", result)
	}
}
