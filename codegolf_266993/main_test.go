package main

import (
    "testing"
)

func Test(t *testing.T) {
    from0to19 := []int{
        1, 1, 1, -1, 1, 1, -1, 1, 1, 1, 1, -1, -1, -1, 1, -1, 1, 1, 1, -1,
    }
    for n, expected := range from0to19 {
        actual := A020985(n)
        if actual != expected {
            t.Fatalf(`A020985(%q) = %q, expected %q`, n, actual, expected)
        }
    }
}
