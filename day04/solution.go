package day04

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func countPossiblePasswords(min, max int) (int, int) {
	countPart1 := 0
	countPart2 := 0
	for i := min; i <= max; i++ {
		if isValidPart1(i) {
			countPart1++
		}
		if isValidPart2(i) {
			countPart2++
		}
	}
	return countPart1, countPart2
}

func isValidPart1(pass int) bool {
	s := strconv.Itoa(pass)

	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	digits := strings.Split(s, "")
	// is increasing?
	if !reflect.DeepEqual(sorted, digits) {
		return false
	}

	digitCounts := make(map[string]int)
	for _, digit := range digits {
		if _, ok := digitCounts[digit]; !ok {
			digitCounts[digit] = 1
		} else {
			digitCounts[digit] = digitCounts[digit] + 1
		}
	}
	for _, value := range digitCounts {
		if value > 1 {
			return true
		}
	}
	return false
}

func isValidPart2(pass int) bool {
	s := strconv.Itoa(pass)

	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	digits := strings.Split(s, "")
	// is increasing?
	if !reflect.DeepEqual(sorted, digits) {
		return false
	}

	digitCounts := make(map[string]int)
	for _, digit := range digits {
		if _, ok := digitCounts[digit]; !ok {
			digitCounts[digit] = 1
		} else {
			digitCounts[digit] = digitCounts[digit] + 1
		}
	}
	for _, value := range digitCounts {
		if value == 2 {
			return true
		}
	}
	return false
}
