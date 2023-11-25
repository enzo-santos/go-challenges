package main

import (
    "testing"
)

func test(t *testing.T, referenceText, expectedVerse string) {
    key, err := ParseBibleKey(referenceText)
    if err != nil {
        t.Fatalf(`ParseBibleKey(%q) threw %q`, referenceText, err)
    }
    actualVerse, err := FindVerse(key)
    if err != nil {
        t.Fatalf(`FindVerse(%q) threw %q`, referenceText, err)
    }
    if actualVerse != expectedVerse {
        t.Fatalf(`FindVerse(%q) = %q, expected %q`, referenceText, actualVerse, expectedVerse)
    }
}

func TestCase01(t *testing.T) {
    test(t, "Genesis 1:1", "In the beginning God created the heaven and the earth.")
}

func TestCase02(t *testing.T) {
    test(t, "1 Samuel 1:1", "Now there was a certain man of Ramathaimzophim, of mount Ephraim, and his name was Elkanah, the son of Jeroham, the son of Elihu, the son of Tohu, the son of Zuph, an Ephrathite:")
}

func TestCase03(t *testing.T) {
    test(t, "1 Kings 1:1", "Now king David was old and stricken in years; and they covered him with clothes, but he gat no heat.")
}

func TestCase04(t *testing.T) {
    test(t, "Psalm 119:11", "Thy word have I hid in mine heart, that I might not sin against thee.")
}

func TestCase05(t *testing.T) {
    test(t, "John 3:16", "For God so loved the world, that he gave his only begotten Son, that whosoever believeth in him should not perish, but have everlasting life.")
}

func TestCase06(t *testing.T) {
    test(t, "1 John 1:9", "If we confess our sins, he is faithful and just to forgive us our sins, and to cleanse us from all unrighteousness.")
}

func TestCase07(t *testing.T) {
    test(t, "3 John 1", "The elder unto the wellbeloved Gaius, whom I love in the truth.")
}

func TestCase8(t *testing.T) {
    test(t, "Jude 21", "Keep yourselves in the love of God, looking for the mercy of our Lord Jesus Christ unto eternal life.")
}

func TestCase09(t *testing.T) {
    test(t, "Revelation 21:11", "Having the glory of God: and her light was like unto a stone most precious, even like a jasper stone, clear as crystal;")
}

func TestCase10(t *testing.T) {
    test(t, "Revelation 21:16", "And the city lieth foursquare, and the length is as large as the breadth: and he measured the city with the reed, twelve thousand furlongs. The length and the breadth and the height of it are equal.")
}

func TestCase11(t *testing.T) {
    test(t, "Revelation 22:20", "He which testifieth these things saith, Surely I come quickly. Amen. Even so, come, Lord Jesus.")
}

func TestCase12(t *testing.T) {
    test(t, "Revelation 22:21", "The grace of our Lord Jesus Christ be with you all. Amen.")
}
