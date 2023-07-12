package main

import (
	"fmt"

	"github.com/Art-S-D/aoc-2022-day-10/cpu"
)

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func main() {
	input, err := cpu.ParseInput("input.txt")
	if err != nil {
		panic(err.Error())
	}

	cpu := cpu.NewCpu(input)

	for !cpu.Done() {
		if cpu.Clock%40 == 0 {
			fmt.Println()
		}
		spriteCenter := cpu.Registers['X']
		shouldDrawPixel := abs(spriteCenter-int(cpu.Clock%40)) <= 1
		if shouldDrawPixel {
			fmt.Print("##")
		} else {
			fmt.Print("  ")
		}

		cpu.Tick()
	}
}
