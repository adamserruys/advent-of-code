package main

import (
	"testing"

	"github.com/adamserruys/advent-of-code/utils"
)

func TestSumUpParts(t *testing.T) {
	s := []int{467, 35, 633, 617, 592, 755, 664, 598}
	want := 4361
	got := sumUpParts(s)

	utils.PrintResults("TestSumUpParts", want, got)

	if want != got {
		t.Fatalf(`sumUpParts(%q) = %d, want match for %d`, s, got, want)
	}
}

func TestFindEngineParts_HappyPath(t *testing.T) {
	s := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
		".........9",
		"........*."}
	want := []int{467, 35, 633, 617, 592, 755, 664, 598, 9}
	got := findEngineParts(s)

	utils.PrintResults("TestFindEngineParts_HappyPath", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findEngineParts(%q) = %q, want match for %#q`, s, got, want)
	}
}

func TestIsPartNumber_HappyPath(t *testing.T) {
	s := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}
	row := 0
	start := 0
	end := 2
	want := true
	got := isPartNumber(row, start, end, s)

	utils.PrintResults("TestIsPartNumber_firstRow", want, got)

	if want != got {
		t.Fatalf(`findEngineParts(%q) = %t, want match for %t`, s, got, want)
	}
}

func TestIsPartNumber_HappyPathTwo(t *testing.T) {
	s := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		".......755",
		"...$..*...",
		".664.598..",
		".........9",
		"........*."}
	row := 10
	start := 9
	end := 9
	want := true
	got := isPartNumber(row, start, end, s)

	utils.PrintResults("TestIsPartNumber_HappyPathTwo", want, got)

	if want != got {
		t.Fatalf(`findEngineParts(%q) = %t, want match for %t`, s, got, want)
	}
}

func TestContainsSpecialChar_True(t *testing.T) {
	s := "617*......"
	want := true
	got := containsSpecialChar(s)

	utils.PrintResults("TestContainsSpecialChar_True", want, got)

	if want != got {
		t.Fatalf(`containsSpecialChar(%q) = %t, want match for %t`, s, got, want)
	}
}

func TestContainsSpecialChar_FalseNoDigits(t *testing.T) {
	s := "......"
	want := false
	got := containsSpecialChar(s)

	utils.PrintResults("TestContainsSpecialChar_FalseNoDigits", want, got)

	if want != got {
		t.Fatalf(`containsSpecialChar(%q) = %t, want match for %t`, s, got, want)
	}
}

func TestContainsSpecialChar_FalseOnlyDigits(t *testing.T) {
	s := "1234556"
	want := false
	got := containsSpecialChar(s)

	utils.PrintResults("TestContainsSpecialChar_FalseOnlyDigits", want, got)

	if want != got {
		t.Fatalf(`containsSpecialChar(%q) = %t, want match for %t`, s, got, want)
	}
}

func TestContainsSpecialChar_FalseMixed(t *testing.T) {
	s := "...123...."
	want := false
	got := containsSpecialChar(s)

	utils.PrintResults("TestContainsSpecialChar_FalseMixed", want, got)

	if want != got {
		t.Fatalf(`containsSpecialChar(%q) = %t, want match for %t`, s, got, want)
	}
}

func TestFindGearRatios_HappyPath(t *testing.T) {
	s := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
		".........9",
		"........*."}
	want := []int{16345, 451490}
	got := findGearRatios(s)

	utils.PrintResults("TestFindGearRatios_HappyPath", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findGearRatios(%q) = %q, want match for %#q`, s, got, want)
	}
}

func TestFindSurroundingParts_FoundOne(t *testing.T) {
	s := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
		".........9",
		"........*."}
	row := 4
	gearIndex := 3
	want := []int{617}
	got := findSurroundingParts(row, gearIndex, s)

	utils.PrintResults("TestFindSurroundingParts_FoundOne", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findSurroundingParts(%q,%q,%q ) = %q, want match for %#q`, row, gearIndex, s, got, want)
	}
}

func TestFindSurroundingParts_FoundTwo(t *testing.T) {
	s := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
		".........9",
		"........*."}
	row := 1
	gearIndex := 3
	want := []int{467, 35}
	got := findSurroundingParts(row, gearIndex, s)

	utils.PrintResults("TestFindSurroundingParts_FoundTwo", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findSurroundingParts(%q,%q,%q ) = %q, want match for %#q`, row, gearIndex, s, got, want)
	}
}

func TestFindSurroundingParts_FoundThree(t *testing.T) {
	s := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"...45.755.",
		"...$.*....",
		".664.598..",
		".........9",
		"........*."}
	row := 8
	gearIndex := 5
	want := []int{45, 755, 598}
	got := findSurroundingParts(row, gearIndex, s)

	utils.PrintResults("TestFindSurroundingParts_FoundThree", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findSurroundingParts(%q,%q,%q ) = %q, want match for %#q`, row, gearIndex, s, got, want)
	}
}

func TestFindAdjPartNums_DiffRowOne(t *testing.T) {
	row := "...45.755."
	gearIndex := 6
	want := []int{755}
	got := findAdjPartNums(row, gearIndex)

	utils.PrintResults("TestFindAdjPartNums_DiffRowOne", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findSurroundingParts(%q,%q) = %q, want match for %#q`, row, gearIndex, got, want)
	}
}

func TestFindAdjPartNums_DiffRowTwo(t *testing.T) {
	row := "...45.755."
	gearIndex := 5
	want := []int{45, 755}
	got := findAdjPartNums(row, gearIndex)

	utils.PrintResults("TestFindAdjPartNums_DiffRowOne", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findSurroundingParts(%q,%q) = %q, want match for %#q`, row, gearIndex, got, want)
	}
}

func TestFindAdjPartNums_DiffNone(t *testing.T) {
	row := "...45.755."
	gearIndex := 1
	want := []int{}
	got := findAdjPartNums(row, gearIndex)

	utils.PrintResults("TestFindAdjPartNums_DiffRowOne", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findSurroundingParts(%q,%q) = %q, want match for %#q`, row, gearIndex, got, want)
	}
}
