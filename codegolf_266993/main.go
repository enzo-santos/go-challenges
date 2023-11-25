package main

import (
    "fmt"
    "math"
)

func countConsecutive1Bits(n int) (count int) {
    var initialized bool
    var previousByte int
    for i := 0; i < int(math.Log2(float64(n)))+1; i++ {
        currentByte := (n >> i) & 1
        if initialized && currentByte == 1 && currentByte == previousByte {
            count++
        }
        initialized = true
        previousByte = currentByte
    }
    return
}

func A020985(n int) int {
    if countConsecutive1Bits(n)%2 == 0 {
        return 1
    } else {
        return -1
    }
}

func main() {
    for i := 0; i < 20; i++ {
        fmt.Printf("%d -> %d\n", i, A020985(i))
    }
}
