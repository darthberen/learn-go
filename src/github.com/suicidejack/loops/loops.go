package main

import "fmt"

func main() {
    // for loops are the only looping construct in go
    // note that () are not used (not even optional) and {} are required
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    // can leave pre and post statements empty
    sum = 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
    // or just implement a for like a C while
    sum = 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)

    // an infinite loop: 
    // for {}
}
