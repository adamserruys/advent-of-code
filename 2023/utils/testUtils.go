package utils

import "fmt"

func PrintResults(testTitle string, want any, got any) {
	fmt.Println("=============================================")
	fmt.Println("Test Title: " + testTitle)
	fmt.Print("WANT: ")
	fmt.Println(want)
	fmt.Print("GOT: ")
	fmt.Println(got)
	fmt.Println("=============================================")
}

func AreIntArraysEqual(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aItem := range a {
		if aItem != b[i] {
			return false
		}
	}
	return true
}
