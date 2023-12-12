package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SpringCondition []rune
type ContiguousGroups []uint
type SpringRow struct {
	Condition        SpringCondition
	ContiguousGroups ContiguousGroups
}

func (g ContiguousGroups) Sum() int {
	count := 0
	for _, n := range g {
		count += int(n)
	}
	return count
}
func (g ContiguousGroups) MinimumSizeNeeded() int {
	if len(g) == 0 {
		return 0
	}
	if len(g) == 1 {
		return int(g[0])
	}
	// len(g) - 1 accounts for spaces
	return g.Sum() + len(g) - 1
}

func Valid(cond SpringCondition, groups ContiguousGroups) bool {
	fields := strings.FieldsFunc(string(cond), func(c rune) bool { return c == '.' })
	if len(fields) != len(groups) {
		return false
	}
	for i := range fields {
		if uint(len(fields[i])) != groups[i] {
			return false
		}
	}
	return true
}

func (c SpringCondition) CanFitGroup(size int) bool {
	if size > len(c) {
		return false
	}
	for i := 0; i < size; i++ {
		if c[i] == '.' {
			return false
		}
	}
	return true
}

type cacheEntry struct {
	cond   string
	groups int
}

var cache map[cacheEntry]int

func (r SpringRow) CountArrangements() int {
	entry := cacheEntry{cond: string(r.Condition), groups: len(r.ContiguousGroups)}
	if _, ok := cache[entry]; ok {
		return cache[entry]
	}
	cond := r.Condition

	if len(r.ContiguousGroups) == 0 {
		if strings.Contains(string(cond), "#") {
			cache[entry] = 0
			return 0
		} else {
			cache[entry] = 1
			return 1
		}
	}
	if len(cond) < int(r.ContiguousGroups.MinimumSizeNeeded()) {
		cache[entry] = 0
		return 0
	}
	sum := r.ContiguousGroups.Sum()
	// unknownCount := strings.Count(string(cond), "?")
	brokenCount := strings.Count(string(cond), ".")
	workingCount := strings.Count(string(cond), "#")
	if len(cond)-brokenCount < sum {
		cache[entry] = 0
		return 0
	}
	if workingCount > sum {
		cache[entry] = 1
		return 0
	}

	res := 0
	groupSize := r.ContiguousGroups[0]

	if cond.CanFitGroup(int(groupSize)) {
		if int(groupSize) == len(cond) {
			cache[entry] = 1
			return 1
		} else if cond[groupSize] != '#' {
			nextRow := SpringRow{Condition: cond[groupSize+1:], ContiguousGroups: r.ContiguousGroups[1:]}
			res += nextRow.CountArrangements()
		}
	}
	if cond[0] == '#' {
		cache[entry] = res
		return res
	} else {
		nextRow := SpringRow{Condition: cond[1:], ContiguousGroups: r.ContiguousGroups}
		res += nextRow.CountArrangements()
		cache[entry] = res
		return res
	}
}

func parseSpringRows() []SpringRow {
	var res []SpringRow
	for _, line := range parseInput() {
		fields := strings.Fields(line)
		var groups ContiguousGroups
		for _, s := range strings.Split(fields[1], ",") {
			n, _ := strconv.Atoi(s)
			groups = append(groups, uint(n))
		}

		row := SpringRow{Condition: []rune(fields[0]), ContiguousGroups: groups}
		for i := 0; i < 4; i++ {
			row.Condition = []rune(string(row.Condition) + "?" + fields[0])
			row.ContiguousGroups = append(row.ContiguousGroups, groups...)
		}
		res = append(res, row)
	}
	return res
}

func main() {
	rows := parseSpringRows()
	res := 0
	for i, row := range rows {
		cache = make(map[cacheEntry]int)
		count := row.CountArrangements()
		fmt.Printf("%d/%d, %v %v %v\n", i, len(rows), string(row.Condition), row.ContiguousGroups, count)

		res += count
	}
	fmt.Println(res)
}
