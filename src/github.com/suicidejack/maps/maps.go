package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex
var m2 = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -7439967,
    },
    "Google": {37.42202, -122.08408}, // top-level is just a type name can omit from elements
}

func main() {
    // maps must be created with make
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])

    fmt.Println(m2)

    //mutating maps
    m3 := make(map[string]int)
    m3["Answer"] = 42
    fmt.Println("The value:", m3["Answer"])

    m3["Answer"] = 48
    fmt.Println("The value:", m3["Answer"])

    delete(m3, "Answer")
    fmt.Println("The value:", m3["Answer"])

    v, ok := m3["Answer"]
    fmt.Println("The value:", v, "Present?", ok)
}
