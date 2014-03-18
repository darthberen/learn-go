package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    fmt.Print("Go runs ons ")
    // case body breaks automatically unless it ends with a fallthrough statement
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X")
    case "linux":
        fmt.Println("Linux")
    default:
        fmt.Printf("%s\n", os)
    }

    fmt.Println("When's Saturday?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("today")
    case today + 1:
        fmt.Println("tomorrow")
    case today + 2:
        fmt.Println("in two days")
    default:
        fmt.Println("too far away")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning")
    case t.Hour() < 17:
        fmt.Println("Good afternoon")
    default:
        fmt.Println("Good evening")
    }
}
