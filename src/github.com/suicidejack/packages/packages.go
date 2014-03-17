package main

import (
    "fmt"
    "math/rand"
)

func main() {
    rand.Seed(5) // can use this to seed the random number generator
    // environment is deterministic so rand.Intn always returns the same number
    fmt.Println("My favorite number is", rand.Intn(10))
}
