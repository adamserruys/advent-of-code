package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
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

	calSum := int64(0)

	for fs.Scan() {
		var calStr string

		if partToRun == "two" {
			calStr = getPartTwoCalibrationString(fs.Text())
		} else {
			calStr = getPartOneCalibrationString(fs.Text())
		}

		calVal := calculateCalibrationValue(calStr)

		calSum += calVal
	}
	fmt.Println("Sum of calibration values is ", calSum)

	f.Close()
}

func getPartOneCalibrationString(s string) string {
	calStr := ""
	for _, char := range s {
		if unicode.IsDigit(char) {
			calStr += string(char)
		}
	}
	return calStr
}

func getPartTwoCalibrationString(s string) string {
	digitWords := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sLength := len(s)
	hasWordIndex := make([][]int, 9)
	populatedDigits := make([]string, sLength)
	calStr := ""

	for i := range hasWordIndex {
		hasWordIndex[i] = make([]int, 0, sLength)
	}

	for i, dw := range digitWords {
		hasWordIndex[i] = findWordIndexes(s, dw)

	}

	for i, char := range s {
		if unicode.IsDigit(char) {
			populatedDigits[i] = string(char)
		}
	}

	for i, wIdxs := range hasWordIndex {
		for _, idx := range wIdxs {
			populatedDigits[idx] = fmt.Sprint(i + 1)
		}
	}

	for _, char := range populatedDigits {
		if char != "" {
			calStr += char
		}
	}

	return calStr

}

func findWordIndexes(w string, dw string) []int {
	var indexes []int
	dwLength := len(dw)
	wLength := len(w)
	lastViableIndex := wLength - dwLength
	startPoint := 0

	for {
		subString := w[startPoint:]
		index := strings.Index(subString, dw)
		if index == -1 {
			break
		}

		offsetIndex := index + startPoint
		indexes = append(indexes, offsetIndex)

		startPoint = offsetIndex + dwLength

		if startPoint >= lastViableIndex {
			break
		}
	}
	return indexes
}

func calculateCalibrationValue(calStr string) int64 {
	calVal := int64(0)
	calStrLength := len(calStr)
	if calStrLength == 0 {
		return calVal
	} else if calStrLength == 1 {
		d, err := strconv.ParseInt(calStr, 10, 64)
		if err == nil {
			calVal = d * 11
		}
	} else {
		fD := calStr[0]
		lD := calStr[calStrLength-1]
		n, err := strconv.ParseInt(string(fD)+string(lD), 10, 64)
		if err == nil {
			calVal = n
		}
	}

	return calVal
}
