package main

import (
    "code.google.com/p/go-tour/pic"
    "fmt"
    "math"
)

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        pic[y] = make([]uint8, dx)
        for x := 0; x < dx; x++ {
            //squares := x^y
            //rightCornerhade := (x+y) / 2
            //swirls := x*y
            //vertLines := math.Pow(float64(x), 2)
            //horiLines := math.Pow(float64(y), 2)
            //horiShad := x
            //vertShade := y
            //bubbles := math.Pow(float64(x), 3) + math.Pow(float64(y), 3)
            //kindaCool := math.Pow(math.Pow(float64(x), 2) + math.Pow(float64(y), 2), 0.5)
            val := math.Pow(math.Pow(float64(x - x/2), 2) + math.Pow(float64(y - y/2), 2), 0.5)
            pic[y][x] = uint8(val)
        }
    }
    return pic
}

func main() {
    fmt.Print("data:image/png;base64,")
    pic.Show(Pic)
}
