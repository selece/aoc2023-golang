package day02

import (
	"log"
	"regexp"
	"strconv"

	"github.com/selece/aoc2023-golang/util"
)

type GameRound struct {
	red   int
	blue  int
	green int
}

func (gr *GameRound) SetColourData(colour string, count int) {
	switch colour {
	case "red":
		gr.red = count

	case "blue":
		gr.blue = count

	case "green":
		gr.green = count

	default:
		log.Fatalf("invalid colour field: %s", colour)
	}
}

type Game struct {
	id     int
	rounds []GameRound

	maxRed   int
	maxGreen int
	maxBlue  int
}

func NewGame(id int) *Game {
	g := Game{id: id}
	return &g
}

func (g *Game) AddRound(round GameRound) {
	g.rounds = append(g.rounds, round)

	if round.blue > g.maxBlue {
		g.maxBlue = round.blue
	}

	if round.green > g.maxGreen {
		g.maxGreen = round.green
	}

	if round.red > g.maxRed {
		g.maxRed = round.red
	}
}

func (g *Game) PassesRequirements(maxRed int, maxGreen int, maxBlue int) bool {
	return g.maxBlue <= maxBlue && g.maxGreen <= maxGreen && g.maxRed <= maxRed
}

func RunDay02(path string) {
	reader := util.NewFileReader(path)

	reqRed := 12
	reqGreen := 13
	reqBlue := 14

	tally := 0

	for _, line := range reader.Contents {
		game := extractGameFromLine(line)
		if game.PassesRequirements(reqRed, reqGreen, reqBlue) {
			tally += game.id
		}
	}

	log.Printf("sum: %d", tally)
}

func extractGameId(input string) int {
	re := regexp.MustCompile(`Game (?P<gameId>\d+)`)
	match := re.FindStringSubmatch(input)

	if match == nil {
		log.Fatalf("no gameId match for input: %s", input)
	}

	converted, err := strconv.Atoi(match[re.SubexpIndex("gameId")])
	if err != nil {
		log.Fatal(err)
	}

	return converted
}

var gameIdFromDataString = regexp.MustCompile(`: `)
var gameDataSetSeperator = regexp.MustCompile(`;\s?`)
var colourDataSeperator = regexp.MustCompile(`,\s?`)
var colourDataParser = regexp.MustCompile(`(?P<count>\d+) (?P<colour>(blue|green|red))`)

func extractGameRounds(input string) []GameRound {
	dataString := gameIdFromDataString.Split(input, -1)[1]
	dataStringSets := gameDataSetSeperator.Split(dataString, -1)

	var rounds []GameRound

	for _, subsplit := range dataStringSets {
		colourSets := colourDataSeperator.Split(subsplit, -1)
		round := GameRound{}

		for _, colourSplit := range colourSets {
			colourData := colourDataParser.FindStringSubmatch(colourSplit)
			colour := colourData[colourDataParser.SubexpIndex("colour")]
			count, err := strconv.Atoi(colourData[colourDataParser.SubexpIndex("count")])

			if err != nil {
				log.Fatalf("failed to convert to int: %s", colourData[colourDataParser.SubexpIndex("count")])
			}

			round.SetColourData(colour, count)
		}

		rounds = append(rounds, round)
	}

	return rounds
}

func extractGameFromLine(input string) *Game {
	gameId := extractGameId(input)
	rounds := extractGameRounds(input)

	game := NewGame(gameId)

	for _, r := range rounds {
		game.AddRound(r)
	}

	return game
}
