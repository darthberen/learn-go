package main

import "fmt"

// functions do not need to be declared before being referenced
// types come after the variable name
// func add(x int, y int) int { is also written as the below since the int type is shared across the params
func add(x, y int) int {
    return x + y;
}

// any number of results can be returned from function
func swap(x, y string) (string, string) {
    return y, x
}

// result parameters can be named and act just like variables
// a return statement without any arguments returns the current values of the results
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(add(42, 13))

    a, b := swap("hello", "world")
    fmt.Println(a, b)

    fmt.Println(split(17))
}


