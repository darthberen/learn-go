package main

import (
    "fmt"
    "math/cmplx"
    "math"
)

/* 
 * basic types are:
 * bool
 * string
 * int  int8  int16  int32  int64
 * uint uint8 uint16 uint32 uint64 uintptr
 * byte // alias for uint8
 * rune // alias for int32 (represents a unicode code print)
 * float32 float64
 * complex64 complex128
 */
var (
    ToBe    bool        = false
    MaxInt  uint64      = 1<<64 - 1
    z       complex128  = cmplx.Sqrt(-5 + 12i)
)

// constants can be character, string, boolean, or numeric values
// they can't be declared using := operator
// an untyped constant takes the type needed by its context
const Pi = 3.14

// numeric constants are high precision values
const (
    Big = 1 << 100
    Small = Big >> 99
)

func needInt(x int) int { 
    return x*10 + 1 
}

func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)

    x, y := 3, 4
    f2 := math.Sqrt(float64(3*3 + 4*4))
    z := int(f2) // assignment between different types requires explicit conversions
    fmt.Println(x, y, z)

    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")
    const Truth = true
    fmt.Println("Go rules?", Truth)

    fmt.Println(needInt(Small))
    //fmt.Println(needInt(Big)) // overflows int
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
