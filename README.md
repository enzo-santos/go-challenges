# Go challenges

Implementation of challenges from Stack Exchange websites.

## Installing

You need to [install Go](https://go.dev/doc/install) to run these snippets locally. 
Check the *go.mod* file for the minimum version.

```shell
$ go version
go version go1.21.3 windows/amd64
# or whatever version you have
```

Clone this repository:

```shell
$ git clone https://github.com/enzo-santos/go-challenges
$ cd go-challenges
```

## Usage

Run the following to check all challenges are implemented according to its specs:

```shell
$ go test ./...
```

## Project structure

All challenges in this repository are implemented in Go.

Every directory on this repository is named after a Stack Exchange's host website name and a post ID from 
that website. Therefore, a folder named `<domain>_<postid>` refers to the URL `https://<domain>.stackexchange.com/q/<postid>`.

A `codegolf_*` folder does not necessarily contain [golfed code](https://en.wikipedia.org/wiki/Code_golf): sometimes the 
challenge was considered interesting enough to be implemented in a normal way.

Every directory on this repository contains at least 

- a *README.md* file, containing the challenge description and metadata;
- a *main.go* file, containing the challenge implementation; and
- a *main_test.go* file, containing the challenge test cases as described on the Stack Exchange post

If a challenge is too complex to be implemented in a single *main.go* file, it may be split into multiple files.
