package util

import (
	"strconv"
	"strings"
)

func Atoi(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func ToInts(s string) []int {
	codes := strings.Split(s, ",")
	var ints []int
	for _, code := range codes {
		ints = append(ints, Atoi(code))
	}
	return ints
}
