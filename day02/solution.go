package day02

import (
	"AdventOfCode2019/util"
)

var puzzleInput = "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,10,19,1,19,6,23,2,23,13,27,1,27,5,31,2,31,10,35,1,9,35,39,1,39,9,43,2,9,43,47,1,5,47,51,2,13,51,55,1,55,9,59,2,6,59,63,1,63,5,67,1,10,67,71,1,71,10,75,2,75,13,79,2,79,13,83,1,5,83,87,1,87,6,91,2,91,13,95,1,5,95,99,1,99,2,103,1,103,6,0,99,2,14,0,0"

func solvePart1() int {
	program := util.ToInts(puzzleInput)
	program[1] = 12
	program[2] = 2
	computer := util.NewComputer(program, nil)
	computer.Run()
	return computer.Get(0)
}

func solvePart2() int {
	want := 19690720
	noun, verb := 0, 0
	for i1 := 0; i1 <= 99; i1++ {
		for i2 := 0; i2 <= 99; i2++ {
			program := util.ToInts(puzzleInput)
			program[1] = i1
			program[2] = i2
			computer := util.NewComputer(program, nil)
			computer.Run()
			if computer.Get(0) == want {
				noun = i1
				verb = i2
				break
			}
		}
	}
	return 100*noun + verb
}