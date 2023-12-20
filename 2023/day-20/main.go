package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func ParseModules() map[string]Module {
	var moduleRegex *regexp.Regexp = regexp.MustCompile(`([a-zA-Z0-9&%]+) -> (.*)`)
	matches := moduleRegex.FindAllStringSubmatch(input, -1)
	res := make(map[string]Module)
	for _, match := range matches {
		name := match[1]
		modules := match[2]
		var destinations []string
		destinations = append(destinations, strings.Split(modules, ", ")...)
		if name == "broadcaster" {
			broadcast := Broadcast{BaseModule{destinations: destinations, Name: "broadcaster"}}
			res["broadcaster"] = &broadcast
		} else if name[0] == '%' {
			flipflop := FlipFlop{BaseModule{destinations: destinations, Name: name[1:]}, false}
			res[name[1:]] = &flipflop
		} else if name[0] == '&' {
			conjunction := Conjunction{BaseModule{destinations: destinations, Name: name[1:]}, make(map[string]Height)}
			res[name[1:]] = &conjunction
		} else {
			panic(fmt.Sprintf("unknown module %s", name))
		}
	}
	// update all conjunctions to remember low pulse on all sources by default
	for _, v := range res {
		if conjunction, ok := v.(*Conjunction); ok {
			for name, v := range res {
				if slices.Contains(v.Destinations(), conjunction.Name) {
					conjunction.MostRecentPulse[name] = Low
				}
			}
		}
	}
	res["button"] = &Button{BaseModule{Name: "button", destinations: []string{"broadcaster"}}}
	return res
}

func main() {
	modules := ParseModules()
	for _, v := range modules {
		fmt.Printf("%+v\n", v)
	}
	// lowPulses := 0
	// highPulses := 0
	presses := 0
	for {
		pulses := []Pulse{{Source: "button", Destination: "broadcaster", Height: Low}}
		// fmt.Println("\n=== pressing the button ===")
		presses += 1
		if presses%100_000 == 0 {
			fmt.Println(presses)
		}
		for len(pulses) > 0 {
			pulse := pulses[0]
			// fmt.Println(pulse)

			if pulse.Destination == "rx" && pulse.Height == Low {
				fmt.Println(presses, "presses")
				return
			}
			if pulse.Destination == "rs" && pulse.Height == High {
				// with this, find the cycles that send high to the modules before rx and the lcm of the cycles is the answer
				fmt.Println(pulse, presses)
			}

			// if pulse.Height == Low {
			// 	lowPulses += 1
			// } else {
			// 	highPulses += 1
			// }

			destModule, ok := modules[pulse.Destination]
			if !ok {
				// because of examples that point to a non defined module
				pulses = pulses[1:]
				continue
			}
			newPulses := destModule.Receive(pulse)
			pulses = append(pulses[1:], newPulses...)
		}
	}
	// fmt.Printf("%d low, %d heigh: %d pulses\n", lowPulses, highPulses, lowPulses*highPulses)
}
