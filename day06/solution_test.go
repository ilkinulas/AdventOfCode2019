package day06

import (
	"testing"
)

func Test_Solve(t *testing.T) {
	part1, part2 := solve()

	if part1 != 227612 {
		t.Errorf("wrong answer %v", part1)
	}
	if part2 != 454 {
		t.Errorf("wrong answer %v", part2)
	}
}
