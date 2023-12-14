# codegolf_000084

- Title: ["Interpret brainfuck"](https://codegolf.stackexchange.com/questions/84)
- Posted at [Code Golf Stack Exchange](https://codegolf.stackexchange.com)
- Posted by user [@Alexandru](https://codegolf.stackexchange.com/users/32)
- Posted on January 28, 2011

## Description

Write the shortest program in your favourite language to interpret a [brainfuck](http://en.wikipedia.org/wiki/brainfuck)
program. The program is read from a file. Input and output are standard input and
standard output.

### Additional description

| **Character** | **Meaning**                                                                                                                                                                         |
|:-------------:|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|      `>`      | Increment the data pointer by one (to point to the next cell to the right).                                                                                                         |
|      `<`      | Decrement the data pointer by one (to point to the next cell to the left).                                                                                                          |
|      `+`      | Increment the byte at the data pointer by one.                                                                                                                                      |
|      `-`      | Decrement the byte at the data pointer by one.                                                                                                                                      |
|      `.`      | Output the byte at the data pointer.                                                                                                                                                |
|      `,`      | Accept one byte of input, storing its value in the byte at the data pointer.                                                                                                        |
|      `[`      | If the byte at the data pointer is zero, then instead of moving the instruction pointer forward to the next command, jump it forward to the command after the matching `]` command. |
|      `]`      | If the byte at the data pointer is nonzero, then instead of moving the instruction pointer forward to the next command, jump it back to the command after the matching `[` command. |
