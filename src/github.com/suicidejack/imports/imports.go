package main

import (
    "fmt"
    "math"
)
/*
 * could also write like:
 * import "fmt"
 * import "math"
 */

func main() {
    fmt.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))

    // a name is exported if it begins with a capital letter
    // fmt.Println(math.pi) this will not work since in pi p is lowercase
    fmt.Println(math.Pi)
}
