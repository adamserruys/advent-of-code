package main

import (
	"github.com/adamserruys/advent-of-code/utils"
	"testing"
)

func TestParseGame_HappyPath(t *testing.T) {
	s := "Game 3: 11 blue, 3 red, 1 green; 15 red, 9 blue, 3 green; 11 blue, 4 red, 4 green; 1 red, 2 green, 14 blue; 18 blue, 4 green, 10 red"
	want := []int{3, 18, 4, 15}
	got := parseGame(s)

	utils.PrintResults("TestParseGame_HappyPath", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`parseGame("%q") = %q, want match for %#q`, s, got, want)
	}
}

func TestParseGame_Empty(t *testing.T) {
	s := ""
	want := []int{0, 0, 0, 0}
	got := parseGame(s)

	utils.PrintResults("TestParseGame_Empty", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`parseGame("%q") = %q, want match for %#q`, s, got, want)
	}
}

func TestGetElegibleId_NotEligible(t *testing.T) {
	s := "Game 3: 11 blue, 3 red, 1 green; 15 red, 9 blue, 3 green; 11 blue, 4 red, 4 green; 1 red, 2 green, 14 blue; 18 blue, 4 green, 10 red"
	maxBlue := 14
	maxGreen := 13
	maxRed := 12
	want := 0
	got := getEligibleId(s, maxBlue, maxGreen, maxRed)

	utils.PrintResults("TestGetEligibleId_NotEligible", want, got)

	if want != got {
		t.Fatalf(`getEligibleId("%q",%q,%q,%q) = %q, want match for %#q`, s, maxBlue, maxGreen, maxRed, got, want)
	}
}

func TestGetElegibleId_Eligible(t *testing.T) {
	s := "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	maxBlue := 14
	maxGreen := 13
	maxRed := 12
	want := 2
	got := getEligibleId(s, maxBlue, maxGreen, maxRed)

	utils.PrintResults("TestGetEligibleId_Eligible", want, got)

	if want != got {
		t.Fatalf(`getEligibleId("%q",%q,%q,%q) = %q, want match for %#q`, s, maxBlue, maxGreen, maxRed, got, want)
	}
}

func TestGetCubePower_HappyPath(t *testing.T) {
	s := "Game 3: 11 blue, 3 red, 1 green; 15 red, 9 blue, 3 green; 11 blue, 4 red, 4 green; 1 red, 2 green, 14 blue; 18 blue, 4 green, 10 red"
	want := 1080
	got := getCubePower(s)

	utils.PrintResults("TestGetCubePower_HappyPath", want, got)

	if want != got {
		t.Fatalf(`getCubePower("%q") = %q, want match for %#q`, s, got, want)
	}
}

func TestGetCubePower_Empty(t *testing.T) {
	s := ""
	want := 0
	got := getCubePower(s)

	utils.PrintResults("TestGetCubePower_Empty", want, got)

	if want != got {
		t.Fatalf(`getCubePower("%q") = %q, want match for %#q`, s, got, want)
	}
}
