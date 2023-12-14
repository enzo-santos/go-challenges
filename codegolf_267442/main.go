package main

import (
    "strconv"
)

type Context struct {
    Stack []int
}

type Stackable interface {
    Push(number int)
    Pop() int
}

func (context *Context) Push(number int) {
    context.Stack = append(context.Stack, number)
}

func (context *Context) Pop() int {
    value := context.Stack[len(context.Stack)-1]
    context.Stack = context.Stack[:len(context.Stack)-1]
    return value
}

func (context *Context) Parse(instructions []string) {
    for _, instruction := range instructions {
        switch instruction {
        case "+":
            arg0 := context.Pop()
            arg1 := context.Pop()
            context.Push(arg1 + arg0)
        case "-":
            arg0 := context.Pop()
            arg1 := context.Pop()
            context.Push(arg1 - arg0)
        case "dup":
            arg := context.Pop()
            context.Push(arg)
            context.Push(arg)
        case "drop":
            context.Pop()
        case "swap":
            arg0 := context.Pop()
            arg1 := context.Pop()
            context.Push(arg0)
            context.Push(arg1)
        }
        value, err := strconv.Atoi(instruction)
        if err != nil {
            continue
        }
        context.Push(value)
    }
}
