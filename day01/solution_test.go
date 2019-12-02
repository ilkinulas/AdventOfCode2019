package day01

import (
	"testing"
)

func TestRequiredFuel(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, test := range tests {
		got := requiredFuel(test.mass)
		if got != test.fuel {
			t.Fatalf("expected %v, got %v", test.fuel, got)
		}
	}
}

func TestRequiredFuelRecursive(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, test := range tests {
		got := requiredFuelRecursive(test.mass)
		if got != test.fuel {
			t.Errorf("expected %v, got %v", test.fuel, got)
		}
	}
}

func TestSolution(t *testing.T) {
	answer := solve1()
	if answer != 3325156 {
		t.Fatalf("wrong answer : %v", answer)
	}

	answer = solve2()
	if answer != 4984866 {
		t.Fatalf("wrong answer : %v", answer)
	}
}
