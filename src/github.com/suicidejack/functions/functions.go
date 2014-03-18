package main

import (
    "fmt"
    "math"
)

// functions do not need to be declared before being referenced
// types come after the variable name
// func add(x int, y int) int { is also written as the below since the int type is shared across the params
func add(x, y int) int {
    return x + y
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

// example of a closure (function value that references variables from outside its body)
// each closure is bound to its own sum variable
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    fmt.Println(add(42, 13))

    a, b := swap("hello", "world")
    fmt.Println(a, b)

    fmt.Println(split(17))

    // functions can be values
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(3, 4))

    // closure example
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}
