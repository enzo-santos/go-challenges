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
    "strconv"
    "strings"
)

func downloadBook(url string) (string, error) {
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

type BibleKey struct {
    BookTitle     string
    ChapterNumber int
    VerseNumber   int
}

var ordinalRegexp *regexp.Regexp = regexp.MustCompile(`^The ([A-Za-z]+) .+ of(?:[^:]+)? (\w+)$`)
var newVerseRegexp *regexp.Regexp = regexp.MustCompile(`(.*?)([0-9]+:[0-9]+) (.+)`)

var ordinalValues []string = []string{"First", "Second", "Third"}

func BookTitleOf(line string) (title string, ok bool) {
    groups := ordinalRegexp.FindStringSubmatch(line)
    if groups != nil {
        ordinalValue := groups[1]
        bookTitle := groups[2]

        index := slices.Index(ordinalValues, ordinalValue)
        if index >= 0 {
            return fmt.Sprintf("%d %s", index+1, bookTitle), true
        }
    }

    var bookTitle string
    words := strings.Split(line, " ")
    switch lastWord := words[len(words)-1]; lastWord {
    case "Bible":
        {
            return
        }
    case "Solomon":
        {
            bookTitle = "Song of Solomon"
        }
    case "Apostles":
        {
            bookTitle = "Acts"
        }
    case "Psalms":
        {
            bookTitle = "Psalm"
        }
    case "Divine":
        {
            bookTitle = "Revelation"
        }
    default:
        {
            if strings.Contains(line, "Lamentations") {
                bookTitle = "Lamentations"
            } else {
                bookTitle = lastWord
            }
        }
    }
    return bookTitle, true
}

func ParseBibleKey(label string) (BibleKey, error) {
    tokens := strings.Split(label, " ")
    var bookTitle string
    if len(tokens) < 3 {
        bookTitle = tokens[0]
    } else {
        bookTitle = fmt.Sprintf("%s %s", tokens[0], tokens[1])
    }

    chapterNumber := 1
    subtokens := strings.Split(tokens[len(tokens)-1], ":")
    verseNumber, err := strconv.Atoi(subtokens[len(subtokens)-1])
    if err != nil {
        return BibleKey{}, fmt.Errorf("Failed to convert verse number to integer: %v", err)
    }
    if len(subtokens) > 1 {
        chapterNumber, err = strconv.Atoi(subtokens[0])
        if err != nil {
            return BibleKey{}, fmt.Errorf("Failed to convert chapter number to integer: %v", err)
        }
    }
    return BibleKey{
        BookTitle:     bookTitle,
        ChapterNumber: chapterNumber,
        VerseNumber:   verseNumber,
    }, nil
}

type Context struct {
    Key        BibleKey
    VerseParts []string
    IndexTitle string
    IgnoreNext int
}

type ParsingStep interface {
    Validate(context *Context, line string) bool
}

type bibleIndexParsingStep struct{}

// Validate of `bibleIndexParsingStep` locates where the Bible index starts.
//
// This function will iterate the Bible line by line until it finds "*** START OF THE
// PROJECT GUTENBERG EBOOK THE KING JAMES VERSION OF THE BIBLE ***".
func (bibleIndexParsingStep) Validate(context *Context, line string) bool {
    return strings.HasPrefix(line, "*** START")
}

type indexTitleParsingStep struct{}

// Validate of `indexTitleParsingStep` locates the corresponding Bible-based book title
// for a given user-based book title.
//
// If "Psalm" is given as the book title, this function will locate its corresponding
// Bible index line ("The Book of Psalms") and store it for future use.
func (indexTitleParsingStep) Validate(context *Context, line string) bool {
    if bookTitle, ok := BookTitleOf(line); ok && bookTitle == context.Key.BookTitle {
        context.IndexTitle = line
        return true
    }
    return false
}

type bookParsingStep struct{}

// Validate of `bookParsingStep` locates the Bible-based book title located by
// `indexTitleParsingStep`.
//
// If the stored book title is "The Book of Psalms", this function will iterate the
// Bible line by line until it finds a line with this exact title.
func (bookParsingStep) Validate(context *Context, line string) bool {
    if line == "Otherwise Called:" {
        context.IgnoreNext = 2
        return false
    }
    return line == context.IndexTitle
}

type startOfVerseParsingStep struct{}

// Validate of `startOfVerseParsingStep` locates the given chapter and verse within the
// book located by `bookParsingStep`.
//
// If "Psalm 192:11" is given as chapter and verse number, this function will iterate
// the "The Book of Psalms" by line until it finds this exact reference.
func (startOfVerseParsingStep) Validate(context *Context, line string) bool {
    groups := newVerseRegexp.FindStringSubmatch(line)
    if groups == nil {
        return false
    }

    ref := fmt.Sprintf("%d:%d", context.Key.ChapterNumber, context.Key.VerseNumber)
    if ref == groups[2] {
        context.VerseParts = append(context.VerseParts, groups[3])
        return true
    }
    return false
}

type endOfVerseParsingStep struct{}

// Validate of `startOfVerseParsingStep` locates the end of the paragraph which the line
// located by `startOfVerseParsingStep` corresponds to.
//
// If "Psalm 192:11" is given as chapter and verse number, this function will iterate
// every line after it until it finds another reference (such as 192:12) or the end of
// the Bible.
func (endOfVerseParsingStep) Validate(context *Context, line string) bool {
    if len(line) == 0 {
        return false
    }

    // The end of the Bible
    if line == "        " {
        return true
    }

    groups := newVerseRegexp.FindStringSubmatch(line)
    if groups == nil {
        context.VerseParts = append(context.VerseParts, line)
        return false
    }

    // Another reference is found
    if len(groups[1]) > 1 {
        context.VerseParts = append(context.VerseParts, strings.TrimSpace(groups[1]))
    }
    return true
}

func FindVerse(key BibleKey) (string, error) {
    fpath, err := downloadBook("https://gutenberg.org/cache/epub/10/pg10.txt")
    if err != nil {
        return "", fmt.Errorf("Failed to download the Bible: %v", err)
    }
    f, err := os.Open(fpath)
    if err != nil {
        return "", fmt.Errorf("Failed to open the downloaded Bible: %v", err)
    }
    defer f.Close()

    context := &Context{Key: key, VerseParts: make([]string, 0)}
    parsingSteps := []ParsingStep{
        bibleIndexParsingStep{},
        indexTitleParsingStep{},
        bookParsingStep{},
        startOfVerseParsingStep{},
        endOfVerseParsingStep{},
    }
    parsingStepIndex := 0

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()

        if context.IgnoreNext > 0 {
            context.IgnoreNext--
            continue
        }

        parsingStep := parsingSteps[parsingStepIndex]
        valid := parsingStep.Validate(context, line)
        if !valid {
            continue
        }
        if parsingStepIndex == len(parsingSteps)-1 {
            break
        }
        parsingStepIndex++
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
    return strings.Join(context.VerseParts, " "), nil
}
