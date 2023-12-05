package main

import (
	"github.com/adamserruys/advent-of-code/utils"
	"testing"
)

func TestFindNumOfMatches_Multiple(t *testing.T) {
	c := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	want := 4
	got := findNumOfMatches(c)

	utils.PrintResults("TestFindNumOfMatches_Multiple", want, got)

	if want != got {
		t.Fatalf(`findNumOfMatches(%q) = %d, want match for %d`, c, got, want)
	}
}

func TestFindNumOfMatches_One(t *testing.T) {
	c := "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"
	want := 1
	got := findNumOfMatches(c)

	utils.PrintResults("TestFindNumOfMatches_One", want, got)

	if want != got {
		t.Fatalf(`findNumOfMatches(%q) = %d, want match for %d`, c, got, want)
	}
}

func TestFindNumOfMatches_Zero(t *testing.T) {
	c := "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"
	want := 0
	got := findNumOfMatches(c)

	utils.PrintResults("TestFindNumOfMatches_Zero", want, got)

	if want != got {
		t.Fatalf(`findNumOfMatches(%q) = %d, want match for %d`, c, got, want)
	}
}

func TestFindScoreOfCard_Multiple(t *testing.T) {
	c := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	want := 8
	got := findScoreOfCard(c)

	utils.PrintResults("TestFindScoreOfCard_Multiple", want, got)

	if want != got {
		t.Fatalf(`findScoreOfCard(%q) = %d, want match for %d`, c, got, want)
	}
}

func TestFindScoreOfCard_One(t *testing.T) {
	c := "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"
	want := 1
	got := findScoreOfCard(c)

	utils.PrintResults("TestFindScoreOfCard_Multiple", want, got)

	if want != got {
		t.Fatalf(`findScoreOfCard(%q) = %d, want match for %d`, c, got, want)
	}
}

func TestFindScoreOfCard_Zero(t *testing.T) {
	c := "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"
	want := 0
	got := findScoreOfCard(c)

	utils.PrintResults("TestFindScoreOfCard_Multiple", want, got)

	if want != got {
		t.Fatalf(`findScoreOfCard(%q) = %d, want match for %d`, c, got, want)
	}
}

func TestFindNumOfCardsWon(t *testing.T) {
	c := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	want := []int{1, 2, 4, 8, 14, 1}
	got := findNumberOfCardsWon(c)

	utils.PrintResults("TestFindNumOfCardsWon", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findNumberOfCardsWons(%q) = %q, want match for %q`, c, got, want)
	}

}
