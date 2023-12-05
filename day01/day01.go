package day01

import (
	"log"
	"strconv"
	"unicode"

	"github.com/selece/aoc2023-golang/util"
)

func RunDay01(short bool) {
	path := ""
	if short {
		path = "day01/test.txt"
	} else {
		path = "day01/input.txt"
	}
	reader := util.New(path)

	tally := make([]int, 0)
	for _, line := range reader.Contents {
		extracted := ""
		log.Printf("line: %s\n", line)

		for _, rune := range line {
			if unicode.IsDigit(rune) {
				extracted += string(rune)
			}
		}

		log.Printf("extracted: %s\n", extracted)

		crunched := ""
		crunched += string(extracted[0])
		crunched += string(extracted[len(extracted)-1])

		log.Printf("crunched: %s\n", crunched)

		converted, err := strconv.Atoi(crunched)
		if err != nil {
			log.Fatal("Failed to convert to integer: " + crunched)
			log.Fatal(err)
		}

		log.Printf("converted: %d\n", converted)
		tally = append(tally, converted)
	}

	sum := 0
	for _, v := range tally {
		sum += v
	}

	println(sum)
}
