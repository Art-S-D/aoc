package monkey

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type operation interface {
	apply(old uint) (new uint)
}
type addition struct{ by uint }

func (a addition) apply(old uint) (new uint) { return old + a.by }

type multiplication struct{ by uint }

func (m multiplication) apply(old uint) (new uint) { return old * m.by }

type square struct{}

func (s square) apply(old uint) (new uint) { return old * old }

func parseOperation(scanner *bufio.Scanner) (operation, error) {
	scanner.Scan()
	line := scanner.Text()

	var operator rune
	var rhs string

	_, err := fmt.Sscanf(line, "  Operation: new = old %c %s", &operator, &rhs)
	if err != nil {
		return nil, err
	}

	if rhs == "old" {
		return &square{}, nil
	} else {
		var by uint
		_, err = fmt.Sscanf(rhs, "%d", &by)
		if err != nil {
			return nil, err
		}
		switch operator {
		case '*':
			return &multiplication{by}, nil
		case '+':
			return &addition{by}, nil
		default:
			return nil, fmt.Errorf("unknown operation %c", operator)
		}
	}
}

type test struct {
	divisibleBy    uint
	ifTrueThrowTo  uint
	ifFalseThrowTo uint
}

func parseTest(scanner *bufio.Scanner) (test, error) {
	var res test

	scanner.Scan()
	testLine := scanner.Text()
	_, err := fmt.Sscanf(testLine, "  Test: divisible by %d", &res.divisibleBy)
	if err != nil {
		return res, err
	}

	scanner.Scan()
	trueLine := scanner.Text()
	_, err = fmt.Sscanf(trueLine, "    If true: throw to monkey %d", &res.ifTrueThrowTo)
	if err != nil {
		return res, err
	}

	scanner.Scan()
	falseLine := scanner.Text()
	_, err = fmt.Sscanf(falseLine, "    If false: throw to monkey %d", &res.ifFalseThrowTo)
	if err != nil {
		return res, err
	}

	return res, nil
}

type MonkeyHerd []Monkey
type Monkey struct {
	holdings  []uint
	inspected uint
	operation operation
	test      test
}

func (m *Monkey) InspectOne(divProduct uint) (worryLevel uint, throwTo uint) {
	m.inspected += 1

	firstItem := m.holdings[0]
	m.holdings = m.holdings[1:]

	firstItem = m.operation.apply(firstItem)

	// MONKEY BOREDOM
	// firstItem /= 3

	if firstItem%m.test.divisibleBy == 0 {
		return firstItem % divProduct, m.test.ifTrueThrowTo
	} else {
		return firstItem % divProduct, m.test.ifFalseThrowTo
	}
}

func (herd MonkeyHerd) Round() {
	var divProduct uint = 1
	for _, monkey := range herd {
		divProduct *= monkey.test.divisibleBy
	}

	for i := range herd {
		monkey := &herd[i]
		for len(monkey.holdings) > 0 {
			worryLevel, throwTo := monkey.InspectOne(divProduct)
			herd[throwTo].holdings = append(herd[throwTo].holdings, worryLevel)
		}
	}
}

func (herd MonkeyHerd) MonkeyBusinessLevel() uint {
	var max1 uint
	var max2 uint

	for _, monkey := range herd {
		if monkey.inspected > max1 {
			max2 = max1
			max1 = monkey.inspected
		} else if monkey.inspected > max2 {
			max2 = monkey.inspected
		}
	}

	fmt.Printf("most active: %d and %d\n", max1, max2)

	return max1 * max2
}

func parseMonkey(scanner *bufio.Scanner) (*Monkey, int, error) {
	// scanner.Scan()
	firstLine := scanner.Text()
	var monkeyIndex int
	_, err := fmt.Sscanf(firstLine, "Monkey %d:", &monkeyIndex)
	if err != nil {
		return nil, -1, err
	}

	var monkey Monkey

	scanner.Scan()
	startingItems := scanner.Text()
	startingItems = startingItems[len("  Starting items: "):]
	for _, item := range strings.Split(startingItems, ", ") {
		var holding uint
		_, err := fmt.Sscanf(item, "%d", &holding)
		if err != nil {
			return nil, monkeyIndex, err
		}
		monkey.holdings = append(monkey.holdings, holding)
	}

	op, err := parseOperation(scanner)
	if err != nil {
		return nil, monkeyIndex, err
	}
	monkey.operation = op

	test, err := parseTest(scanner)
	if err != nil {
		return nil, monkeyIndex, err
	}
	monkey.test = test

	return &monkey, monkeyIndex, nil
}

func ParseMonkeyHerd(filename string) (MonkeyHerd, error) {
	var res MonkeyHerd

	inputFile, err := os.Open(filename)
	if err != nil {
		return res, err
	}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		monkey, _, err := parseMonkey(scanner)
		if err != nil {
			return res, err
		}
		res = append(res, *monkey)

		scanner.Scan()
		if scanner.Text() != "" {
			return res, fmt.Errorf("missing blank line after parsing a monkey")
		}
	}

	return res, nil
}
