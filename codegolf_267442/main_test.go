package main

import (
    "slices"
    "testing"
)

func test(t *testing.T, instructions []string, expectedStack []int) {
    context := &Context{Stack: make([]int, 0)}
    context.Parse(instructions)
    actualStack := context.Stack
    if !slices.Equal(actualStack, expectedStack) {
        t.Fatalf(`Parse(%q) = %q, expected %q`, instructions, actualStack, expectedStack)
    }
}

func TestSum(t *testing.T) {
    test(t, []string{"1", "2", "+"}, []int{3})
}

func TestDrop(t *testing.T) {
    test(t, []string{"1", "drop"}, []int{})
}

func TestSubtractAndSum(t *testing.T) {
    test(t, []string{"10", "2", "-", "3", "+"}, []int{11})
}

func TestDupAndSubtract(t *testing.T) {
    test(t, []string{"9", "dup", "-"}, []int{0})
}

func TestIdentity(t *testing.T) {
    test(t, []string{"1", "2", "3"}, []int{1, 2, 3})
}

func TestDupAndSum(t *testing.T) {
    test(t, []string{"1", "dup", "2", "+"}, []int{1, 3})
}

func TestSwap(t *testing.T) {
    test(t, []string{"3", "2", "5", "swap"}, []int{3, 5, 2})
}
