package main

import (
	"github.com/adamserruys/advent-of-code/utils"
	"testing"
)

func TestGetPartOneCalibrationString_HappyPath(t *testing.T) {
	s := "ab4i89"
	want := "489"
	got := getPartOneCalibrationString(s)

	utils.PrintResults("TestGetPartOneCalibrationString_happy_path", want, got)

	if want != got {
		t.Fatalf(`getPartOneCalibrationString("ab4i89") = %q, want match for %#q`, got, want)
	}
}

func TestGetPartOneCalibrationString_Empty(t *testing.T) {
	s := ""
	want := ""
	got := getPartOneCalibrationString(s)

	utils.PrintResults("TestGetPartOneCalibrationString_empty", want, got)

	if want != got {
		t.Fatalf(`getPartOneCalibrationString("") = %q, want match for %#q`, got, want)
	}
}

func TestFindWordIndexes_OneInstance(t *testing.T) {
	w := "abcone2threexyz"
	dw := "one"
	want := []int{3}
	got := findWordIndexes(w, dw)

	utils.PrintResults("TestFindWordIndexes_OneInstance", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findWordIndexes("abcone2threexyz","one") = %q, want match for %#q`, got, want)
	}
}

func TestFindWordIndexes_TwoInstances(t *testing.T) {
	w := "abcone2threexyzone9z"
	dw := "one"
	want := []int{3, 15}
	got := findWordIndexes(w, dw)

	utils.PrintResults("TestFindWordIndexes_TwoInstances", want, got)

	if !utils.AreIntArraysEqual(want, got) {
		t.Fatalf(`findWordIndexes("abcone2threexyzone9z","one") = %q, want match for %#q`, got, want)
	}
}

func TestGetPartTwoCalibrationString_HappyPath(t *testing.T) {
	s := "abcone2threex6yzeighte"
	want := "12368"
	got := getPartTwoCalibrationString(s)

	utils.PrintResults("TestGetPartTwoCalibrationString_HappyPath", want, got)

	if want != got {
		t.Fatalf(`getPartOneCalibrationString("abcone2threexyz") = %q, want match for %#q`, got, want)
	}
}

func TestGetPartTwoCalibrationString_Empty(t *testing.T) {
	s := ""
	want := ""
	got := getPartTwoCalibrationString(s)

	utils.PrintResults("TestGetPartTwoCalibrationString_Empty", want, got)

	if want != got {
		t.Fatalf(`getPartOneCalibrationString("") = %q, want match for %#q`, got, want)
	}
}

func TestCalculateCalibrationValue_SingleDigit(t *testing.T) {
	s := "4"
	want := int64(44)
	got := calculateCalibrationValue(s)

	utils.PrintResults("TestCalculateCalibrationValue_SingleDigit", want, got)

	if want != got {
		t.Fatalf(`calculateCalibrationValue("4") = %q, want match for %#q`, got, want)
	}
}

func TestCalculateCalibrationValue_MultiDigit(t *testing.T) {
	s := "38432"
	want := int64(32)
	got := calculateCalibrationValue(s)

	utils.PrintResults("TestCalculateCalibrationValue_MultiDigit", want, got)

	if want != got {
		t.Fatalf(`calculateCalibrationValue("38432") = %q, want match for %#q`, got, want)
	}
}

func TestCalculateCalibrationValue_Empty(t *testing.T) {
	s := ""
	want := int64(0)
	got := calculateCalibrationValue(s)

	utils.PrintResults("TestCalculateCalibrationValue_Empty", want, got)

	if want != got {
		t.Fatalf(`calculateCalibrationValue("") = %q, want match for %#q`, got, want)
	}
}

func TestCalculateCalibrationValue_SingleNonDigit(t *testing.T) {
	s := "h"
	want := int64(0)
	got := calculateCalibrationValue(s)

	utils.PrintResults("TestCalculateCalibrationValue_SingleNonDigit", want, got)

	if want != got {
		t.Fatalf(`calculateCalibrationValue("h") = %q, want match for %#q`, got, want)
	}
}

func TestCalculateCalibrationValue_MultiNonDigit(t *testing.T) {
	s := "hello"
	want := int64(0)
	got := calculateCalibrationValue(s)

	utils.PrintResults("TestCalculateCalibrationValue_MultiNonDigit", want, got)

	if want != got {
		t.Fatalf(`calculateCalibrationValue("hello") = %q, want match for %#q`, got, want)
	}
}
