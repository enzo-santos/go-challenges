# codegolf_266951

- Title: ["Get the Bible verse"](https://codegolf.stackexchange.com/questions/266951)
- Posted at [Code Golf Stack Exchange](https://codegolf.stackexchange.com)
- Posted by user [@Somebody](https://codegolf.stackexchange.com/users/112554)
- Posted on November 21, 2023

## Description

Write a program that accesses [Project Gutenberg's KJV Bible][1], either from a file
called `pg10.txt` in the working directory or in any other way that is typical for the
language, takes a Bible verse reference as input, and prints the corresponding verse,
with no newlines except optionally one at the end, as output.

The standard format for a reference is `[book] [chapter]:[verse]`, e.g. `John 3:13`.
However, some books have only one chapter. For these books, the format is `[book]
[verse]`, e.g. `Jude 2`. Your program must accept this format for single-chapter books.
Behavior when using the standard format with single-chapter books, or when using the
single-chapter format with multi-chapter books, is undefined.

Note that the "title" for the book of Psalms, for the purposes of this question, is
`Psalm`, not `Psalms`, because references are usually written like `Psalm 119:11`, not
`Psalms 119:11`.

\[...\]

| Reference | Text |
|-----------|------|
| Genesis 1:1 | In the beginning God created the heaven and the earth. |
| 1 Samuel 1:1 | Now there was a certain man of Ramathaimzophim, of mount Ephraim, and his name was Elkanah, the son of Jeroham, the son of Elihu, the son of Tohu, the son of Zuph, an Ephrathite: |
| 1 Kings 1:1 | Now king David was old and stricken in years; and they covered him with clothes, but he gat no heat. |
| Psalm 119:11 | Thy word have I hid in mine heart, that I might not sin against thee. |
| John 3:16 | For God so loved the world, that he gave his only begotten Son, that whosoever believeth in him should not perish, but have everlasting life. | Then said the Jews unto him, Now we know that thou hast a devil. Abraham is dead, and the prophets; and thou sayest, If a man keep my saying, he shall never taste of death. |
| 1 John 1:9 | If we confess our sins, he is faithful and just to forgive us our sins, and to cleanse us from all unrighteousness. |
| 3 John 1 | The elder unto the wellbeloved Gaius, whom I love in the truth. |
| Jude 21 | Keep yourselves in the love of God, looking for the mercy of our Lord Jesus Christ unto eternal life. |
| Revelation 21:11 | Having the glory of God: and her light was like unto a stone most precious, even like a jasper stone, clear as crystal; |
| Revelation 21:16 | And the city lieth foursquare, and the length is as large as the breadth: and he measured the city with the reed, twelve thousand furlongs. The length and the breadth and the height of it are equal. |
| Revelation 22:20 | He which testifieth these things saith, Surely I come quickly. Amen. Even so, come, Lord Jesus. |
| Revelation 22:21 | The grace of our Lord Jesus Christ be with you all. Amen. |

  [1]: https://gutenberg.org/cache/epub/10/pg10.txt
