package main

import (
	"log"
	"os"

	"github.com/selece/aoc2023-golang/day01"
	"github.com/selece/aoc2023-golang/day02"
)

func main() {
	day := os.Args[1]
	path := os.Args[2]

	switch day {
	case "1":
		day01.RunDay01(path)

	case "2":
		day02.RunDay02(path)

	default:
		log.Fatal("invalid day choice: " + day)
	}
}
