package main

import (
    "slices"
    "testing"
)

func test(t *testing.T, leftText, rightText string, expectedValues []int) {
    actualValues := SquareAlphametic(leftText, rightText)
    if !slices.Equal(actualValues, expectedValues) {
        t.Fatalf("SquareAlphametic(%q, %q) = %v, expected %v", leftText, rightText, actualValues, expectedValues)
    }
}

func TestIlluminateLight(t *testing.T) {
    test(t, "illuminate", "light", []int{75978})
}
