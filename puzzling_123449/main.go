package main

import (
    "bufio"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "slices"
    "sort"
)

func downloadUrl(url string) (string, error) {
    os.MkdirAll("cache", 755)
    fpath := filepath.Join("cache", filepath.Base(url))
    if _, err := os.Stat(fpath); err == nil {
        return fpath, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return "", fmt.Errorf("Failed to create the download request: %v", err)
    }

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("Failed to execute the download request: %v", err)
    }
    defer res.Body.Close()

    f, err := os.Create(fpath)
    if err != nil {
        return "", fmt.Errorf("Failed to open the file to write: %v", err)
    }
    defer f.Close()

    if _, err := io.Copy(f, res.Body); err != nil {
        return "", fmt.Errorf("Failed to save the downloaded data into the file: %v", err)
    }
    return fpath, nil
}

type Dictionary struct {
    Words []string
}

func NewDictionary() (d Dictionary, rerr error) {
    fpath, err := downloadUrl("https://websites.umich.edu/~jlawler/wordlist")
    if err != nil {
        rerr = fmt.Errorf("Failed to download the dictionary: %v", err)
        return
    }
    f, err := os.Open(fpath)
    if err != nil {
        rerr = fmt.Errorf("Failed to open the downloaded dictionary: %v", err)
        return
    }
    defer f.Close()

    words := make([]string, 0)

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        word := scanner.Text()
        if matched, err := regexp.MatchString(`^[a-z]+$`, word); !matched || err != nil {
            continue
        }
        words = append(words, word)
    }
    return Dictionary{words}, scanner.Err()
}

func (dictionary Dictionary) OfLength(length int) Dictionary {
    words := make([]string, 0)
    for _, word := range dictionary.Words {
        if len(word) == length {
            words = append(words, word)
        }
    }
    return Dictionary{words}
}

type Context struct {
    Dictionary Dictionary
    Cache      map[string][][]string
}

func FindPaths(context *Context, word string) [][]string {
    path := []string{word}

    if len(word) == 1 {
        return [][]string{path}
    }

    childDictionary := context.Dictionary.OfLength(len(word) - 1)

    paths := make([][]string, 0)
    for i, _ := range word {
        actualWord := word[:i] + word[i+1:]
        actualSortedRunes := []byte(actualWord)
        sort.Slice(actualSortedRunes, func(i, j int) bool {
            return actualSortedRunes[i] < actualSortedRunes[j]
        })

        for _, expectedWord := range childDictionary.Words {
            expectedSortedRunes := []byte(expectedWord)
            sort.Slice(expectedSortedRunes, func(i, j int) bool {
                return expectedSortedRunes[i] < expectedSortedRunes[j]
            })
            if !slices.Equal(actualSortedRunes, expectedSortedRunes) {
                continue
            }

            childPaths, found := context.Cache[expectedWord]
            if !found {
                childPaths = FindPaths(context, expectedWord)
                context.Cache[expectedWord] = childPaths
            }
            for _, childPath := range childPaths {
                paths = append(paths, append(path, childPath...))
            }
        }
    }

    return paths
}
