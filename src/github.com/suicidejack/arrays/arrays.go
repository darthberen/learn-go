package main

import "fmt"

func main() {
    // arrays can't be resized
    var arr [2]string
    arr[0] = "Hello"
    arr[1] = "World"
    fmt.Println(arr[0], arr[1])
    fmt.Println(arr)

    // slice points to an array of values and includes a length
    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)
    for i := 0; i < len(p); i++ {
        fmt.Printf("p[%d] == %d\n", i, p[i])
    }

    fmt.Println("p[1:4] ==", p[1:4])

    // missing low index implies 0
    fmt.Println("p[:3] ==", p[:3])

    // missing high index implies len(s)
    fmt.Println("p[4:] ==", p[4:])

    // slices still point to the original array
    x := p[:3]
    x[0] = 0
    fmt.Println(p)

    // slices are created with the make function
    // creates a zeroed array and returns slice that refers to that array
    a := make([]int, 5)
    printSlice("a", a)
    b := make([]int, 0, 5) // len(b)=0 cap(b)=5
    printSlice("b", b)
    c := b[:2]
    printSlice("c", c)
    d := c[2:5]
    printSlice("d", d)

    // nil slice as len and cap of 0
    var z []int;
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("nil!")
    }

    // range form of for lop iterates over a slice/map
    pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    bam := make([]int, 10)
    // drop ", value" to get just the index
    for i := range bam {
        bam[i] = 1 << uint(i)
    }
    // _ drops index or value
    for _, value := range bam {
        fmt.Printf("%d\n", value)
    }
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}
