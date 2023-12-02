package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	targetSum := 0
	maxBlue := 14
	maxGreen := 13
	maxRed := 12

	for fs.Scan() {
		if partToRun == "two" {

			targetVal := getCubePower(fs.Text())
			targetSum += targetVal
		} else {

			targetVal := getEligibleId(fs.Text(), maxBlue, maxGreen, maxRed)
			targetSum += targetVal

		}

	}
	fmt.Printf("Sum: %d ", targetSum)

	f.Close()
}

func getCubePower(g string) int {
	parsedGame := parseGame(g)
	return parsedGame[1] * parsedGame[2] * parsedGame[3]
}

func getEligibleId(g string, maxBlue int, maxGreen int, maxRed int) int {
	parsedGame := parseGame(g)
	eligibleId := parsedGame[0]

	if parsedGame[1] > maxBlue || parsedGame[2] > maxGreen || parsedGame[3] > maxRed {
		eligibleId = 0
	}

	return eligibleId
}

// return an array of id, maxblue, maxgreen, maxred
func parseGame(g string) []int {
	gameId := 0
	maxBlue := 0
	maxGreen := 0
	maxRed := 0

	idAndResults := strings.Split(g, ":")
	if len(idAndResults) < 2 {
		return []int{0, 0, 0, 0}
	}
	gameIdStr := strings.Fields(idAndResults[0])[1]
	n, err := strconv.Atoi(gameIdStr)
	if err == nil {
		gameId = n
	}

	gameResults := strings.Split(idAndResults[1], ";")
	for _, r := range gameResults {
		colorResults := strings.Split(r, ",")
		for _, cr := range colorResults {

			cubesAndColor := strings.Fields(cr)

			numOfCubes, err := strconv.Atoi(cubesAndColor[0])
			if err != nil {
				numOfCubes = 0
			}
			color := cubesAndColor[1]

			switch color {
			case "blue":
				maxBlue = getMaxInt(numOfCubes, maxBlue)
			case "green":
				maxGreen = getMaxInt(numOfCubes, maxGreen)
			default:
				maxRed = getMaxInt(numOfCubes, maxRed)
			}
		}
	}

	return []int{gameId, maxBlue, maxGreen, maxRed}
}

func getMaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
