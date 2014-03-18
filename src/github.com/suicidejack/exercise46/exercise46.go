package main

import "fmt"

func fibonacci() func() int {
    seq := []int{0, 1}
    return func() int {
        result := seq[0] + seq[1]
        seq[0] = seq[1]
        seq[1] = result
        return result
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
