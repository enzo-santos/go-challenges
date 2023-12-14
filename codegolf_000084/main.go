package main

import (
    "bufio"
    "fmt"
    "io"
    "strings"
)

type Interpreter struct {
    Data      []int
    dataIndex int
    buffer    []byte

    loopSkip    *LoopSkip
    activeLoops []*Loop
}

func (interpreter *Interpreter) Set(value int) {
    index := interpreter.dataIndex
    interpreter.Data[index] = value
}

func (interpreter *Interpreter) Get() int {
    index := interpreter.dataIndex
    return interpreter.Data[index]
}

func (interpreter *Interpreter) CurrentLoop() *Loop {
    activeLoops := interpreter.activeLoops
    if activeLoops == nil {
        return nil
    }
    if len(activeLoops) == 0 {
        return nil
    }
    return activeLoops[len(activeLoops)-1]
}

func (interpreter *Interpreter) PopLoop() {
    activeLoops := interpreter.activeLoops

    loop := activeLoops[len(activeLoops)-1]
    activeLoops = activeLoops[:len(activeLoops)-1]
    if len(activeLoops) > 0 {
        var index int
        if loop.index+1 < len(interpreter.buffer) {
            index = loop.index
        } else {
            index = -1
        }
        activeLoops[len(activeLoops)-1].index = index
    }

    interpreter.activeLoops = activeLoops
}

func (interpreter *Interpreter) PushLoop() {
    var index, startIndex int
    if len(interpreter.activeLoops) > 0 {
        lastActiveLoop := interpreter.activeLoops[len(interpreter.activeLoops)-1]
        if lastActiveLoop.index >= 0 {
            index = lastActiveLoop.index
            startIndex = index
        } else {
            index = -1
            startIndex = len(interpreter.buffer)
        }

    } else {
        index = -1
        startIndex = len(interpreter.buffer)
    }
    for _, loop := range interpreter.activeLoops {
        loop.index = loop.startIndex
    }
    interpreter.activeLoops = append(interpreter.activeLoops, &Loop{
        index:      index,
        startIndex: startIndex,
    })
}

type LoopSkip struct {
    count int
}

type Loop struct {
    startIndex int
    index      int
}

type InterpreterSession struct {
    Reader *bufio.Reader
    Stdin  *bufio.Reader
    Stdout io.Writer
    Logger io.Writer
}

func (interpreter *Interpreter) Process(session InterpreterSession) {
    log := session.Logger
    for {
        var i int
        var b byte
        var err error

        if loop := interpreter.CurrentLoop(); loop != nil && loop.index >= 0 {
            b = interpreter.buffer[loop.index]
            i = loop.index
            loop.index++

        } else {
            b, err = session.Reader.ReadByte()
            if err != nil {
                return
            }
            i = len(interpreter.buffer)
            interpreter.buffer = append(interpreter.buffer, b)
        }
        if skip := interpreter.loopSkip; skip != nil && b != '[' && b != ']' {
            continue
        }

        fmt.Fprintf(log, "%sReading %c at %d, ", strings.Repeat("  ", len(interpreter.activeLoops)), b, i)

        switch b {
        case '>':
            interpreter.dataIndex++
            fmt.Fprintf(log, "   next slot, now  i = %d\n", interpreter.dataIndex)
        case '<':
            interpreter.dataIndex--
            fmt.Fprintf(log, "   prev slot, now  i = %d\n", interpreter.dataIndex)
        case '+':
            interpreter.Set(interpreter.Get() + 1)
            fmt.Fprintf(log, "incrementing, now data[%d] = %d (%v)\n", interpreter.dataIndex, interpreter.Get(), interpreter.Data[:10])
        case '-':
            interpreter.Set(interpreter.Get() - 1)
            fmt.Fprintf(log, "decrementing, now data[%d] = %d (%v)\n", interpreter.dataIndex, interpreter.Get(), interpreter.Data[:10])
        case '.':
            fmt.Fprintf(session.Stdout, "%c", interpreter.Get())
            fmt.Fprintf(log, "%c\n", interpreter.Get())

        case ',':
            bi, err := session.Stdin.ReadByte()
            if err != nil {
                return
            }
            interpreter.Set(int(bi))
            fmt.Fprintf(log, "   consuming, now data[%d] = %d (%v)\n", interpreter.dataIndex, interpreter.Get(), interpreter.Data[:10])
        case '[':
            if skip := interpreter.loopSkip; skip == nil {
                if interpreter.Get() == 0 {
                    interpreter.loopSkip = &LoopSkip{}
                    fmt.Fprintln(log, "LOOP: skip")
                } else {
                    fmt.Fprintln(log, "LOOP: enter")
                    interpreter.PushLoop()
                }
            } else {
                skip.count++
                fmt.Fprintf(log, "FWD, %d\n", interpreter.loopSkip.count)
            }
        case ']':
            if skip := interpreter.loopSkip; skip == nil {
                if interpreter.Get() == 0 {
                    interpreter.PopLoop()
                    fmt.Fprintln(log, "LOOP: break")
                } else if loop := interpreter.CurrentLoop(); loop != nil {
                    fmt.Fprintf(log, "LOOP: continue; data[%d] = %d (non-zero)\n", interpreter.dataIndex, interpreter.Get())
                    loop.index = loop.startIndex
                }
            } else {
                skip.count--
                fmt.Fprintf(log, "FWD, %d", interpreter.loopSkip.count)
                if skip.count == -1 {
                    fmt.Fprintln(log, " [break]")
                    interpreter.loopSkip = nil
                } else {
                    fmt.Fprintln(log)
                }
            }
        }

    }
}

func NewInterpreter() *Interpreter {
    return &Interpreter{
        Data:      make([]int, 200),
        dataIndex: 100,
        buffer:    make([]byte, 0),
    }
}
