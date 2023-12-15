package main

import (
    "slices"
    "testing"
)

func test(t *testing.T, word string, expectedPath []string) {
    dictionary, err := NewDictionary()
    if err != nil {
        panic(err)
    }

    context := Context{
        Dictionary: dictionary,
        Cache:      make(map[string][][]string, 0),
    }

    paths := FindPaths(&context, word)
    ok := slices.ContainsFunc(paths, func(actualPath []string) bool {
        return slices.Equal(actualPath, expectedPath)
    })
    if !ok {
        t.Fatalf(`FindPaths(%s) does not contain %q`, word, expectedPath)
    }
}

func TestEmpathise(t *testing.T) {
    test(t, "empathise", []string{"empathise", "shipmate", "atheism", "theism", "times", "time", "tie", "it", "i"})
}
