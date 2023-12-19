package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type WorkFlowRule struct {
	Category   byte
	Comparator byte
	Value      int
	SendTo     string
}
type Workflow struct {
	Name  string
	Rules []WorkFlowRule
}
type Part struct {
	X int
	M int
	A int
	S int
}

//go:embed input.txt
var input string

var workflowRegex *regexp.Regexp = regexp.MustCompile(`(\w+){(.*)}`)
var ruleRegex *regexp.Regexp = regexp.MustCompile(`(\w+)([><])(\d+):(\w+)`)

func ParseWorkflows() map[string]Workflow {
	res := make(map[string]Workflow)
	for _, workflowMatch := range workflowRegex.FindAllStringSubmatch(input, -1) {
		workflow := Workflow{Name: workflowMatch[1]}
		for _, r := range strings.Split(workflowMatch[2], ",") {
			ruleMatch := ruleRegex.FindStringSubmatch(r)
			if len(ruleMatch) > 0 {
				value, _ := strconv.Atoi(ruleMatch[3])
				rule := WorkFlowRule{
					Category:   ruleMatch[1][0],
					Comparator: ruleMatch[2][0],
					Value:      value,
					SendTo:     ruleMatch[4],
				}
				workflow.Rules = append(workflow.Rules, rule)
			} else {
				workflow.Rules = append(workflow.Rules, WorkFlowRule{Category: 'a', Comparator: '-', SendTo: r})
			}
		}
		res[workflow.Name] = workflow
	}
	return res
}

var partRegex *regexp.Regexp = regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)

func ParseParts() []Part {
	var res []Part
	for _, partMatch := range partRegex.FindAllStringSubmatch(input, -1) {
		x, _ := strconv.Atoi(partMatch[1])
		m, _ := strconv.Atoi(partMatch[2])
		a, _ := strconv.Atoi(partMatch[3])
		s, _ := strconv.Atoi(partMatch[4])
		part := Part{X: x, M: m, A: a, S: s}
		res = append(res, part)
	}
	return res
}

func (r WorkFlowRule) String() string {
	if r.Comparator == '-' {
		return r.SendTo
	}
	return fmt.Sprintf("%c%c%d:%s", r.Category, r.Comparator, r.Value, r.SendTo)
}
func (w Workflow) String() string {
	res := w.Name + "{"
	for _, rule := range w.Rules {
		res += rule.String()
		res += ","
	}
	return res + "}"
}
func (p Part) Category(categ byte) int {
	switch categ {
	case 'x':
		return p.X
	case 'm':
		return p.M
	case 'a':
		return p.A
	case 's':
		return p.S
	}
	panic(fmt.Sprintf("unknown category %v\n", categ))
}
func (w Workflow) ProcessPart(part Part) (arrivesAt string) {
	for _, rule := range w.Rules {
		category := part.Category(rule.Category)
		if rule.Comparator == '<' && category < rule.Value {
			return rule.SendTo
		} else if rule.Comparator == '>' && category > rule.Value {
			return rule.SendTo
		} else if rule.Comparator == '-' {
			return rule.SendTo
		}
	}
	panic("no rule found")
}
func ProcessPart(part Part, workflows map[string]Workflow) (accepted bool) {
	currentWorkflow := "in"
	for currentWorkflow != "A" && currentWorkflow != "R" {
		fmt.Printf("-> %s", currentWorkflow)
		currentWorkflow = workflows[currentWorkflow].ProcessPart(part)
	}
	fmt.Printf("-> %s\n", currentWorkflow)
	return currentWorkflow == "A"
}

type PartRange struct {
	X [2]int
	M [2]int
	A [2]int
	S [2]int
}

func (p PartRange) Category(categ byte) [2]int {
	switch categ {
	case 'x':
		return p.X
	case 'm':
		return p.M
	case 'a':
		return p.A
	case 's':
		return p.S
	}
	panic(fmt.Sprintf("unknown get category %v\n", categ))
}
func (p *PartRange) SetCategory(categ byte, value [2]int) {
	switch categ {
	case 'x':
		p.X = value
	case 'm':
		p.M = value
	case 'a':
		p.A = value
	case 's':
		p.S = value
	}
}

func (r PartRange) Process(workflows map[string]Workflow, at string) (accepted []PartRange) {
	if at == "R" {
		return nil
	} else if at == "A" {
		return []PartRange{r}
	}
	current := workflows[at]
	for _, rule := range current.Rules {
		category := r.Category(rule.Category)
		if rule.Category == '-' {
			return r.Process(workflows, rule.SendTo)
		}
		if rule.Comparator == '<' {
			if category[1] < rule.Value {
				return r.Process(workflows, rule.SendTo)
			} else if category[0] >= rule.Value {
				continue
			} else {
				r1 := r
				r1.SetCategory(rule.Category, [2]int{category[0], rule.Value - 1})
				r2 := r
				r2.SetCategory(rule.Category, [2]int{rule.Value, category[1]})
				return append(r1.Process(workflows, at), r2.Process(workflows, at)...)
			}
		}
		if rule.Comparator == '>' {
			if category[0] > rule.Value {
				return r.Process(workflows, rule.SendTo)
			} else if category[1] <= rule.Value {
				continue
			} else {
				r1 := r
				r1.SetCategory(rule.Category, [2]int{category[0], rule.Value})
				r2 := r
				r2.SetCategory(rule.Category, [2]int{rule.Value + 1, category[1]})
				return append(r1.Process(workflows, at), r2.Process(workflows, at)...)
			}
		}
		if rule.Comparator == '-' {
			return r.Process(workflows, rule.SendTo)
		}
	}
	fmt.Println(current)
	panic("no rule to apply found")
}

func (r PartRange) Possibilities() int {
	return ((r.X[1] + 1) - r.X[0]) * ((r.M[1] + 1) - r.M[0]) * ((r.A[1] + 1) - r.A[0]) * ((r.S[1] + 1) - r.S[0])
}

func main() {
	workflows := ParseWorkflows()
	// parts := ParseParts()

	// res := 0
	// for _, part := range parts {
	// 	fmt.Printf("%+v", part)
	// 	if ProcessPart(part, workflows) {
	// 		res += part.X + part.M + part.A + part.S
	// 	}
	// }
	// fmt.Println(res)

	initialRange := PartRange{
		X: [2]int{1, 4000},
		M: [2]int{1, 4000},
		A: [2]int{1, 4000},
		S: [2]int{1, 4000},
	}
	finalRanges := initialRange.Process(workflows, "in")

	res := 0
	for _, r := range finalRanges {
		res += r.Possibilities()
	}
	fmt.Println(res)
}
