package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    // () are optional but {} are required
    if (x < 0) {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x)) 
}

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        // variables declared inside if statements are available inside else blocks
        fmt.Printf("%g >= %g\n", v, lim) 
    }
    return lim
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 10), // note that the , is required here
    )
}
