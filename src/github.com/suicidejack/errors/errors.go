package main

import (
    "fmt"
    "time"
)

// an error is anything that descirbes itself with an error string
// built in interface, error, with Error() method:
/*
type error interface {
    Error() string
}
*/

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err) // fmt knows to call Error() when asked to print an error
    }
}
