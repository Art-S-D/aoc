package main

import (
	"bufio"
	"os"
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
