package main

import "fmt"

// struct is a collection of fields
type Vertex struct {
    X int
    Y int
}

var (
    a = Vertex{1, 2}
    b = &Vertex{1, 2} // pointer to newly allocated struct
    c = Vertex{X: 1} // struct literal - lists the values of its fields
    d = Vertex{}
)

func main() {
    v := Vertex{1, 2}
    fmt.Println(v)

    v.X = 4
    fmt.Println(v.X)

    // Go has pointers but no pointer arithmetic
    p := &v
    p.X = 1e9
    fmt.Println(v)

    fmt.Println(a, b, c, d)

    // new allocates a zeroed T value and returns pointer to it
    v1 := new(Vertex)
    var v2 *Vertex = new(Vertex) // equivalent to the above line
    fmt.Printf("v1: %T %v, v2: %T %v\n", v1, v1, v2, v2)
    v1.X, v1.Y = 11, 9
    fmt.Println(v1)
}
