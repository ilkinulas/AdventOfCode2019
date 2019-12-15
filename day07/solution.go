package day07

import (
	"AdventOfCode2019/util"
)

func MaxThrusterSignal(input string) int {
	program := util.ToInts(input)
	permutations := util.PermutationsOf([]int{0, 1, 2, 3, 4})

	var outputs []int
	for _, phases := range permutations {
		output := 0
		for _, phaseSetting := range phases {
			amp := util.NewComputer(program, []int{output, phaseSetting})
			output = amp.Run()
			outputs = append(outputs, output)
		}
	}
	return util.Max(outputs)
}

func FeedbackLoop(input string) int {
	program := util.ToInts(input)
	permutations := util.PermutationsOf([]int{5, 6, 7, 8, 9})

	var outputs []int

	for _, phases := range permutations {
		amplifiers := []*util.Computer{
			util.NewComputer(program, [] int{phases[0]}),
			util.NewComputer(program, [] int{phases[1]}),
			util.NewComputer(program, [] int{phases[2]}),
			util.NewComputer(program, [] int{phases[3]}),
			util.NewComputer(program, [] int{phases[4]}),
		}

		previousOutput := 0
		output := 0
		for allRunning(amplifiers) {
			for i := range amplifiers {
				amplifiers[i].Input(previousOutput)
				output = amplifiers[i].Run()
				if !amplifiers[i].IsHalted() {
					previousOutput = output
				}
			}
		}
		outputs = append(outputs, previousOutput)
	}

	return util.Max(outputs)
}

func allRunning(amplifiers [] *util.Computer) bool {
	for _, a := range amplifiers {
		if a.IsHalted() {
			return false
		}
	}
	return true
}
