package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

// Go doesn't have classes but can define methods on struct types
// method receiver appears in argument list betwen func keyword and method name
// doesn't necessarily need to be a pointer reciever since it isjust reading values (no mutation)
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer reciever avoids copying the value on each method call and can modify the value that its receiver points to
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

// can define a method on any type except types from another pkg or a basic type
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    v := &Vertex{3, 4}
    fmt.Println(v.Abs())

    v.Scale(5)
    fmt.Println(v, v.Abs())

    f := MyFloat(-math.Sqrt(2))
    fmt.Println(f.Abs())
}
