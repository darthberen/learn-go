package main

import (
    "fmt"
    "time"
)

/*
 * goroutines share the same address space - shared memory access
 * must be synchronized. sync package provides some primitives for this
 */

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world") // starts a goroutine
    say("hello")
}
