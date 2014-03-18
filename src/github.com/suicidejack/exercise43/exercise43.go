package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    words := strings.Fields(s)
    count := make(map[string]int)
    for _, val := range words {
        count[val]++
    }
    return count
}

func main() {
    wc.Test(WordCount)
}
