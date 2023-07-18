package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Valve struct {
	Ppm          uint8 // pressure per minute
	Label        string
	LeadsTo      []*Valve
	LeadsToLabel []string
}

var re *regexp.Regexp = regexp.MustCompile(`^Valve (\w+) has flow rate=(\d+); (tunnel|tunnels) (lead|leads) to (valve|valves) (.+)$`)

// returns the valve with the label AA
func ParseInput() (start *Valve, allValves []*Valve) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	valves := []*Valve{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if !re.MatchString(line) {
			panic(fmt.Sprintf("line <%s> does not match regexp", line))
		}

		submatches := re.FindStringSubmatch(line)
		var valve Valve
		valve.Label = submatches[1]
		_, err := fmt.Sscanf(submatches[2], "%d", &valve.Ppm)
		if err != nil {
			panic(err.Error())
		}
		valve.LeadsToLabel = strings.Split(submatches[6], ", ")
		valves = append(valves, &valve)
	}

	var res *Valve
	for i := range valves {
		valve := valves[i]
		if valve.Label == "AA" {
			res = valve
		}
		valve.LeadsTo = []*Valve{}
		for _, label := range valve.LeadsToLabel {
			for i := range valves {
				if valves[i].Label == label {
					valve.LeadsTo = append(valve.LeadsTo, valves[i])
				}
			}
		}
	}
	if res == nil {
		panic("valve with label AA not found")
	}
	return res, valves
}
