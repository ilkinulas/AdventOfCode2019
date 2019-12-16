package day08

import (
	"AdventOfCode2019/util"
	"fmt"
	"regexp"
)

func solvePart1() int {
	layers := readLayers()

	zero := regexp.MustCompile("0")
	one := regexp.MustCompile("1")
	two := regexp.MustCompile("2")

	minZeros := util.MaxInt
	minZeroLayer := ""
	for _, layer := range layers {
		matches := zero.FindAllStringIndex(layer, -1)
		numZeros := len(matches)
		if numZeros < minZeros {
			minZeros = numZeros
			minZeroLayer = layer
		}
	}

	numOnes := len(one.FindAllStringIndex(minZeroLayer, -1))
	numTwos := len(two.FindAllStringIndex(minZeroLayer, -1))

	return numOnes * numTwos
}

func render() {
	layers := readLayers()
	var image []rune
	for i := 0; i < 150; i++ {
		image = append(image, pixelFromLayers(i, layers))
	}

	for r := 0; r < 6; r++ {
		for c := 0; c < 25; c++ {
			switch image[c+25*r] {
			case '0':
				fmt.Print(" ")
			case '1':
				fmt.Print("#")
			}
		}
		fmt.Print("\n")
	}

}

func pixelFromLayers(pos int, layers []string) rune {
	for _, layer := range layers {
		switch rune(layer[pos]) {
		case '0':
			return '0'
		case '1':
			return '1'
		case '2':
		}
	}
	return '0'
}

func readLayers() []string {
	pixels := util.ReadLines("input.txt")[0]
	totalPixels := len(pixels)
	pixelsPerLayer := 25 * 6
	ptr := 0
	var layers []string
	for ptr+pixelsPerLayer <= totalPixels {
		layers = append(layers, pixels[ptr:ptr+pixelsPerLayer])
		ptr += pixelsPerLayer
	}
	return layers
}
