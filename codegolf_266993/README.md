# codegolf_266993

- Title: ["Rudin-Shapiro sequence"](https://codegolf.stackexchange.com/questions/266993)
- Posted at [Code Golf Stack Exchange](https://codegolf.stackexchange.com)
- Posted by user [@alephalpha](https://codegolf.stackexchange.com/users/9288)
- Posted on November 23, 2023

## Description

The [Rudin-Shapiro sequence](https://en.wikipedia.org/wiki/Rudin%E2%80%93Shapiro_sequence)
is a sequence of *1*s and *-1*s defined as follows: `r_{n} = (-1)^{u_n}`, where `u_n` is
the number of occurrences of (possibly overlapping) `11` in the binary representation
of `n`.

For example, `r_{461} = -1`, because 461 in binary is `111001101`, which
contains 3 occurrences of `11`: `{11}1001101`, `1{11}001101`, `11100{11}01`.

This is sequence [A020985](https://oeis.org/A020985) in the OEIS.

The first few terms of the sequence are:

```
1, 1, 1, -1, 1, 1, -1, 1, 1, 1, 1, -1, -1, -1, 1, -1, 1, 1, 1, -1, 1, 1, -1, 1, -1, -1, -1, 1, 1, 1, -1, 1, 1, 1, 1, -1, 1, 1, -1, 1, 1, 1, 1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, 1, -1, 1, 1, 1, -1, -1, -1, 1, -1, 1, 1, 1, -1, 1, 1, -1, 1, 1, 1, 1, -1, -1, -1, 1, -1, 1
```

\[...\]

Generate the Rudin-Shapiro sequence.
