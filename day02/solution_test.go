package day02

import (
	"AdventOfCode2019/util"
	"testing"
)

func TestSimpleComputer(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 0, 0, 0, 99}, 2},
		{[]int{2, 3, 0, 3, 99}, 2},
		{[]int{2, 4, 4, 5, 99, 0}, 2},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 30},
	}

	for _, test := range tests {
		computer := util.NewComputer(test.input, nil)
		computer.Run()
		got := computer.Get(0)
		if got != test.want {
			t.Errorf("Expecting %v, got %v", test.want, got)
		}
	}
}

func TestSolve1(t *testing.T) {
	answer := solvePart1()
	if answer != 3790645 {
		t.Errorf("wrong answer : %v", answer)
	}
}

func TestSolve2(t *testing.T) {
	if solvePart2() != 6577 {
		t.Errorf("wrong answer!")
	}
}
