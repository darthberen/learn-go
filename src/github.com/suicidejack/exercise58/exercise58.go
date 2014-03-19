package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    // fmt.Print(e) - calls string method of e - like a recursive call - infinite loop exception
    return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

const delta = 0.000000000000001

func calc(start, num float64) float64 {
    return num - ((math.Pow(num, 2) - start) / (2 * num))
}

func Sqrt(x float64, it int) (val float64, iterations int, err error) {
    if x < 0 {
        return 0.0, 0, ErrNegativeSqrt(x)
    } else if x == 0 {
        return 0.0, 0, nil
    }
    val = x
    iterations = 0
    if it >= 0 {
        for ; it > 0; it-- {
            val = calc(x, val)
            iterations++
        }
    } else {
        prevVal := 0.0
        for (val-prevVal >= delta) || (prevVal-val >= delta) {
            prevVal = val
            val = calc(x, val)
            iterations++
        }
    }
    return val, iterations, nil
}

func main() {
    for i := -1; i < 10; i++ {
        pkg := math.Sqrt(float64(i))
        it, itTimes, err := Sqrt(float64(i), 10)
        noIt, noItTimes, _ := Sqrt(float64(i), -1)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("i %d, Math package: %g, mine (%d): %g [%g], mine (%d): %g [%g]\n",
                i,
                pkg,
                itTimes,
                it,
                pkg-it,
                noItTimes,
                noIt,
                pkg-noIt,
            )
        }
    }
}
