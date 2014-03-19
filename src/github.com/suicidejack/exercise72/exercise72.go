package main

import (
    "code.google.com/p/go-tour/tree"
    "fmt"
)

func walkImpl(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        walkImpl(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        walkImpl(t.Right, ch)
    }
}

// Walk walks the tree t sending all values
// from the tree to channel ch
func Walk(t *tree.Tree, ch chan int) {
    walkImpl(t, ch)
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    for {
        v1, ok1 := <-ch1
        v2, ok2 := <-ch2

        if v1 != v2 || ok1 != ok2 {
            return false
        } else if !ok1 || !ok2 {
            break
        }
    }
    return true
}

func main() {
    /* example of Walk
       ch := make(chan int)
       go Walk(tree.New(2), ch)
       for v := range ch {
           fmt.Println(v)
       }
    */
    if Same(tree.New(1), tree.New(1)) {
        fmt.Println("equal")
    }

    if !Same(tree.New(1), tree.New(2)) {
        fmt.Println("not equal")
    }
}
