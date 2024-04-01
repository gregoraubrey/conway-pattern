# conway-pattern
A program that someone asked me to write for generating a number pattern that John Conway found interesting.

You start with two positive integers, then sum them. If this sum is prime, you add it to the end of the sequence. If not, you take its lowest prime factor, divide the sum by said factor, then make *that* the next element in the sequence.

For example, if you start with `2` and `3`, then the sum is `5`. `5` is prime so it becomes the next element in the sequence. Then you repeat with the last two elements, which are now `3` and `5`. These two sum to `8`, which is not prime. Its lowest prime factor is `2`, so you divide `8` by `2`, giving `4`, which becomes the next element in the sequence.

The program asks you for two positive integers and the desired length of the sequence. Here are some examples (all containing 20 elements):

`[2 3 5 4 3 7 5 6 11 17 14 31 15 23 19 21 20 41 61 51]`

`[1 1 2 3 5 4 3 7 5 6 11 17 14 31 15 23 19 21 20 41]`

`[2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2 2]`

`[44 55 33 44 11 11 11 11 11 11 11 11 11 11 11 11 11 11 11 11]`

`[67 21 44 13 19 16 7 23 15 19 17 18 7 5 6 11 17 14 31 15]`

To run the program, just clone down the repo, have Go installed on your computer, `cd` into the repo and run `go run main.go` in a terminal emulator.