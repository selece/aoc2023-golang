package main

import (
	"log"
	"os"

	"github.com/selece/aoc2023-golang/day01"
)

func main() {
	day := os.Args[1]
	short := false

	if os.Args[2] == "true" {
		short = true
	} else {
		short = false
	}

	switch day {
	case "1":
		day01.RunDay01(short)

	default:
		log.Fatal("invalid day choice: " + day)
	}
}
