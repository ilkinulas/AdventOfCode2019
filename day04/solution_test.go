package day04

import (
	"fmt"
	"testing"
)

func Test_IsValid(t *testing.T) {
	tests := []struct {
		pass int
		want bool
	}{
		//{111111, true},
		//{223450, false},
		//{123789, false},
		{112233, true},
		{123444, false},
		{111122, true},
	}

	for _, test := range tests {
		if isValidPart2(test.pass) != test.want {
			t.Errorf("expecting %v for %v", test.want, test.pass)
		}
	}
}

func Test_Solution(t *testing.T) {
	count1, count2 := countPossiblePasswords(256310, 732736)
	fmt.Printf("%v\n", count2)
	if count1 != 979 {
		t.Errorf("wrong answer for part1 %v", count1)
	}

	if count2 != 635 {
		t.Errorf("wrong answer for part2 %v", count2)
	}
}
