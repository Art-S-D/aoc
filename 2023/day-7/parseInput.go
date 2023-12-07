package main

import (
	"bufio"
	"os"
	"strconv"
)

func parseInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, line)
	}
	return res
}

func MustAtoi(in string) int {
	res, err := strconv.Atoi(in)
	if err != nil {
		panic(err.Error())
	}
	return res
}
