# rndnumguess
> A simple "Random number guessing game" implementation of the [stdio](https://github.com/pbreedt/stdio) module that reads input from standard IO (os.Stdin)

## General info
This project is simply playing around with:
- Go modules (with nothing in the main working folder)
- Reading user input from command line

## Technologies
* Go - version 1.15

## Setup
You need [Go](https://golang.org/) to run/build this project.
If you have that in place, get the source code...

```bash
git clone https://github.com/pbreedt/rndnumguess.git
```

...then, either run with Go:

```bash
go run main.go
```

...or build with Go and run:

```bash
go build
./rndnumguess
Guess my number:50
Your guess was HIGH, try again
Guess my number:25
Your guess was LOW, try again
Guess my number:35
Your guess was LOW, try again
Guess my number:42
Your guess was LOW, try again
Guess my number:46
Your guess was LOW, try again
Guess my number:48
CORRECT!
```

## Features
Current features include:
* A random number is generated
* User gets 10 attempts to guess it

## Status
Project is: _in progress_

## Credits
Thanks to [ritaly](https://github.com/ritaly/README-cheatsheet) for a quick readme template

## Contact
Created by [pbreedt](mailto:petrus.breedt@gmail.com)