package day06

import (
	"AdventOfCode2019/util"
	"strings"
)

func solve() (int, int) {
	input := util.ReadLines("input.txt")
	planets := map[string]bool{}
	orbits := map[string]string{}

	for _, line := range input {
		parts := strings.Split(line, ")")
		from := parts[0]
		to := parts[1]

		planets[from] = true
		planets[to] = true
		orbits[to] = from
	}

	sum := 0
	for planet := range planets {
		sum += countOrbits(planet, orbits)
	}

	return sum, minOrbitsBetween("YOU", "SAN", orbits)
}

func minOrbitsBetween(from, to string, orbits map[string]string) int {
	youPath := path("COM", from, orbits)
	sanPath := path("COM", to, orbits)
	for i, planetYouPath := range youPath {
		for j, planetSanPath := range sanPath {
			if planetYouPath == planetSanPath {
				return i + j - 2
			}
		}
	}
	return 0
}

func countOrbits(planet string, orbits map[string]string) int {
	count := 0
	for {
		if _, ok := orbits[planet]; !ok {
			break
		}
		planet = orbits[planet]
		count += 1
	}
	return count
}

func path(to, from string, orbits map[string]string) []string {
	var path []string
	for from != to {
		path = append(path, from)
		from = orbits[from]
	}
	return path
}
