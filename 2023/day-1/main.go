package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	res := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}
	return res
}

var numbers []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func findNumber(line string, at int) (hasNumber bool, value uint8) {
	if line[at] >= '0' && line[at] <= '9' {
		return true, uint8(line[at] - '0')
	}
	for i, number := range numbers {
		if strings.HasPrefix(line[at:], number) {
			return true, uint8(i + 1)
		}
	}
	return false, 0
}

func firstNumber(line string) uint8 {
	for i := range line {
		hasNumber, value := findNumber(line, i)
		if hasNumber {
			return value
		}
	}
	return 0
}
func lastNumber(line string) uint8 {
	lastNumber := uint8(0)
	for i := range line {
		hasNumber, value := findNumber(line, i)
		if hasNumber {
			lastNumber = value
		}
	}
	return lastNumber
}

func calibrationValue(line string) uint8 {
	return firstNumber(line)*10 + lastNumber(line)
}

func main() {
	input := parseInput()

	result := uint(0)
	for _, line := range input {
		result += uint(calibrationValue(line))
		// fmt.Println(i+1, uint(calibrationValue(line)))
	}
	fmt.Println("result", result)
}
