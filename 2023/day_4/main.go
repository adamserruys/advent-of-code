package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal("ERROR!!! No path to input file provided or part to run")
	}

	pathToInput := args[1]
	partToRun := args[2]

	f, err := os.Open(pathToInput)

	if err != nil {
		log.Fatal(err)
	}

	fs := bufio.NewScanner(f)

	var cards []string

	for fs.Scan() {
		cards = append(cards, fs.Text())
	}

	if partToRun == "one" {
		scoreSum := 0
		for _, card := range cards {
			score := findScoreOfCard(card)
			scoreSum += score
		}
		fmt.Println("Sum of scores: " + fmt.Sprint(scoreSum))
	} else if partToRun == "two" {
		cardsWon := findNumberOfCardsWon(cards)
		cardsWonSum := 0

		for _, numOfCards := range cardsWon {
			cardsWonSum += numOfCards
		}
		fmt.Println("Sum of cards won: " + fmt.Sprint(cardsWonSum))
	}

	f.Close()
}

func findNumberOfCardsWon(cards []string) []int {
	numOfCards := len(cards)
	cardCounter := make([]int, numOfCards)

	for cardIdx, card := range cards {
		cardCounter[cardIdx] += 1
		currentCount := cardCounter[cardIdx]
		numOfMatches := findNumOfMatches(card)
		stoppingPoint := int(math.Min(float64(cardIdx+numOfMatches), float64(numOfCards-1)))
		for i := cardIdx + 1; i < stoppingPoint+1; i++ {
			cardCounter[i] += currentCount
		}
	}

	return cardCounter
}

func findScoreOfCard(card string) int {
	numOfMatches := findNumOfMatches(card)

	if numOfMatches == 0 {
		return 0
	} else if numOfMatches == 1 {
		return 1
	} else {
		return int(math.Pow(2, float64(numOfMatches-1)))
	}

}

func findNumOfMatches(card string) int {
	numberOfMatches := 0
	idAndNumbers := strings.Split(card, ":")
	winnersAndNums := strings.Split(idAndNumbers[1], "|")

	winners := strings.Fields(winnersAndNums[0])
	gameNums := strings.Fields(winnersAndNums[1])

	for _, gn := range gameNums {
		for _, wn := range winners {
			if gn == wn {
				numberOfMatches += 1
			}
		}
	}

	return numberOfMatches
}
