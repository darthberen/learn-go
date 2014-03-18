package main

import "fmt"

// var statement declares list of variables
// if an initializer is prent the type can be omitted since it is taken from the initializer type
var i, j int = 1, 2
var c, python, java = true, false, "no!"

func main() {
    fmt.Println(i, j, c, python, java)

    // inside function the short assignment statement (:=) can be used in place of var with implicit type
    k := 3
    ex1, ex2 := false, "im a short assignment example"
    fmt.Println(k, ex1, ex2)
}
