package main

import "fmt"

/*
 * channels are a typed conduit through which you can send/recieve values
 * by default sends and recieves block until the other side is ready - this
 * allows goroutines to synchronize without explicit locks/conditions
 */
func sum(a []int, c chan int) {
    sum := 0
    for _, v := range a {
        sum += v
    }
    c <- sum // sends sum to c
}

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, y+x
    }
    /*
     * sender can close a channel to indicate that no more values will be sent
     * v, ok := <-c if ok is false than channel is closed
     * only sender should close a channel - sending on a close channel will
     * case a panic
     * closing is only necessary when receiver needs to know there are no more
     * values coming - such as terminating a range loop
     */
    close(c)
}

func main() {
    a := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(a[:len(a)/2], c)
    go sum(a[len(a)/2:], c)
    x, y := <-c, <-c // receive from c and assigns to x and y

    fmt.Println(x, y, x+y)

    /*
     * channels can be buffered - buffer length as secnd argument to make to
     * initialize a buffered channel
     * receives will block when the buffer is empty
     */
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    //ch <- 3 - overfilling the buffer causes deadlock
    fmt.Println(<-ch)
    fmt.Println(<-ch)

    c1 := make(chan int, 10)
    go fibonacci(cap(c1), c1)
    for i := range c1 { // recieves values repeatedly until closed
        fmt.Println(i)
    }
}
