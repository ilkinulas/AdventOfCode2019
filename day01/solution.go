package day01

import (
	"bufio"
	"os"
	"strconv"
)

func solve1() int {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		mass, _ := strconv.Atoi(line)
		sum += requiredFuel(mass)
	}
	return sum
}

func solve2() int {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		mass, _ := strconv.Atoi(line)
		sum += requiredFuelRecursive(mass)
	}
	return sum
}

func requiredFuel(mass int) int {
	return (mass / 3) - 2
}

func requiredFuelRecursive(mass int) int {
	required := requiredFuel(mass)

	if required <= 0 {
		return 0
	}

	return required + requiredFuelRecursive(required)
}
