# Hello World

## Saying Hello

There are a bunch of ways we could get to an output that contains “Hello, world!” Preferably this would include a blockchain and neural network on a serverless architecture for maximum hype factor, but doing that would be a complete overkill — not to mention be overbudget for doing something as simple as printing to a standard output.
Instead, we’ll just do a permutation sort on the input string until we get the desired output of “Hello, world!”

Implementing this in Go is fairly straightforward.
We’ll take a slice out of the target string to turn it into runes and shuffle it using rand.Shuffle until it matches the target string.

## Links
https://medium.com/better-programming/hello-go-6721933be560
