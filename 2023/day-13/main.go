package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type Pattern []string

func ParsePatterns() []Pattern {
	var res []Pattern
	var currentPattern Pattern
	for _, line := range parseInput() {
		if len(line) > 0 {
			currentPattern = append(currentPattern, line)
		} else {
			res = append(res, currentPattern)
			currentPattern = nil
		}
	}
	if currentPattern != nil {
		res = append(res, currentPattern)
	}
	return res
}

func (p Pattern) IsHorizontalReflection(at int) bool {
	for i := 0; i < at; i++ {
		reflected := at*2 - i - 1
		if reflected >= len(p) {
			continue
		}
		if p[i] != p[reflected] {
			return false
		}
	}
	return true
}
func (p Pattern) IsVerticalReflection(at int) bool {
	for i := 0; i < at; i++ {
		reflected := at*2 - i - 1
		if reflected >= len(p[0]) {
			continue
		}
		for j := range p {
			if p[j][i] != p[j][reflected] {
				return false
			}
		}
	}
	return true
}

func (p Pattern) Summarize() int {
	for i := 1; i < len(p); i++ {
		if p.IsHorizontalReflection(i) {
			return i * 100
		}
	}
	for i := 1; i < len(p[0]); i++ {
		if p.IsVerticalReflection(i) {
			return i
		}
	}
	panic(fmt.Sprintf("no reflection found in pattern \n%+v\n", strings.Join(p, "\n")))
}

func (p Pattern) IsApproxHorizontalReflection(at int) bool {
	count := 0
	for i := 0; i < at; i++ {
		reflected := at*2 - i - 1
		if reflected >= len(p) {
			continue
		}
		for j := range p[i] {
			if p[i][j] != p[reflected][j] {
				count += 1
			}
		}
	}
	return count == 1
}
func (p Pattern) IsApproxVerticalReflection(at int) bool {
	count := 0
	for i := 0; i < at; i++ {
		reflected := at*2 - i - 1
		if reflected >= len(p[0]) {
			continue
		}
		for j := range p {
			if p[j][i] != p[j][reflected] {
				count += 1
			}
		}
	}
	return count == 1
}

func (p Pattern) SummarizeApprox() int {
	for i := 1; i < len(p); i++ {
		if p.IsApproxHorizontalReflection(i) {
			return i * 100
		}
	}
	for i := 1; i < len(p[0]); i++ {
		if p.IsApproxVerticalReflection(i) {
			return i
		}
	}
	panic(fmt.Sprintf("no approximal reflection found in pattern \n%+v\n", strings.Join(p, "\n")))
}

func main() {
	patterns := ParsePatterns()
	res := 0
	for _, pattern := range patterns {
		count := pattern.SummarizeApprox()
		fmt.Println(strings.Join(pattern, "\n"))
		fmt.Println(count)
		res += count
	}
	fmt.Println(res)
}
