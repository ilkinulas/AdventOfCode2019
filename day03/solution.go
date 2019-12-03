package day03

import (
	"AdventOfCode2019/util"
	"math"
	"strings"
)

type Direction string

type Position struct {
	x int
	y int
}

type Move struct {
	direction string
	distance  int
}

var coeffs = map[string]Position{
	"R": {x: 1, y: 0},
	"U": {x: 0, y: 1},
	"L": {x: -1, y: 0},
	"D": {x: 0, y: -1},
}

func solve(input1, input2 string) (uint, uint) {
	visited1 := moveAll(Position{}, input1)
	visited2 := moveAll(Position{}, input2)

	var intersections []Position
	for _, pos1 := range visited1 {
		for _, pos2 := range visited2 {
			if pos1.x == pos2.x && pos1.y == pos2.y {
				intersections = append(intersections, pos1)
				break
			}
		}
	}
	minDistance := ^uint(0) // max uint
	for _, pos := range intersections {
		distance := distance(pos)
		if distance < minDistance {
			minDistance = distance
		}
	}

	minSteps := ^uint(0) // max uint
	for _, intersection := range intersections {
		steps := calculateStepsTo(visited1, intersection) + calculateStepsTo(visited2, intersection)
		if steps < minSteps {
			minSteps = steps
		}
	}

	return minDistance, minSteps
}

func calculateStepsTo(path []Position, intersection Position) uint {
	for i, pos := range path {
		if pos.x == intersection.x && pos.y == intersection.y {
			return uint(i) + 1
		}
	}
	return 0
}

func distance(pos Position) uint {
	return uint(math.Abs(float64(pos.x))) + uint(math.Abs(float64(pos.y)))
}

func move(position Position, move Move) []Position {
	var visited []Position
	currentPos := position
	for i := 1; i <= move.distance; i++ {
		nextPos := Position{
			x: currentPos.x + coeffs[move.direction].x,
			y: currentPos.y + coeffs[move.direction].y,
		}
		visited = append(visited, nextPos)
		currentPos = nextPos
	}
	return visited
}

func moveAll(position Position, input string) []Position {
	var visited []Position
	var currentPos = position
	for _, s := range strings.Split(input, ",") {
		visited = append(visited, move(currentPos, toMove(s))...)
		currentPos = visited[len(visited)-1]
	}
	return visited
}

func toMove(s string) Move {
	return Move{
		direction: string(s[0]),
		distance:  util.Atoi(s[1:]),
	}
}
