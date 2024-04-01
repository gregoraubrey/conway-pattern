package main

import (
	"reflect"
	"testing"
)

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

func TestBuildSequence(t *testing.T) {
	t.Run("Test buildSequence with 1, 1, 5", func(t *testing.T) {
		result := buildSequence(1, 1, 5)
		expected := []int{1, 1, 2, 3, 5}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})
	t.Run("Test buildSequence with 2, 3, 5", func(t *testing.T) {
		result := buildSequence(2, 3, 5)
		expected := []int{2, 3, 5, 4, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})
	t.Run("Test buildSequence with 3, 5, 5", func(t *testing.T) {
		result := buildSequence(3, 5, 5)
		expected := []int{3, 5, 4, 3, 7}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})
	t.Run("Test buildSequence with 5, 8, 5", func(t *testing.T) {
		result := buildSequence(5, 8, 5)
		expected := []int{5, 8, 13, 7, 10}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})
	t.Run("Test buildSequence with 30, 20, 5", func(t *testing.T) {
		result := buildSequence(30, 20, 5)
		expected := []int{30, 20, 25, 15, 20}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})
	t.Run("Test buildSequence with 1, 1, 2", func(t *testing.T) {
		result := buildSequence(1, 1, 2)
		expected := []int{1, 1}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("result %v expected %v", result, expected)
		}
	})
}
