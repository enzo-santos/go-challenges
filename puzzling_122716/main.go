package main

func intPow(n int, m int) int {
    if m == 0 {
        return 1
    }
    result := n
    for i := 2; i <= m; i++ {
        result *= n
    }
    return result
}

func encodeTextToCharIndices(text string) map[string][]int {
    result := make(map[string][]int, 0)
    for i, rune := range text {
        char := string(rune)
        indices, found := result[char]
        if !found {
            indices = make([]int, 0)
        }
        result[char] = append(indices, len(text)-i-1)
    }
    return result
}

func encodeIntToCharDigits(value int, charIndices map[string][]int) (map[string]int, bool) {
    result := make(map[string]int, len(charIndices))
    for char, indices := range charIndices {
        for _, index := range indices {
            actualDigit := value / intPow(10, index) % 10
            expectedDigit, found := result[char]
            if found && actualDigit != expectedDigit {
                return nil, false
            }
            result[char] = actualDigit
        }
    }
    return result, true
}

func compareCharDigits(c0, c1 map[string]int) bool {
    for char, actualDigit := range c0 {
        expectedDigit, found := c1[char]
        if found && actualDigit != expectedDigit {
            return false
        }
    }
    for char, actualDigit := range c1 {
        expectedDigit, found := c0[char]
        if found && actualDigit != expectedDigit {
            return false
        }
    }
    return true
}

// SquareAlphametic solves an alphametic of form sqrt([leftText]) = [rightText].
func SquareAlphametic(leftText, rightText string) []int {
    leftCharIndices := encodeTextToCharIndices(leftText)
    rightCharIndices := encodeTextToCharIndices(rightText)

    minValue := intPow(10, len(rightText)-1)
    maxValue := intPow(10, len(rightText)) - 1
    minSqrValue := intPow(10, len(leftText)-1)
    maxSqrValue := intPow(10, len(leftText)) - 1

    result := make([]int, 0)
    for value := minValue; value <= maxValue; value++ {
        rightCharDigits, ok := encodeIntToCharDigits(value, rightCharIndices)
        if !ok {
            continue
        }
        sqrValue := value * value
        if sqrValue < minSqrValue || sqrValue > maxSqrValue {
            continue
        }

        leftCharDigits, ok := encodeIntToCharDigits(sqrValue, leftCharIndices)
        if !ok {
            continue
        }

        if !compareCharDigits(leftCharDigits, rightCharDigits) {
            continue
        }
        result = append(result, value)
    }
    return result
}
