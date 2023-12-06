package day01

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/selece/aoc2023-golang/util"
)

func RunDay01(path string) {
	reader := util.NewFileReader(path)
	log.Printf("sum: %d\n", tallyNumbersFromStringArray(reader.Contents))
}

func tallyNumbersFromStringArray(input []string) int {
	tally := 0
	for _, line := range input {
		tally += processNumbersFromString(line)
	}

	return tally
}

func processNumbersFromString(input string) int {
	return crunchNumericStringArrayToInt(findNumbers(replaceStringsWithNumbers(input)))
}

var replacer = strings.NewReplacer(
	"one", "on1e",
	"two", "tw2o",
	"three", "thre3e",
	"four", "fou4r",
	"five", "fiv5e",
	"six", "si6x",
	"seven", "seve7n",
	"eight", "eigh8t",
	"nine", "nin9e",
)

func replaceStringsWithNumbers(input string) string {
	result := input
	for {
		prev := result
		result = replacer.Replace(result)

		if prev == result {
			break
		}
	}

	return result
}

func findNumbers(input string) []string {
	re := regexp.MustCompile("[1-9]")
	return re.FindAllString(input, -1)
}

func crunchNumericStringArrayToInt(input []string) int {
	crunched := ""
	crunched += string(input[0])
	crunched += string(input[len(input)-1])

	converted, err := strconv.Atoi(crunched)
	if err != nil {
		log.Fatalf("failed to convert to string: %s\n", crunched)
	}

	return converted
}
