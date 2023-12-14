package main

import (
    "testing"

    "bufio"
    "io"
    "strings"
)

func test(t *testing.T, code string, input string, expectedOutput string) {
    interpreter := NewInterpreter()

    var stdout strings.Builder
    interpreter.Process(InterpreterSession{
        Reader: bufio.NewReader(strings.NewReader(code)),
        Stdin:  bufio.NewReader(strings.NewReader(input)),
        Stdout: &stdout,
        Logger: io.Discard,
    })
    actualOutput := stdout.String()
    if actualOutput != expectedOutput {
        t.Fatalf(`Brainfuck(%q, %q) = %q, expected %q`, code, input, actualOutput, expectedOutput)
    }
}

func TestAdd(t *testing.T) {
    // https://en.wikipedia.org/wiki/Brainfuck#Adding_two_values
    test(t, "++>+++++[<+>-]++++++++[<++++++>-]<.", "", "7")
}

func TestAddAsciiInputs(t *testing.T) {
    // https://codegolf.stackexchange.com/a/220474/91472
    test(t, ",>,<[->+<]>.", "\x01\x02", "\x03")
    test(t, ",>,<[->+<]>.", "\x13\x54", "\x67")
}

func TestIsAsciiInputPrime(t *testing.T) {
    // https://codegolf.stackexchange.com/a/86586/91472
    test(t, ",-[+[<+>>+<-]>[-<[>+<-]<[>+>->+<[>]>[<+>-]<<[<]>-]>>>]<-[<]]<.", "C", "C")
    test(t, ",-[+[<+>>+<-]>[-<[>+<-]<[>+>->+<[>]>[<+>-]<<[<]>-]>>>]<-[<]]<.", "D", "\x00")
}

func TestStringMultiply(t *testing.T) {
    // https://codegolf.stackexchange.com/a/132343/91472
    test(t, ",>>,[<<[->+>.<<]>[-<+>]>,]", "\x03Hello, world!", "HHHeeellllllooo,,,   wwwooorrrlllddd!!!")
}

func TestHelloWorld1(t *testing.T) {
    // https://en.wikipedia.org/wiki/Brainfuck#Hello_World!
    test(t, "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.", "", "Hello World!")
}

func TestHelloWorld2(t *testing.T) {
    // https://therenegadecoder.com/code/hello-world-in-brainfuck/
    test(t, ">++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+.", "", "Hello, World!")
}

func TestGoodbyeWorld(t *testing.T) {
    // https://rosettacode.org/wiki/Hello_world/Text#Brainf***
    test(t, "++++++++++[>+>+++>++++>+++++++>++++++++>+++++++++>++++++++++>+++++++++++>++++++++++++<<<<<<<<<-]>>>>+.>>>>+..<.<++++++++.>>>+.<<+.<<<<++++.<++.>>>+++++++.>>>.+++.<+++++++.--------.<<<<<+.", "", "Goodbye, World!")
}

// TODO Hanging
func testAddNumericInputs(t *testing.T) {
    // https://codegolf.stackexchange.com/a/91335/91472
    test(t, "+[-->++++++[-<------>]+>>,----------]<,[<+++++[->--------<]+[<<<]>>[-]>[>[-<<<+<[-]+>>>>]>>]<<,]<-<<<[>[->+<]>[-<+>[-<+>[-<+>[-<+>[-<+>[-<+>[-<+>[-<+>[-<+>[-<[-]<<+<<[-]+>>>>>[-<+>]]]]]]]]]]]<<<<<]>>>[+++++[->++++++++<]>.>>]", "31415926535897932384626433832795\n27182818284590452353602874713527\n", "58598744820488384738229308546322")
}
