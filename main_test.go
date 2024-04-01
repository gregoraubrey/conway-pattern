package main

import "testing"

func TestCalculateNextNumber(t *testing.T) {
	t.Run("Test CalculateNextNumber with 1 and 1", func(t *testing.T) {
		result := calculateNextNumber(1, 1)
		expected := 2
		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
	t.Run("Test CalculateNextNumber with 2 and 3", func(t *testing.T) {
		result := calculateNextNumber(2, 3)
		expected := 5
		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
	t.Run("Test CalculateNextNumber with 3 and 5", func(t *testing.T) {
		result := calculateNextNumber(3, 5)
		expected := 4
		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
	t.Run("Test CalculateNextNumber with 5 and 8", func(t *testing.T) {
		result := calculateNextNumber(5, 8)
		expected := 13
		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
	t.Run("Test CalculateNextNumber with 30 and 20", func(t *testing.T) {
		result := calculateNextNumber(30, 20)
		expected := 25
		if result != expected {
			t.Errorf("result %d expected %d", result, expected)
		}
	})
}
