package main

import (
    "fmt"
    "math"
    "math/cmplx"
)

const delta = 1e-10

func calc(start, num complex128) complex128 {
    return num - ((cmplx.Pow(num, complex128(3)) - start) / (3 * cmplx.Pow(num, complex128(2))))
}

func Cbrt(x complex128) complex128 {
    prevVal := complex128(0)
    val := x
    for cmplx.Abs(prevVal-val) > delta {
        prevVal = val
        val = calc(x, val)
    }
    return val
}

func main() {
    fmt.Println("[cube root 8]", "math:", math.Cbrt(8), "cmplx:", Cbrt(8))
    fmt.Println("[cube root 2]", "math:", math.Cbrt(2), "cmplx:", Cbrt(2))
}
