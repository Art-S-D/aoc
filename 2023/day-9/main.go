package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Values []int
type History []Values

func (v Values) Differences() Values {
	var res Values
	for i := 0; i < len(v)-1; i++ {
		res = append(res, v[i+1]-v[i])
	}
	return res
}

func (v Values) AllZeros() bool {
	for _, val := range v {
		if val != 0 {
			return false
		}
	}
	return true
}

func (v Values) Extrapolate() int {
	if v.AllZeros() {
		return 0
	}
	lastValue := v[len(v)-1]
	return lastValue + v.Differences().Extrapolate()
}

func (v Values) ExtrapolateBackwards() int {
	if v.AllZeros() {
		return 0
	}
	firstValue := v[0]
	return firstValue - v.Differences().ExtrapolateBackwards()
}

func parseHistory() History {
	var res History
	for _, line := range parseInput() {
		var values Values
		for _, value := range strings.Fields(line) {
			v, _ := strconv.Atoi(value)
			values = append(values, v)
		}
		res = append(res, values)
	}
	return res
}

func main() {
	history := parseHistory()

	res := 0
	for _, values := range history {
		// fmt.Println(values.ExtrapolateBackwards())
		// values.PrettyPrint()
		res += values.ExtrapolateBackwards()
	}
	fmt.Println(res)
}
