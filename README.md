# Go-Chess
A fully functioning [chess](https://en.wikipedia.org/wiki/Chess) game developed using GoLang which you can play as a CLI application

Currently I'm learning GoLang and this project helped me a lot to understand concepts such as OOP, Pointers, Generics, etc... in GoLang and hope the same will be for you

## How it works
When you input a move in to the application, first it will check whether it's a legal move or a not. Then it will check whether opponent is check mate or stalemate by playing all possible moves in a virtual board

## Input format
You can play moves in the game by inserting data according to the following format

`Character:From->To</i>` <br />
i.e. `N:a3->b5` (Move knight in a3 square to b5)

## What's implemented upto now

 - [x] Format Input
 - [ ] Input validation and error handling
 - [x] Check for legal moves
 - [x] Check mate
 - [x] Stale mate
 - [ ] Repetition draw
 - [ ] Dead draw
 - [ ] Setup a log file
