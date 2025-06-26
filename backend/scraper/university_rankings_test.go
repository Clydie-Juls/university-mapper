package main

import "testing"

func TestCleanRank(t *testing.T) {
	rank := "50"
	result := CleanRank(rank)

	if result != 50 {
		t.Errorf("expected %d, got %d", 50, result)
	}

	rank = "100-200"
	result = CleanRank(rank)

	if result != 100 {
		t.Errorf("expected %d, got %d", 100, result)
	}

	rank = "reporter"
	result = CleanRank(rank)

	if result != 0 {
		t.Errorf("expected %d, got %d", 0, result)
	}
}
