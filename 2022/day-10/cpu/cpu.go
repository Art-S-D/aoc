package cpu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cpu struct {
	Registers               map[rune]int
	Clock                   uint
	instructions            []Instruction
	currentInstructionDelay uint
}

func NewCpu(instructions []Instruction) Cpu {
	res := Cpu{
		Clock:        0,
		Registers:    make(map[rune]int),
		instructions: instructions,
	}
	res.Registers['X'] = 1
	return res
}
func (c *Cpu) Run(instruction Instruction) {
	instruction.apply(c)
}
func (c *Cpu) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("clock %d\n", c.Clock))
	for k, v := range c.Registers {
		sb.WriteString(fmt.Sprintf("\t%c: %d\n", k, v))
	}
	return sb.String()
}
func (c *Cpu) CurrentInstruction() Instruction {
	return c.instructions[0]
}
func (c *Cpu) Done() bool {
	return len(c.instructions) == 0
}
func (c *Cpu) Tick() {
	c.Clock += 1
	c.currentInstructionDelay += 1
	if c.currentInstructionDelay >= c.CurrentInstruction().ClockTicks() {
		c.currentInstructionDelay = 0
		c.Run(c.CurrentInstruction())
		c.instructions = c.instructions[1:]
	}
}

type Noop struct{}

func (n Noop) apply(cpu *Cpu) {
	// cpu.Clock += n.ClockTicks()
}
func (n Noop) String() string {
	return "noop"
}
func (n Noop) ClockTicks() uint { return 1 }

type AddX struct{ Value int }

func (a AddX) apply(cpu *Cpu) {
	cpu.Registers['X'] += a.Value
	// cpu.Clock += a.ClockTicks()
}
func (a AddX) ClockTicks() uint { return 2 }
func (a AddX) String() string {
	return fmt.Sprintf("addx %d", a.Value)
}

type Instruction interface {
	apply(cpu *Cpu)
	ClockTicks() uint
}

func instructionFromString(line string) (Instruction, error) {
	switch line[:4] {
	case "noop":
		return Noop{}, nil
	case "addx":
		var add AddX
		_, err := fmt.Sscanf(line[5:], "%d", &add.Value)
		if err != nil {
			return nil, err
		}
		return add, nil
	default:
		return nil, fmt.Errorf("unknown instruction %s", line)
	}

}

func ParseInput(file string) ([]Instruction, error) {
	var input []Instruction

	inputFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}
	buf := bufio.NewScanner(inputFile)
	for buf.Scan() {
		line := buf.Text()
		instruction, err := instructionFromString(line)
		if err != nil {
			return nil, err
		}
		input = append(input, instruction)
	}

	return input, nil
}
