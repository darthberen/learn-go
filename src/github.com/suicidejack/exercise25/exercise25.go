package main

import (
    "fmt"
    "math"
)

const delta = 0.000000000000001;

func calc(start, num float64) float64 {
    return num - ((math.Pow(num, 2) - start) / (2 * num))
}

func Sqrt(x float64, it int) (val float64, iterations int) {
    val = x
    iterations = 0
    if it >= 0 {
        for ; it > 0; it-- {
            val = calc(x, val)
            iterations++
        }
    } else {
        prevVal := 0.0;
        for (val - prevVal >= delta) || (prevVal - val >= delta) {
            prevVal = val
            val = calc(x, val)
            iterations++
        }
    }
    return
}

func main() {
    for i := 1; i < 100; i++ {
        pkg := math.Sqrt(float64(i))
        it, itTimes := Sqrt(float64(i), 10)
        noIt, noItTimes := Sqrt(float64(i), -1)
        fmt.Printf("i %d, Math package: %g, mine (%d): %g [%g], mine (%d): %g [%g]\n",
            i,
            pkg,
            itTimes,
            it,
            pkg - it,
            noItTimes,
            noIt,
            pkg - noIt,
        )
    }
}
