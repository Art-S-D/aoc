package main

import (
	"fmt"
	"log"

	"github.com/Art-S-D/aoc-2022-day-11/monkey"
)

func main() {
	herd, err := monkey.ParseMonkeyHerd("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10_000; i++ {
		// fmt.Printf("round %d:\n", i+1)
		herd.Round()
	}

	fmt.Println(herd.MonkeyBusinessLevel())
}
