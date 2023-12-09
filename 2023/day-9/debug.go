package main

import "fmt"

func (v Values) offsetPrint(offset int) {
	for i := 0; i < offset; i++ {
		fmt.Printf(" ")
	}
	for _, val := range v {
		fmt.Printf(" %d", val)
	}
	fmt.Printf("\n")
}

func (v Values) PrettyPrint() {
	current := v
	offset := 0
	for !current.AllZeros() {
		current.offsetPrint(offset)
		offset++
		current = current.Differences()
	}
	current.offsetPrint(offset)
}
