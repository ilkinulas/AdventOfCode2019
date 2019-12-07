package util

import (
	"bufio"
	"os"
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

// ignore erros
func ReadLines(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
