package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type pn struct {
	id       string
	startIdx int
	endIdx   int
}

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

	var engineSchema []string

	for fs.Scan() {
		engineSchema = append(engineSchema, fs.Text())
	}

	if partToRun == "one" {
		partNumbers := findEngineParts(engineSchema)
		fmt.Println("Sum of part numbers: " + fmt.Sprint(sumUpParts(partNumbers)))
	} else if partToRun == "two" {
		gearRatios := findGearRatios(engineSchema)
		fmt.Println("Sum of gear ratios: " + fmt.Sprint(sumUpParts(gearRatios)))
	}

	f.Close()
}

func sumUpParts(partNums []int) int {
	partNumberSum := 0

	for _, pn := range partNums {
		partNumberSum += pn
	}
	return partNumberSum

}

func findGearRatios(engineSchema []string) []int {
	var gearRatios []int
	for rowIdx, r := range engineSchema {
		for charIdx, char := range r {
			if string(char) == "*" {
				surroundingParts := findSurroundingParts(rowIdx, charIdx, engineSchema)
				if len(surroundingParts) == 2 {
					gearRatios = append(gearRatios, surroundingParts[0]*surroundingParts[1])
				}
			}

		}

	}
	return gearRatios
}

func findSurroundingParts(row int, gearIndex int, engineSchema []string) []int {
	var partNums []int

	if row > 0 {
		adj := findAdjPartNums(engineSchema[row-1], gearIndex)
		for _, n := range adj {
			partNums = append(partNums, n)
		}
	}

	if row < len(engineSchema)-1 {
		adj := findAdjPartNums(engineSchema[row+1], gearIndex)
		for _, n := range adj {
			partNums = append(partNums, n)
		}
	}

	adj := findAdjPartNums(engineSchema[row], gearIndex)
	for _, n := range adj {
		partNums = append(partNums, n)
	}

	return partNums

}

func findAdjPartNums(row string, gearIdx int) []int {
	var partNums []int
	i := gearIdx - 1
	for {
		if i >= 0 && unicode.IsDigit(rune(row[i])) {
			i -= 1
		} else {
			i += 1
			break
		}
	}
	currentNumber := ""
	for {
		if i > len(row)-1 {
			break
		}
		currentChar := row[i]
		if unicode.IsDigit(rune(currentChar)) {
			currentNumber += string(currentChar)
		} else if i != gearIdx {
			break
		} else {
			pn, err := strconv.Atoi(currentNumber)
			if err == nil {
				partNums = append(partNums, pn)
			}
			currentNumber = ""
		}

		i += 1
	}

	pn, err := strconv.Atoi(currentNumber)
	if err == nil {
		partNums = append(partNums, pn)
	}

	return partNums

}

func findEngineParts(engineSchema []string) []int {
	var partNumbers []int
	for rowIdx, r := range engineSchema {
		var potentialPartNumbers []pn
		currentNumber := ""
		startIdx := -1
		endIdx := -1
		for charIdx, char := range r {
			if unicode.IsDigit(char) {
				currentNumber += string(char)
				if startIdx == -1 {
					startIdx = charIdx
					endIdx = charIdx
				} else {
					endIdx = charIdx
				}
			} else {
				if currentNumber != "" {
					potentialPartNumbers = append(potentialPartNumbers, pn{id: currentNumber, startIdx: startIdx, endIdx: endIdx})
				}

				currentNumber = ""
				startIdx = -1
				endIdx = -1
			}
		}

		if currentNumber != "" {
			potentialPartNumbers = append(potentialPartNumbers, pn{id: currentNumber, startIdx: startIdx, endIdx: endIdx})
		}

		for _, ppn := range potentialPartNumbers {
			if isPartNumber(rowIdx, ppn.startIdx, ppn.endIdx, engineSchema) {
				pn, err := strconv.Atoi(ppn.id)
				if err == nil {
					partNumbers = append(partNumbers, pn)
				}
			}
		}
	}
	return partNumbers
}

func isPartNumber(row int, start int, end int, engineSchema []string) bool {
	startIdx := start - 1
	endIdx := end + 1
	wordAbove := ""
	wordBelow := ""

	if startIdx < 0 {
		startIdx = start
	}

	if endIdx > len(engineSchema[0])-1 {
		endIdx = end
	}

	if row > 0 {
		wordAbove = string(engineSchema[row-1][startIdx : endIdx+1])
	}
	if row < len(engineSchema)-1 {
		wordBelow = string(engineSchema[row+1][startIdx : endIdx+1])
	}

	return isSpecialChar(rune(engineSchema[row][startIdx])) || isSpecialChar(rune(engineSchema[row][endIdx])) || containsSpecialChar(wordAbove) || containsSpecialChar(wordBelow)
}

func containsSpecialChar(s string) bool {
	for _, char := range s {
		if isSpecialChar(char) {
			return true
		}
	}
	return false
}

func isSpecialChar(char rune) bool {
	convertedChar := string(char)
	return !unicode.IsDigit(char) && convertedChar != "."
}
